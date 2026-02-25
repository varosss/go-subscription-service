package http

import "github.com/gin-gonic/gin"

func (h *SubscriptionHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/subscriptions/total-cost", h.CalculateTotalCost)
	r.POST("/subscriptions", h.Create)
	r.GET("/subscriptions", h.List)

	r.GET("/subscriptions/:id", h.GetByID)
	r.PATCH("/subscriptions/:id", h.Update)
	r.DELETE("/subscriptions/:id", h.Delete)
}
