package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

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

// TestData struct
type TestData struct {
	Text string
}

// GetTestData func
func GetTestData(ctx *gin.Context) {
	testData := TestData{Text: "This is text."}
	ctx.JSON(200, gin.H{
		"TestData": testData,
	})
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
	router.LoadHTMLGlob(filepath.Join(os.Getenv("TEMPLATE_DIR"), "*"))
	router.Static("/static", os.Getenv("STATIC_DIR"))
	router.Use(static.Serve("/", static.LocalFile(os.Getenv("BUILD_DIR"), true)))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",
			gin.H{"title": "Home Page"},
		)
	})

	/*	test := router.Group("/")
		{
			test.GET("/test", GetTestData)
		}*/

	apiPort := GetPort()
	api := "Handling REST-API calls on " + ":" + apiPort
	fmt.Println(api)
	router.Run(":" + apiPort)
}
