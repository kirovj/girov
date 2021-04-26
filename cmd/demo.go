package main

import (
	"girov"
	"net/http"
)

func main() {
	r := girov.New()
	r.GET("/", func(c *girov.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *girov.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
