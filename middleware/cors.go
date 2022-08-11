package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CORSMiddleware struct{}

func (m *CORSMiddleware) CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins:    []string{},
		// AllowOriginFunc: func(origin string) bool {},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"content-type", "token"},
		AllowCredentials: false,
		// ExposeHeaders:          []string{},
		MaxAge: 0,
		// AllowWildcard:          false,
		// AllowBrowserExtensions: false,
		// AllowWebSockets:        false,
		// AllowFiles:             false,
	})
}

func NewCORSMiddleware() *CORSMiddleware {
	return &CORSMiddleware{}
}
