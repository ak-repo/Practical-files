package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomLogger(param gin.LogFormatterParams) string {
	// Example log format:
	return fmt.Sprintf("[%s] %s | %3d | %13v | %15s | %-7s  %s\n",
		// time
		param.TimeStamp.Format(time.RFC1123),
		// client IP
		param.ClientIP,
		// status code
		param.StatusCode,
		// latency
		param.Latency,
		// method
		param.Method,
		// path
		param.Path,
		// error if any
		param.ErrorMessage,
	)
}
