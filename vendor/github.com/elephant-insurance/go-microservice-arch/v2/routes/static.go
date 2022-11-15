package routes

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// These funcs and methods replace the matching Gin funcs and methods
// This is necessary so that we can handle the base (app name) of the path
// Gin will not allow ANY overlapping routes

// StaticFile registers a single route in order to serve a single file of the local filesystem.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (r *Router) StaticFile(relativePath, filepath string) {
	if strings.Contains(relativePath, ":") || strings.Contains(relativePath, "*") {
		panic("URL parameters can not be used when serving a static file")
	}

	handler := func(c *gin.Context) {
		c.File(filepath)
	}

	r.get(routeLabelStatic, relativePath, handler, true)
}

// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//     router.Static("/static", "/var/www")
func (r *Router) Static(relativePath, root string) {
	r.StaticFS(relativePath, gin.Dir(root, false))
}

// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
// Gin by default user: gin.Dir()
func (r *Router) StaticFS(relativePath string, fs http.FileSystem) {
	if strings.Contains(relativePath, ":") || strings.Contains(relativePath, "*") {
		panic("URL parameters can not be used when serving a static folder")
	}
	handler := r.createStaticHandler(relativePath, fs)
	urlPattern := path.Join(relativePath, "/*filepath")

	// Register GET handler
	r.get(routeLabelStatic, urlPattern, handler, true)

	// need this?
	// r.HEAD(urlPattern, handler)
}

func (r *Router) createStaticHandler(relativePath string, fs http.FileSystem) gin.HandlerFunc {
	absolutePath := r.calculateAbsolutePath(relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *gin.Context) {
		if _, noListing := fs.(*onlyFilesFS); noListing {
			c.Writer.WriteHeader(http.StatusNotFound)
		}

		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		f, err := fs.Open(file)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		f.Close()

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

// LoadHTMLGlob loads HTML files identified by glob pattern
// and associates the result with HTML renderer.
func (r *Router) LoadHTMLGlob(pattern string) {
	if r == nil || r.Engine == nil {
		return
	}

	de := r.Engine.Delims(HTMLTemplateLeftDelimiter, HTMLTemplateRightDelimiter)
	templ := template.Must(template.New("").Delims(HTMLTemplateLeftDelimiter, HTMLTemplateRightDelimiter).Funcs(de.FuncMap).ParseGlob(pattern))

	de.SetHTMLTemplate(templ)
}

// LoadHTMLFiles loads a slice of HTML files
// and associates the result with HTML renderer.
func (r *Router) LoadHTMLFiles(files ...string) {
	filePaths := make([]string, len(files))
	testMode := r != nil && r.Engine != nil && gin.Mode() == gin.TestMode && r.TestPathPrefix != ``
	for i := 0; i < len(files); i++ {
		if testMode {
			filePaths[i] = r.TestPathPrefix + files[i]
		} else {
			filePaths[i] = files[i]
		}
	}
	de := r.Engine.Delims(HTMLTemplateLeftDelimiter, HTMLTemplateRightDelimiter)
	templ := template.Must(template.New("").Delims(HTMLTemplateLeftDelimiter, HTMLTemplateRightDelimiter).Funcs(de.FuncMap).ParseFiles(filePaths...))
	de.SetHTMLTemplate(templ)
}

// these are all copies of small Gin internal utility funcs

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func (r *Router) calculateAbsolutePath(relativePath string) string {
	return joinPaths(r.pathBase, relativePath)
}

type neuteredReaddirFile struct {
	http.File
}

type onlyFilesFS struct {
	fs http.FileSystem
}

// Open conforms to http.Filesystem.
func (fs onlyFilesFS) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

// Readdir overrides the http.File default implementation.
func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	// this disables directory listing
	return nil, nil
}
