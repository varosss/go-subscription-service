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
	createUC             *usecase.CreateSubscriptionUseCase
	getUC                *usecase.GetSubscriptionUseCase
	updateUC             *usecase.UpdateSubscriptionUseCase
	deleteUC             *usecase.DeleteSubscriptionUseCase
	listUC               *usecase.ListSubscriptionsUseCase
	calculateTotalCostUC *usecase.CalculateTotalCostUseCase
}

func NewSubscriptionHandler(
	createUC *usecase.CreateSubscriptionUseCase,
	getUC *usecase.GetSubscriptionUseCase,
	updateUC *usecase.UpdateSubscriptionUseCase,
	deleteUC *usecase.DeleteSubscriptionUseCase,
	listUC *usecase.ListSubscriptionsUseCase,
	calculateTotalCostUC *usecase.CalculateTotalCostUseCase,
) *SubscriptionHandler {
	return &SubscriptionHandler{
		createUC:             createUC,
		getUC:                getUC,
		updateUC:             updateUC,
		deleteUC:             deleteUC,
		listUC:               listUC,
		calculateTotalCostUC: calculateTotalCostUC,
	}
}

// @Summary Create a new subscription
// @Description Creates a subscription for a user
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body CreateSubRequest true "Subscription to create"
// @Success 201 {object} CreateSubResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(c *gin.Context) {
	var req CreateSubRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	userID, err := valueobject.ParseUserID(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	startDate, err := time.Parse("01-2006", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var endDate *time.Time
	if req.EndDate != nil {
		t, err := time.Parse("01-2006", *req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
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
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, CreateSubResponse{ID: res.Subscription.ID})
}

// UpdateSubscription godoc
// @Summary Update a subscription
// @Description Update a subscription by ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Param subscription body UpdateSubRequest true "Subscription update"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := valueobject.ParseSubscriptionID(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid subscription id"})
		return
	}

	var req UpdateSubRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var startDate *time.Time
	if req.StartDate != nil {
		t, err := time.Parse("01-2006", *req.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}
		startDate = &t
	}

	var endDate *time.Time
	if req.EndDate != nil {
		t, err := time.Parse("01-2006", *req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
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
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetSubscription godoc
// @Summary Get subscription by ID
// @Description Get a subscription by subscription ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} dto.Subscription
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	subscriptionID, err := valueobject.ParseSubscriptionID(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid subscription id"})
		return
	}

	res, err := h.getUC.Execute(c.Request.Context(), usecase.GetSubscriptionCommand{
		SubscriptionID: subscriptionID,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "subscription not found"})
		return
	}

	c.JSON(http.StatusOK, res.Subscription)
}

// DeleteSubscription godoc
// @Summary Delete a subscription
// @Description Delete subscription by ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	subscriptionID, err := valueobject.ParseSubscriptionID(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid subscription id"})
		return
	}

	err = h.deleteUC.Execute(c.Request.Context(), usecase.DeleteSubscriptionCommand{
		SubscriptionID: subscriptionID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListSubscriptions godoc
// @Summary List subscriptions
// @Description List subscriptions with optional filters
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param user_id query string false "User ID filter"
// @Param service_name query string false "Service name filter"
// @Param from_date query string false "Start date filter (MM-YYYY)"
// @Param to_date query string false "End date filter (MM-YYYY)"
// @Param limit query int false "Limit number of results"
// @Param offset query int false "Offset"
// @Success 200 {array} dto.Subscription
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [get]
func (h *SubscriptionHandler) List(c *gin.Context) {
	var userID *valueobject.UserID
	if uid := c.Query("user_id"); uid != "" {
		parsed, err := valueobject.ParseUserID(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user_id"})
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
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid from_date"})
			return
		}
		fromDate = &t
	}
	if td := c.Query("to_date"); td != "" {
		t, err := time.Parse("01-2006", td)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid to_date"})
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
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res.Subs)
}

// CalculateTotalCost godoc
// @Summary      Calculate total subscription cost
// @Description  Calculates the total cost of all subscriptions for a user and/or service within a given period. If from_date/to_date are not provided, defaults to the current month.
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        user_id      query     string  false  "UUID of the user"
// @Param        service_name query     string  false  "Filter by subscription service name"
// @Param        from_date    query     string  false  "Start of period, format MM-YYYY, defaults to start of current month"
// @Param        to_date      query     string  false  "End of period, format MM-YYYY, defaults to end of current month"
// @Success      200  {object}  CalculateTotalCostResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /subscriptions/total-cost [get]
func (h *SubscriptionHandler) CalculateTotalCost(c *gin.Context) {
	var userID *valueobject.UserID
	if uid := c.Query("user_id"); uid != "" {
		parsed, err := valueobject.ParseUserID(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user_id"})
			return
		}
		userID = &parsed
	}

	var serviceName *string
	if sn := c.Query("service_name"); sn != "" {
		serviceName = &sn
	}

	var fromDate, toDate time.Time
	var err error

	if fd := c.Query("from_date"); fd != "" {
		fromDate, err = time.Parse("01-2006", fd)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid from_date, expected MM-YYYY"})
			return
		}
	} else {
		now := time.Now()
		fromDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	}

	if td := c.Query("to_date"); td != "" {
		toDate, err = time.Parse("01-2006", td)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid to_date, expected MM-YYYY"})
			return
		}
		toDate = time.Date(toDate.Year(), toDate.Month()+1, 0, 23, 59, 59, 0, toDate.Location())
	} else {
		toDate = fromDate.AddDate(0, 1, -1).Add(time.Hour*23 + time.Minute*59 + time.Second*59)
	}

	res, err := h.calculateTotalCostUC.Execute(c.Request.Context(), usecase.CalculateTotalCostCommand{
		UserID:      userID,
		ServiceName: serviceName,
		FromDate:    fromDate,
		ToDate:      toDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CalculateTotalCostResponse{
		Total:       res.Total,
		FromDate:    fromDate.Format("01-2006"),
		ToDate:      toDate.Format("01-2006"),
		UserID:      (*string)(userID),
		ServiceName: serviceName,
	})
}
