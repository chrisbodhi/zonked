package main

import "net/http"
import "time"
import "github.com/gin-gonic/gin"

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
	    "title": "Main SSSsssite",
	})
    })

    router.GET("/api", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
            "author": "cb",
            "time": time.Now(),
        })
    })
    router.Run() // listen and server on 0.0.0.0:8080
}
