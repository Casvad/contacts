package middlewares

import (
	"bytes"
	"contacts/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		detectedErrors := c.Errors.ByType(gin.ErrorTypeAny)
		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError errors.UserError
			switch err.(type) {
			case errors.UserError:
				parsedError = err.(errors.UserError)
				c.IndentedJSON(parsedError.GetStatusServer(), parsedError)
				c.Abort()
			default:
				defaultError := errors.NewUserError("errors", http.StatusInternalServerError)
				defaultError.Message = err.(error).Error()
				c.IndentedJSON(http.StatusInternalServerError, defaultError)
				c.Abort()
			}
		}
	}
}
