package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

//	type Mid struct {
//		a *auth.Auth
//	}
type key string

const TraceIdKey key = "1"

func (m Mid) Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.NewString()
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, TraceIdKey, traceId)
		req := c.Request.WithContext(ctx)
		c.Request = req

		log.Info().Str("Trace Id", traceId).Str("Method", c.Request.Method).
			Str("URL Path", c.Request.URL.Path).Msg("request started")
		defer log.Info().Str("Trace Id", traceId).Str("Method", c.Request.Method).
			Str("URL Path", c.Request.URL.Path).
			Int("status Code", c.Writer.Status()).Msg("Request processing completed")
		c.Next()
	}
}

// func NewMid(a *auth.Auth) (Mid, error) {
// 	// It first checks if 'a' is nil
// 	// 'a' should not be nil because 'nil' indicates that the 'Auth' object does not exist.
// 	if a == nil {
// 		// An error is returned when 'a' is 'nil'.
// 		return Mid{}, errors.New("auth can't be nil")
// 	}
// 	//If 'a' is not 'nil', a new 'Mid' instance is returned with 'a' as a field.
// 	// A nil error is returned, indicating that there were no issues with the initialization.
// 	return Mid{a: a}, nil
// }
