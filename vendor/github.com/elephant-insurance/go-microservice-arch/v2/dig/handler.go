package dig

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/pprof"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/timing"
	"github.com/gin-gonic/gin"
)

// ServeDiagnostics is a Gin HandlerFunc for serving diagnostics info
func ServeDiagnostics(c *gin.Context) {
	params := c.Param("params")
	paramList := strings.Split(params, "/")
	// this can come to us a couple of different ways, so we just skip ahead to an action we understand
	for _, thisParam := range paramList {
		switch strings.ToLower(thisParam) {
		case "":
			continue
		case timingsPath:
			ServeTimings(c)
			return
		case memStatsPath:
			ServeMemStats(c)
			return
		case pprofPath:
			pprofOption := ""
			if len(paramList) > 1 {
				pprofOption = strings.ToLower(paramList[1])
			}
			ServePProf(c, pprofOption)
			return
		}
	}

	// either there were no params, or none that we recognized, so just serve up the plain dig page
	serveDiagnostics(c.Writer, c.Request)
}

// addHandler adds a diagnostic handler to the selected path + /diagnostics
func addHandler(engine *gin.Engine, basePath string) {
	fs, ss := "/", "/"
	if strings.HasPrefix(basePath, "/") {
		fs = ""
	}
	if strings.HasSuffix(basePath, "/") {
		ss = ""
	}
	rp := fmt.Sprintf("%v%v%v%v", fs, basePath, ss, DefaultPath)

	engine.GET(rp, ServeDiagnostics)
}

// serveDiagnostics does the actual work of serving up the basic diagnostics page
func serveDiagnostics(w http.ResponseWriter, r *http.Request) {
	var c context.Context
	if r != nil {
		c = r.Context()
	}

	lw := log.ForFunc(c)
	if r == nil {
		// not much we can do here
		lw.Info(ErrorMsgNilRequest)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	url := r.URL
	if url == nil {
		lw.Info(ErrorMsgNilURL)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	q := url.Query().Get(runTestsParameterName)

	code := http.StatusOK
	runTests := q != "" && q != "false"

	result := getCurrentSystemInfo(runTests, r)
	if runTests && result.FailedTests != nil && *result.FailedTests > 0 {
		code = http.StatusInternalServerError
	}

	resultJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		lw.WithError(err).Warn(ErrorMsgSerializeResults)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, err = w.Write(resultJSON)
	if err != nil {
		lw.WithError(err).Warn(ErrorMsgWritingResult)
	}
}

// ServeMemStats dumps mem stats to the client
func ServeMemStats(c *gin.Context) {
	afterparm := c.Query(paramNameAfter)

	var sq uint64
	var rows []*memStat
	var err error
	if sq, err = strconv.ParseUint(afterparm, 10, 64); err == nil && sq > 0 {
		rows = memStats.After(sq)
	} else {
		rows = memStats.Earliest(0)
	}

	csvParm := c.Query(paramNameCSV)
	if csv, err := strconv.ParseBool(csvParm); csv && err == nil {
		buf := strings.Builder{}
		buf.WriteString(memStatCSVHeader + "\n")
		for _, v := range rows {
			if v != nil {
				buf.WriteString(v.ToCSV() + "\n")
			}
		}
		c.Data(http.StatusOK, mimeTypeCSV, []byte(buf.String()))
		return
	}

	c.JSON(http.StatusOK, rows)
}

// ServeTimings dumps timings data to the client
func ServeTimings(c *gin.Context) {
	afterparm := c.Query(paramNameAfter)

	var sq uint64
	var rows []*timing.Timing
	var err error
	if sq, err = strconv.ParseUint(afterparm, 10, 64); err == nil && sq > 0 {
		rows = timing.ApplicationTimings.After(sq)
	} else {
		rows = timing.ApplicationTimings.Latest(0)
	}

	csvParm := c.Query(paramNameCSV)
	if csv, err := strconv.ParseBool(csvParm); csv && err == nil {
		buf := strings.Builder{}
		buf.WriteString(timing.TimingCSVHeader + "\n")
		for _, v := range rows {
			if v != nil {
				buf.WriteString(v.ToCSV() + "\n")
			}
		}
		c.Data(http.StatusOK, mimeTypeCSV, []byte(buf.String()))
		return
	}
	spew.Dump(rows)
	c.JSON(http.StatusOK, rows)
}

func ServePProf(c *gin.Context, action string) {
	switch action {
	case "cmdline":
		pprof.Cmdline(c.Writer, c.Request)
		return
	case "profile":
		pprof.Profile(c.Writer, c.Request)
		return
	case "symbol":
		pprof.Symbol(c.Writer, c.Request)
		return
	case "trace":
		pprof.Trace(c.Writer, c.Request)
		return
	case "allocs":
		pprof.Handler("allocs").ServeHTTP(c.Writer, c.Request)
		return
	case "block":
		pprof.Handler("block").ServeHTTP(c.Writer, c.Request)
		return
	case "goroutine":
		pprof.Handler("goroutine").ServeHTTP(c.Writer, c.Request)
		return
	case "heap":
		pprof.Handler("heap").ServeHTTP(c.Writer, c.Request)
		return
	case "mutex":
		pprof.Handler("mutex").ServeHTTP(c.Writer, c.Request)
		return
	case "threadcreate":
		pprof.Handler("threadcreate").ServeHTTP(c.Writer, c.Request)
		return
	default:
		pprof.Index(c.Writer, c.Request)
		return
	}
}
