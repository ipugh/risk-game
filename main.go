package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func main() {
	r := gin.Default()
    r.Static("/", "./assets")

    r.POST("/pick", func(c *gin.Context) {
        var territory Pick
        err := c.BindJSON(&territory)
        if err != nil {
            c.AbortWithError(400, err).SetType(gin.ErrorTypeBind)
        } else {
            fmt.Println(territory)

            c.JSON(200, gin.H{
                "success": true,
            })
        }
    })

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
