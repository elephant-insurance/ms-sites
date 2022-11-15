package glog

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/gin-gonic/gin"
)

// IMP NOTES:
// 1. This Handler is design for GIN to replace the GIN in build Recovery Handler
// 2. It is based on GIN Recover Handler that is out of the box for value. In future any new release of GIN,
//    after upgrade following items could add value to be reverified with gin.Recovery
//    a. Value for "pre-known" (dunno, centerDot, dot, slash) they are also subjected to change with new GoLang versions
//    b. Stack index current index that is in use is 3 to extract the Stack details needed - this rarely change but still worth checking
// 3. RecoverHandler itself
//    a. Only verification that needs to be done is related to "broken pipe" exception - very rare to change
// 4. Or Easiest every upgrade verify if there is any change in following functions and pull those changes to this file:
//    a. stack
//    b. source
//    c. function

// NOTE: stack-handling code has been moved to the log package

// RecoveryHandler returns a middleware for a given writer that recovers from any panics and writes a 500 if there was one.
func RecoveryHandler() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		defer func() {
			lw := log.ForFunc(ginContext)
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// err => Here isn't an error type object, logging it as WithConsoleField
				lw.WithStack().WithError(errors.New(fmt.Sprint(err))).Error("[Recovery] panic recovered")

				// If the connection is dead, we can't write a status to it.
				if brokenPipe {
					ginContext.Abort()
				} else {
					ginContext.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		ginContext.Next()
	}
}
