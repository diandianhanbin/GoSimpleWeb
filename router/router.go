package router

import (
	"SvBlogApi/handler/sd"
	"SvBlogApi/handler/user"
	"SvBlogApi/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
	}

	svcd := g.Group("/sd")
	{
		svcd.GET("/Health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
