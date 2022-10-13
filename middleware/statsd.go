package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"

	"fmt"
	"time"
)

var client *statsd.Client
var userOptions Options

var handlerFunc = func(c *gin.Context) {
	startTime := time.Now()
	c.Next()

	if client != nil {
		prefix := userOptions.getPrefix()
		path := strings.TrimPrefix(c.FullPath(), "/")
		path = strings.ReplaceAll(path, "/", "_")
		path = strings.ReplaceAll(path, "*", "_")
		path = strings.ReplaceAll(path, ":", "_")

		metricPrefix := fmt.Sprintf("%s.%s", prefix, path)
		metricPrefix = strings.TrimPrefix(metricPrefix, ".")

		// send status code
		status := c.Writer.Status()
		client.Increment(fmt.Sprintf("%s.status_code.%d", metricPrefix, status))

		// send response time
		duration := time.Since(startTime).Seconds() * 1000 // in milliseconds
		client.Timing(fmt.Sprintf("%s.response_time", metricPrefix), duration)
	}
}

// New will setup middleware and return handler
func New(opts Options) gin.HandlerFunc {
	userOptions = opts
	addr := userOptions.getAddress()
	var clientErr error
	client, clientErr = statsd.New(statsd.Address(addr))
	if clientErr != nil {
		client = nil
		printLog(fmt.Sprintf("Failed connecting to statsd - %s", clientErr.Error()), errorLevel)
	} else {
		printLog("Successfully connected to statsd", infoLevel)
	}

	return handlerFunc
}
