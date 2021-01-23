package main

import (
	"flag"

	"github.com/gin-gonic/gin"
)

var (
	certFile = flag.String("cert", "/cert/server.pem", "cert file")
	keyFile  = flag.String("key", "/cert/server.key", "key file")
)

func main() {
	flag.Parse()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	r.RunTLS(":8443", *certFile, *keyFile)
}
