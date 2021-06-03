package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true))) //vue
	// router.Static("/static", "./ui-dist/static")//react

	router.LoadHTMLGlob("./frontend/dist/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", "text/html")
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", "text/html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
		log.Printf("Defaulting to port %s", port)
	}
	router.Run(":" + port)
}
