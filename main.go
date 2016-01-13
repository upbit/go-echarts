package main

import (
    // "fmt"
    "net/http"
    "time"
    // GitHub
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Assets
    r.Static("/static", "./assets/static")
    r.StaticFile("/favicon.ico", "./assets/favicon.ico")
    r.StaticFile("/robots.txt", "./assets/robots.txt")
    r.Static("/c", "./charts")
    // Templates
    r.LoadHTMLGlob("templates/*")

    // Routers
    r.GET("/", func(c *gin.Context) {
        c.Redirect(301, "/index")
    })

    r.GET("/index", func(c *gin.Context) {
        c.HTML(200, "index.tmpl", gin.H{
            "ECharts入门示例 - 柱状图": "bar.html",
        })
    })

    r.GET("/stat", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Error handling
    r.NoRoute(func(c *gin.Context) {
        c.HTML(404, "not_found.tmpl", gin.H{})
    })

    s := &http.Server{
        Addr:           ":4000",
        Handler:        r,
        ReadTimeout:    2 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
