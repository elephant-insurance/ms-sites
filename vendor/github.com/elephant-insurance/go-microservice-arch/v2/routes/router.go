package routes

import (
	"context"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/bc"
	"github.com/elephant-insurance/go-microservice-arch/v2/cfg"
	dig "github.com/elephant-insurance/go-microservice-arch/v2/dig"
	log "github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
	"github.com/gin-gonic/gin"
)

// The Elephant router wraps the Gin router, making arbitrary URL structures possible
// The built-in Gin router is very inflexible, making routes with multiple route parameters impossible
// This package adds a default handler for each RESTful HTTP method (Delete, Get, Patch, Post, Put)
// This default handler parses the URL of the request and calls the "real" handler for the intended route

// Here's how it works:

// 1. At startup, we call the New(g *gin.Engine) method to create a router struct and save the Gin Engine reference.
// 2. Also at startup, routes are added to the router struct using the usual route functions (GET, POST, PUT, etc.).
// 		Each call to one of these functions appends a Route object to the corresponding array in the Router object.
// 3. The route functions also add a catch-all wildcard handler (if one hasn't been added already) for the specified
//		method to the Gin Engine. This is the handler that will be called by Gin.
// 4. When the app services a client call, Gin passes control to the catch-all wildcard handler for the appropriate method.
//		This catch-all function then cycles through the list of Route objects to see whether one matches.
// 5. If a matching route is found, any route parameters specified in the path argument will be parsed out and added to
//		the gin.Context, so that they can be retrieved easily by the controller. The handler corresponding to the matching
//		route is called.

type Router struct {
	// we use the sync.Mutex to lock down access to a single thread during startup
	// We assume that there will be no changes to the routing tables once the app is fully initialized
	sync.Mutex
	Engine              *gin.Engine
	LastRouteParameters map[string]string
	// LastRouteSelected holds a pointer to the last route we resolved -- use this for testing the routes
	LastRouteSelected *Route
	finalized         bool
	listenPort        string
	pathBase          string
	routesForDelete   []*Route
	routesForGet      []*Route
	routesForPatch    []*Route
	routesForPost     []*Route
	routesForPut      []*Route
	// set testMode to true to test routes with Gin in test mode
	// otherwise, the router will not run the actual route code
	testMode bool
	// TestPathPrefix is the string we prepend to the relative path for static files in testing
	TestPathPrefix string
}

const slash = "/"

// methodsAdded keeps track of which HTTP methods we have added routes for
// This is used to validate CORS Allowed Methods settings
var methodsAdded = map[string]bool{
	http.MethodConnect: false,
	http.MethodDelete:  false,
	http.MethodGet:     true,
	http.MethodHead:    true,
	http.MethodOptions: true,
	http.MethodPatch:   false,
	http.MethodPost:    false,
	http.MethodPut:     false,
	http.MethodTrace:   false,
}

// methodsAllowed keeps track of which HTTP methods we have allowed in CORS
// This is used to validate CORS Accepted Methods settings
var methodsAllowed = map[string]bool{
	http.MethodConnect: false,
	http.MethodDelete:  false,
	http.MethodGet:     true,
	http.MethodHead:    true,
	http.MethodOptions: true,
	http.MethodPatch:   false,
	http.MethodPost:    false,
	http.MethodPut:     false,
	http.MethodTrace:   false,
}

// New creates a router stuct for us to populate and sets up the main controller action
// This func WILL THROW A FATAL ERROR if it is sent a bad basePath OR if the path regex does not compile
func New(requiredConfig cfg.Configurator, g *gin.Engine) *Router {
	lw := log.ForFunc(nil)
	if g == nil {
		lw.Fatal(`attempt to initialize routes with nil Gin engine`)
	}

	basePath := requiredConfig.GetInstanceName()

	methods := requiredConfig.GetAllowedMethods()
	for _, v := range methods {
		// make sure this is a valid method:
		if _, ok := methodsAllowed[v]; !ok {
			lw.Fatal(`attempted to Allow unrecognized HTTP method: ` + v)
		}
		methodsAllowed[v] = true
	}

	port := requiredConfig.GetListenPort()

	return new(basePath, port, g)
}

