// Package middleware pprof serves via its HTTP server runtime profiling data
// in the format expected by the pprof visualization tool.
// Ref: https://github.com/gin-contrib/pprof
package middleware

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

const (
	// PprofDefaultPrefix url prefix of pprof
	PprofDefaultPrefix = "/debug/pprof"
)

func getPprofPrefix(prefixOptions ...string) string {
	prefix := PprofDefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}

// PprofRegister the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.Engine. prefixOptions is a optional. If not prefixOptions,
// the default path prefix is used, otherwise first prefixOptions will be path prefix.
func PprofRegister(r *gin.Engine, prefixOptions ...string) {
	PprofRouteRegister(&(r.RouterGroup), prefixOptions...)
}

// PprofRouteRegister the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.GrouterGroup. prefixOptions is a optional. If not prefixOptions,
// the default path prefix is used, otherwise first prefixOptions will be path prefix.
func PprofRouteRegister(rg *gin.RouterGroup, prefixOptions ...string) {
	prefix := getPprofPrefix(prefixOptions...)

	prefixRouter := rg.Group(prefix)
	{
		prefixRouter.GET("/", gin.WrapF(pprof.Index))
		prefixRouter.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		prefixRouter.GET("/profile", gin.WrapF(pprof.Profile))
		prefixRouter.POST("/symbol", gin.WrapF(pprof.Symbol))
		prefixRouter.GET("/symbol", gin.WrapF(pprof.Symbol))
		prefixRouter.GET("/trace", gin.WrapF(pprof.Trace))
		prefixRouter.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		prefixRouter.GET("/block", gin.WrapH(pprof.Handler("block")))
		prefixRouter.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		prefixRouter.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		prefixRouter.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		prefixRouter.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}
}
