package main

import (
	"github.com/gin-gonic/gin"
)

func (r routes) initRoutes() {
	r.router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"Status": "alive"}) })
}

func (r routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}
