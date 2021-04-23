package main

import (
	"girov"
	"net/http"
)

func main() {
	r := girov.New()

	r.GET("/", func(c *girov.Context) {
		c.HTML(http.StatusOK, "<h1>hello</h1>")
	})

	r.GET("/hello", func(c *girov.Context) {
		// expect /hello?name=jack
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *girov.Context) {
		c.JSON(http.StatusOK, girov.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
