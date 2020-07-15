package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"
	"weshierNext/handler"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging is a middleware func that logs each request
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		reg := regexp.MustCompile("(/v1/user/login)")
		if !reg.MatchString(path) {
			return
		}

		// Skip for the health check requests.
		if path == "/api/sd/health" || path == "/api/sd/ram" || path == "/api/sd/cpu" || path == "/api/sd/disk" {
			return
		}

		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// The basic informations.
		method := c.Request.Method
		ip := c.ClientIP()

		logger.Logger.Debug(fmt.Sprintf("New request come in, path: %s, Method: %s, body `%s`", path, method, string(bodyBytes)))
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		// get code and message
		var response handler.Response
		var byteData = blw.body.Bytes()
		if err := json.Unmarshal(byteData, &response); err != nil {
			logger.Logger.Error("response body can not unmarshal to model.Response struct",
				zap.String("error", err.Error()),
				zap.String("body", string(byteData)),
			)
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}

		logger.Logger.Info(fmt.Sprintf("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message))
	}
}
