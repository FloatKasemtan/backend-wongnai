package endpoints

import (
	"github.com/FloatKasemtan/endpoints/covid"
	"github.com/gin-gonic/gin"
)

func RegisterEndpoint(r *gin.Engine) {
	// Create group
	group := r.Group("/covid")

	group.GET("/summary", covid.SummaryHandler)
}
