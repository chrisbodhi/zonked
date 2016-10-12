package main

import "time"
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/api", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
            "author": "cb",
            "time": time.Now(),
        })
    })
    r.Run() // listen and server on 0.0.0.0:8080
}
