package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/static"
    "fmt"
)

func main() {
	r := gin.Default()
    api := r.Group("/api")

    api.POST("/pick", func(c *gin.Context) {
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

    api.POST("/pickreinforce", func(c *gin.Context) {
        var territory PickReinforce
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

    api.POST("/regroup", func(c *gin.Context) {
        var regroup Regroup
        err := c.BindJSON(&regroup)
        if err != nil {
            c.AbortWithError(400, err).SetType(gin.ErrorTypeBind)
        } else {
            fmt.Println(regroup)

            c.JSON(200, gin.H{
                "success": true,
            })
        }
    })

    api.POST("/reinforce", func(c *gin.Context) {
        var reinforce Reinforce
        err := c.BindJSON(&reinforce)
        if err != nil {
            c.AbortWithError(400, err).SetType(gin.ErrorTypeBind)
        } else {
            fmt.Println(reinforce)

            c.JSON(200, gin.H{
                "success": true,
                "remaining": 100,
            })
        }
    })

    api.POST("/attack", func(c *gin.Context) {
        var attack Attack
        err := c.BindJSON(&attack)
        if err != nil {
            c.AbortWithError(400, err).SetType(gin.ErrorTypeBind)
        } else {
            fmt.Println(attack)

            c.JSON(200, gin.H{
                "success": true,
                "remaining": 100,
            })
        }
    })

    api.GET("/board", func(c *gin.Context) {
        var board Gameboard
        board.Turn = "Isaac"
        board.Gamestate = "pick"
        board.Turnstate = "start"

        // Add players
        var p Player
        p.Name = "Isaac"
        p.Color = "00ff00"
        board.Players = append(board.Players, p) 

        // Add territories
        var a = Territory{"steve", 1, 1}
        var b = Territory{"steve", 1, 2}
        var w World
        w.Territories = append(w.Territories, a, b)
        board.Board = w

        fmt.Println(board)
        c.JSON(200, board)
    })


    // thanks to gin for not supporting advanced enough routing to do this
    // r.Static("/", "./assets")
    r.Use(static.Serve("/", static.LocalFile("./assets", true)))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
