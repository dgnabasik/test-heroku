package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// ErrorHandler func
func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": ctx.Errors,
		})
	}
}

// GetPort func:
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return port
}

func main() {
	gin.SetMode(gin.ReleaseMode) // Switch to "release" mode in production; or export GIN_MODE=release
	router := gin.Default()
	router.Use(cors.New(cors.Config{ // assume this config is overwritten by ContextOptions()?
		AllowOrigins: []string{"*"}, // hostName
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//AllowHeaders: []string{"X-CSRF-Token", "Authorization"},
		AllowHeaders:     []string{"User-Agent", "Referrer", "Host", "Token", "Accept", "Content-Type", "Origin", "Content-Length", "X-Requested-With", "Accept-Encoding"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowOriginFunc: func(origin string) bool {
			return true // origin == hostName
		},
		MaxAge: 86400,
	}))

	router.Use(ErrorHandler)
	router.LoadHTMLGlob("templates/*")
	router.Use(static.Serve("/", static.LocalFile("./build", false)))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	apiPort := GetPort()
	api := "Handling REST-API calls on " + ":" + apiPort
	fmt.Println(api)
	router.Run(":" + apiPort)
}