// new is a private version of New for testing
func new(basePath, port string, g *gin.Engine) *Router {
	lw := log.ForFunc(nil)
	// make sure basePath is either empty or in /path form:
	if basePath != "" {
		if !strings.HasPrefix(basePath, slash) {
			basePath = slash + basePath
		}

		basePath = strings.TrimSuffix(basePath, slash)
	}

	digPath := "/diagnostics"
	digReggie, err := regexp.Compile(basePath + digPath + "$")
	lw.IfError(err).Fatal("failed to create regex for base diagnostics path")
	digReggieWithParams, err := regexp.Compile(basePath + digPath + "/(?P<params>.*)")
	lw.IfError(err).Fatal("failed to create regex for diagnostics path with parameters")

	// "heartbeat" route
	g.HEAD("/", bc.SayOK)
	if basePath != "" {
		g.HEAD(basePath, bc.SayOK)
	}

	// We add the two diagnostics routes here so that they're always available
	newRoutesForGet := []*Route{
		{
			Handler:   dig.ServeDiagnostics,
			Label:     routeLabelDiagnostics,
			PathRegex: digReggie,
		},
		{
			Handler:    dig.ServeDiagnostics,
			Label:      routeLabelDiagnostics,
			PathParams: []string{"params"},
			PathRegex:  digReggieWithParams,
		},
	}

	rtn := &Router{
		Engine:         g,
		finalized:      false,
		listenPort:     port,
		pathBase:       basePath,
		routesForGet:   newRoutesForGet,
		testMode:       false,
		TestPathPrefix: defaultTestPathPrefix,
	}

	g.GET(basePath+"/*allparms", rtn.handleGet)

	return rtn
}

// Listen performs some final checks and marks the router finalized.
// The router will not work until it is finalized.
func (r *Router) Listen() {
	lw := log.ForFunc(context.Background())

	// make sure we have not allowed any methods we aren't listening for
	for k, v := range methodsAllowed {
		if v {
			added, ok := methodsAdded[k]
			if !ok || !added {
				lw.Fatal(`method allowed in CORS for which there is no route: ` + k)
			}
		}
	}

	if r != nil && r.Engine != nil {
		r.finalized = true
		msg := `application started and listening`
		evt := uf.EventFactory.New(&enumerations.Event.ApplicationStartup.ID, nil, msg)
		lw.WithEvent(evt).Info(msg)
		err := r.Engine.Run(r.listenPort)
		lw.IfError(err).Fatal(`Gin Run returned error`)
	} else {
		lw.Fatal(`attempt to Run with nil Gin Engine or router`)
	}
}

func (r *Router) FinalizeForTest(requiredConfig cfg.Configurator) {
	lw := log.ForFunc(context.Background())
	if requiredConfig.GetEnvironment() != enumerations.ServiceEnvironment.Testing.ID {
		lw.Fatal(`attempted to finalize router for testing but not in testing environment`)
	}

	r.finalized = true
}

// handle fulfills the request specified in the context
// If a Route is found in the Router that matches the URL of the request, that Route's handler is run
// Otherwise, we don't have an appropriate handler, so return 404
// Every request that runs through here is timed in the diagnostics package
func (r *Router) handle(c *gin.Context) {
	if c == nil || c.Request == nil || c.Request.URL == nil || c.Request.URL.Path == "" || c.Request.Method == "" {
		// nothing we can do
		bc.SayNotFound(c)
		return
	}

	lw := log.ForFunc(c, "method", c.Request.Method, "request-path", c.Request.URL.Path)

	// reset the last route pointer
	r.LastRouteSelected = nil

	var routes []*Route

	switch c.Request.Method {
	case http.MethodDelete:
		routes = r.routesForDelete
	case http.MethodGet:
		routes = r.routesForGet
	case http.MethodPatch:
		routes = r.routesForPatch
	case http.MethodPost:
		routes = r.routesForPost
	case http.MethodPut:
		routes = r.routesForPut
	default:
		// this should never happen, because we let Gin handle other methods
		// But if it does happen, we have to return 404
		lw.Warn("no handlers found for method")
		bc.SayNotFound(c)
		return
	}

	for i := 0; i < len(routes); i++ {
		thisRoute := routes[i]
		if thisRoute == nil {
			continue
		}

		if ok, pmap := thisRoute.handles(c); ok {
			// lw.WithConsoleField("route-label", thisRoute.Label).WithConsoleField("regex", thisRoute.PathRegex.String()).WithConsoleField("params", c.Param("allparms")).Debug("route handles path, dispatching")
			// if TXTYPE hasn't been set yet, we set it now
			lw.SetTransactionType(thisRoute.Label)
			r.LastRouteSelected = thisRoute
			// spew.Dump("LRS: " + r.LastRouteSelected.Label)
			r.LastRouteParameters = pmap
			label, path := thisRoute.timingParams()
			writeLogs := !(thisRoute.Silent || label == nil || *label == routeLabelDiagnostics)
			var timer dig.Timer
			if writeLogs {
				timer = dig.StartServiceTiming(c, label, path)
			}
			if gin.Mode() != gin.TestMode || r.testMode {
				thisRoute.Handler(c)
			}
			status := 0
			if c.Writer != nil {
				status = c.Writer.Status()
			}
			if timer != nil {
				timer.Stop(status)
			}
			return
		}

		// lw.WithConsoleField("route-label", thisRoute.Label).WithConsoleField("regex", thisRoute.PathRegex.String()).Debug("route does not match path")
	}

	// didn't find a matching route
	// lw.Debug("no route matched path")
	bc.SayNotFound(c)
}

