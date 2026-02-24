package http

import "github.com/gin-gonic/gin"

func (h *SubscriptionHandler) RegisterRoutes(r *gin.Engine) {
	s := r.Group("/subscriptions")
	{
		s.POST("", h.Create)
		s.GET("/:id", h.GetByID)
		s.PATCH("/:id", h.Update)
		s.DELETE("/:id", h.Delete)
		s.GET("", h.List)
	}
}
