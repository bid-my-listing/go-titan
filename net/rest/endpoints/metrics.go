package endpoints

import (
	"expvar"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jucardi/go-titan/net/rest/middleware/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// AddLogLevel adds the `/metrics` endpoint to the given router.
func AddMetrics(router *gin.Engine) {
	router.GET("/memory-details", gin.WrapH(expvar.Handler()))
	router.GET("/metrics", func(context *gin.Context) {
		start := time.Now()
		resp := map[string]interface{}{
			"hardware":  metrics.GetHardwareStats(),
			"resources": metrics.GetStats(),
		}
		latency := time.Since(start)
		resp["metrics_latency"] = latency.String()
		context.IndentedJSON(200, resp)
	})
	router.GET("/metrics/prometheus", func(context *gin.Context) {
		h := promhttp.Handler()
		h.ServeHTTP(context.Writer, context.Request)
	})
}