////////////////////////////////////////////////
//             Route add funcs                //
//											  //
// Use these to add routes to the application //
////////////////////////////////////////////////

// Each of these funcs checks to see whether it is the first handler we have added for this HTTP method
// If it is, we initialize the list of handlers for that method, AND we assign the wildcard handler to Gin.

// DELETE adds an HTTP Delete route to the router
func (r *Router) DELETE(label, relativePath string, handler gin.HandlerFunc) {
	lw := log.ForFunc(context.Background())
	if r == nil {
		lw.Fatal(errMessageNilRouter)
	}

	if r.finalized {
		lw.Fatal(errMessageRouterFinalized)
	}

	if allowed, ok := methodsAllowed[http.MethodDelete]; !ok || !allowed {
		lw.Fatal(`attempted to add DELETE route without allowing DELETE method in CORS`)
	}

	r.Lock()
	defer r.Unlock()

	methodsAdded[http.MethodDelete] = true

	// we use the nullity of the routes list to keep track of whether we've set up this method yet
	if r.routesForDelete == nil {
		r.routesForDelete = []*Route{}
		r.Engine.DELETE("/*allparms", r.handleDelete)
	}

	newRoute := makeRoute(label, r.pathBase+relativePath, handler)

	r.routesForDelete = append(r.routesForDelete, newRoute)
}

// GET adds an HTTP Get route to the router
func (r *Router) GET(label, relativePath string, handler gin.HandlerFunc) {
	r.get(label, relativePath, handler, false)
}

func (r *Router) get(label, relativePath string, handler gin.HandlerFunc, silent bool) {
	lw := log.ForFunc(context.Background(), "label", label, "relative-path", relativePath).Debug("adding route for GET")
	if r == nil {
		lw.Fatal(errMessageNilRouter)
	}

	if r.finalized {
		lw.Fatal(errMessageRouterFinalized)
	}

	// TODO: diagnostics

	r.Lock()
	defer r.Unlock()

	// we use the nullity of the routes list to keep track of whether we've set up this method yet
	if r.routesForGet == nil {
		lw.Warn("called GET while routesForGet was nil, should never happen")
		r.routesForGet = []*Route{}
		r.Engine.GET("/*allparms", r.handleGet)
	}

	newRoute := makeRoute(label, r.pathBase+relativePath, handler)
	newRoute.Silent = silent

	r.routesForGet = append(r.routesForGet, newRoute)
}

// PATCH adds an HTTP Patch route to the router
func (r *Router) PATCH(label, relativePath string, handler gin.HandlerFunc) {
	lw := log.ForFunc(context.Background())
	if r == nil {
		lw.Fatal(errMessageNilRouter)
	}

	if r.finalized {
		lw.Fatal(errMessageRouterFinalized)
	}

	if allowed, ok := methodsAllowed[http.MethodPatch]; !ok || !allowed {
		lw.Fatal(`attempted to add PATCH route without allowing PATCH method in CORS`)
	}
	// TODO: diagnostics

	r.Lock()
	defer r.Unlock()

	methodsAdded[http.MethodPatch] = true

	// we use the nullity of the routes list to keep track of whether we've set up this method yet
	if r.routesForPatch == nil {
		r.routesForPatch = []*Route{}
		r.Engine.PATCH("/*allparms", r.handlePatch)
	}

	newRoute := makeRoute(label, r.pathBase+relativePath, handler)

	r.routesForPatch = append(r.routesForPatch, newRoute)
}

