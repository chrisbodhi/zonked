package main

import "net/http"
import "time"
import "github.com/gin-gonic/gin"

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    router.GET("/", func(c *gin.Context) {
	type Message struct {
	  Message  string
	  Author   string
	}

	type Messages []Message

	var msgs = Messages{
	  Message{
		  "Can no longer stand people's excitement about ping-pong.",
		  "structy",
	  },

	  Message{
		  "So tired",
		  "GH botnet",
	  },

	}

        c.HTML(http.StatusOK, "index.tmpl", gin.H{
	    "title": "iamtiredofeverything.today",
	    "msgs": msgs,
	})
    })

    router.GET("/api", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Can no longer stand people's excitement about ping-pong.",
            "author": "cb",
            "time": time.Now(),
        })
    })
    router.Run() // listen and server on 0.0.0.0:8080
}
