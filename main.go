package main

import (
	"github.com/FloatKasemtan/endpoints"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	endpoints.RegisterEndpoint(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