// POST adds an HTTP Post route to the router
func (r *Router) POST(label, relativePath string, handler gin.HandlerFunc) {
	lw := log.ForFunc(context.Background())
	if r == nil {
		lw.Fatal(errMessageNilRouter)
	}

	if r.finalized {
		lw.Fatal(errMessageRouterFinalized)
	}

	if allowed, ok := methodsAllowed[http.MethodPost]; !ok || !allowed {
		lw.Fatal(`attempted to add POST route without allowing POST method in CORS`)
	}

	r.Lock()
	defer r.Unlock()

	methodsAdded[http.MethodPost] = true

	// we use the nullity of the routes list to keep track of whether we've set up this method yet
	if r.routesForPost == nil {
		r.routesForPost = []*Route{}
		r.Engine.POST("/*allparms", r.handlePost)
	}

	newRoute := makeRoute(label, r.pathBase+relativePath, handler)

	r.routesForPost = append(r.routesForPost, newRoute)
}

// PUT adds an HTTP Put route to the router
func (r *Router) PUT(label, relativePath string, handler gin.HandlerFunc) {
	lw := log.ForFunc(context.Background())
	if r == nil {
		lw.Fatal(errMessageNilRouter)
	}

	if r.finalized {
		lw.Fatal(errMessageRouterFinalized)
	}

	if allowed, ok := methodsAllowed[http.MethodPut]; !ok || !allowed {
		lw.Fatal(`attempted to add PUT route without allowing PUT method in CORS`)
	}

	r.Lock()
	defer r.Unlock()

	methodsAdded[http.MethodPut] = true

	// we use the nullity of the routes list to keep track of whether we've set up this method yet
	if r.routesForPut == nil {
		r.routesForPut = []*Route{}
		r.Engine.PUT("/*allparms", r.handlePut)
	}

	newRoute := makeRoute(label, r.pathBase+relativePath, handler)

	r.routesForPut = append(r.routesForPut, newRoute)
}

/*/////////////////////////////////////////////////
//               Wildcard handlers               //
//											     //
// These are the handlers that Gin runs directly //
/////////////////////////////////////////////////*/

// Each of these functions checks the list of handlers in the Router for the specified method.
// If there could be a match (i.e., the list is not empty), we call r.handle(c) to complete processing of the request.

func (r *Router) handleDelete(c *gin.Context) {
	if r == nil || !r.finalized {
		log.ForFunc(context.Background()).Error(errMessageRouterNotFinalized)
		bc.SayNotFound(c)
		return
	}

	if len(r.routesForDelete) < 1 {
		bc.SayNotFound(c)
		return
	}

	r.handle(c)
}

func (r *Router) handleGet(c *gin.Context) {
	if r == nil || !r.finalized {
		log.ForFunc(context.Background()).Error(errMessageRouterNotFinalized)
		bc.SayNotFound(c)
		return
	}

	if len(r.routesForGet) < 1 {
		bc.SayNotFound(c)
		return
	}

	r.handle(c)
}

func (r *Router) handlePatch(c *gin.Context) {
	if r == nil || !r.finalized {
		log.ForFunc(context.Background()).Error(errMessageRouterNotFinalized)
		bc.SayNotFound(c)
		return
	}

	if len(r.routesForPatch) < 1 {
		bc.SayNotFound(c)
		return
	}

	r.handle(c)
}

func (r *Router) handlePost(c *gin.Context) {
	if r == nil || !r.finalized {
		log.ForFunc(context.Background()).Error(errMessageRouterNotFinalized)
		bc.SayNotFound(c)
		return
	}

	if len(r.routesForPost) < 1 {
		bc.SayNotFound(c)
		return
	}

	r.handle(c)
}

func (r *Router) handlePut(c *gin.Context) {
	if r == nil || !r.finalized {
		log.ForFunc(context.Background()).Error(errMessageRouterNotFinalized)
		bc.SayNotFound(c)
		return
	}

	if len(r.routesForPut) < 1 {
		bc.SayNotFound(c)
		return
	}

	r.handle(c)
}
