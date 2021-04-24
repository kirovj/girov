package main

import (
	"girov"
	"net/http"
)

func main() {
	r := girov.New()
	r.GET("/", func(c *girov.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *girov.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *girov.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *girov.Context) {
		c.JSON(http.StatusOK, girov.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
