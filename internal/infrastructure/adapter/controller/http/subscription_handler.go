package http

import (
	"go-subscription-service/internal/application/usecase"
	"go-subscription-service/internal/domain/valueobject"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	createUC *usecase.CreateSubscriptionUseCase
	getUC    *usecase.GetSubscriptionUseCase
	updateUC *usecase.UpdateSubscriptionUseCase
	deleteUC *usecase.DeleteSubscriptionUseCase
	listUC   *usecase.ListSubscriptionsUseCase
}

func NewSubscriptionHandler(
	createUC *usecase.CreateSubscriptionUseCase,
	getUC *usecase.GetSubscriptionUseCase,
	updateUC *usecase.UpdateSubscriptionUseCase,
	deleteUC *usecase.DeleteSubscriptionUseCase,
	listUC *usecase.ListSubscriptionsUseCase,
) *SubscriptionHandler {
	return &SubscriptionHandler{
		createUC: createUC,
		getUC:    getUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
		listUC:   listUC,
	}
}

func (h *SubscriptionHandler) Create(c *gin.Context) {
	var req struct {
		UserID      string  `json:"user_id"`
		ServiceName string  `json:"service_name"`
		Price       int64   `json:"price"`
		StartDate   string  `json:"start_date"`
		EndDate     *string `json:"end_date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := valueobject.ParseUserID(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("01-2006", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var endDate *time.Time
	if req.EndDate != nil {
		t, err := time.Parse("01-2006", *req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		endDate = &t
	}

	res, err := h.createUC.Execute(c.Request.Context(), usecase.CreateSubscriptionCommand{
		UserID:      userID,
		ServiceName: req.ServiceName,
		Price:       req.Price,
		StartDate:   startDate,
		EndDate:     endDate,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"subscription_id": res.Subscription.ID,
	})
}

func (h *SubscriptionHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := valueobject.ParseSubscriptionID(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subscription id"})
		return
	}

	var req struct {
		ServiceName *string `json:"service_name"`
		Price       *int64  `json:"price"`
		StartDate   *string `json:"start_date"`
		EndDate     *string `json:"end_date"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var startDate *time.Time
	if req.StartDate != nil {
		t, err := time.Parse("01-2006", *req.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		startDate = &t
	}

	var endDate *time.Time
	if req.EndDate != nil {
		t, err := time.Parse("01-2006", *req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		endDate = &t
	}

	err = h.updateUC.Execute(c.Request.Context(), usecase.UpdateSubscriptionCommand{
		SubscriptionID: id,
		ServiceName:    req.ServiceName,
		Price:          req.Price,
		StartDate:      startDate,
		EndDate:        endDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *SubscriptionHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	subscriptionID, err := valueobject.ParseSubscriptionID(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subscription id"})
		return
	}

	res, err := h.getUC.Execute(c.Request.Context(), usecase.GetSubscriptionCommand{
		SubscriptionID: subscriptionID,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	c.JSON(http.StatusOK, res.Subscription)
}

func (h *SubscriptionHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	subscriptionID, err := valueobject.ParseSubscriptionID(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subscription id"})
		return
	}

	err = h.deleteUC.Execute(c.Request.Context(), usecase.DeleteSubscriptionCommand{
		SubscriptionID: subscriptionID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *SubscriptionHandler) List(c *gin.Context) {
	var userID *valueobject.UserID
	if uid := c.Query("user_id"); uid != "" {
		parsed, err := valueobject.ParseUserID(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
			return
		}
		userID = &parsed
	}

	var serviceName *string
	if sn := c.Query("service_name"); sn != "" {
		serviceName = &sn
	}

	var fromDate, toDate *time.Time
	if fd := c.Query("from_date"); fd != "" {
		t, err := time.Parse("01-2006", fd)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from_date"})
			return
		}
		fromDate = &t
	}
	if td := c.Query("to_date"); td != "" {
		t, err := time.Parse("01-2006", td)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to_date"})
			return
		}
		toDate = &t
	}

	var limit, offset *int
	if l := c.Query("limit"); l != "" {
		v, err := strconv.Atoi(l)
		if err == nil {
			limit = &v
		}
	}
	if o := c.Query("offset"); o != "" {
		v, err := strconv.Atoi(o)
		if err == nil {
			offset = &v
		}
	}

	res, err := h.listUC.Execute(c.Request.Context(), usecase.ListSubscriptionsCommand{
		UserID:      userID,
		ServiceName: serviceName,
		FromDate:    fromDate,
		ToDate:      toDate,
		Limit:       limit,
		Offset:      offset,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res.Subs)
}
