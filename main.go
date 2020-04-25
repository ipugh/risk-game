package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
	r := gin.Default()
    r.LoadHTMLGlob("assets/*.html")
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
