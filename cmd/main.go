package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"servicerepository.git/internal/utils"
)

type routes struct {
	router *gin.Engine
}

func main() {
	err := run()
	if err != nil {
		zap.L().Fatal("Service to run", zap.Error(err))
		panic(err)
	}
}

func run() error {
	l, _ := zap.NewDevelopment()
	defer l.Sync()
	zap.ReplaceGlobals(l)

	if os.Getenv("NODE_ENV") == "development" {
		setLocalEnvironmentVariables()
	}

	zap.L().Info("=== Starting Service ===")

	r := routes{
		router: gin.Default(),
	}

	r.initRoutes()

	port := utils.GetEnvVariable("PORT", "8080")

	if os.Getenv("NODE_ENV") == "development" {
		r.Run("127.0.0.1:" + port)
	} else {
		r.Run(":" + port)
	}

	return nil

}

func setLocalEnvironmentVariables() {}
