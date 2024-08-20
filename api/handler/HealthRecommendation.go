package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Exam4/ApiGatawey/genproto/health_sync"
	"github.com/gin-gonic/gin"
)

// @Summary CreateHealthRecommendation
// @Description CreateHealthRecommendation
// @Tags HealthRecommendation
// @Accept json
// @Produce json
// @Param request body health_sync.HealthRecommendationCreate true "CreateHealthRecommendation"
// @Success 201 {object} health_sync.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health-recommendation/create [post]
func (h *Handler) CreateHealthRecommendation(c *gin.Context) {
	var request health_sync.HealthRecommendationCreate

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req, err := json.Marshal(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.Producer.ProduceMessages("healthRecommendationscreate", req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"HealthRecommendation": &request,
	})
}

// @Summary UpdateHealthRecommendation
// @Description UpdateHealthRecommendation updates a health_recommendation record based on the provided id and user_id.
// @Tags HealthRecommendation
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Param body body health_sync.HealthRecommendationUpdate true "HealthRecommendation Update Request"
// @Success 200 {object} health_sync.HealthRecommendationUpdate
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health-recommendation/update/{id} [put]
func (h *Handler) UpdateHealthRecommendation(c *gin.Context) {
	id := c.Query("id")

	var request health_sync.HealthRecommendationUpdate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	request.Id = id
	req, err := json.Marshal(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.Producer.ProduceMessages("healthRecommendationsupdate", req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"HealthRecommendation": &request,
	})
}

// @Summary DeleteHealthRecommendation
// @Description Delete a HealthRecommendation record based on the provided id
// @Tags HealthRecommendation
// @Accept json
// @Produce json
// @Param id path string true "HealthRecommendation ID"
// @Success 200 {object} health_sync.GetById
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health-recommendation/delete/{id} [delete]
func (h *Handler) DeleteHealthRecommendation(c *gin.Context) {
	id := c.Param("id")
	req := health_sync.GetById{Id: id}

	_, err := h.HealthRecommendations.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "HealthRecommendation deleted successfully",
	})
}

// @Summary GetHealthRecommendation
// @Description Retrieve a HealthRecommendation record based on the provided id
// @Tags HealthRecommendation
// @Accept json
// @Produce json
// @Param id query string true "HealthRecommendation ID"
// @Success 200 {object} health_sync.GetById
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health-recommendation/get/{id} [get]
func (h *Handler) GetHealthRecommendation(c *gin.Context) {
	id := c.Query("id")
	req := health_sync.GetById{Id: id}

	res, err := h.HealthRecommendations.Get(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"HealthRecommendation": res,
	})
}

// / @Summary ListHealthRecommendation
// @Description List HealthRecommendation records with pagination
// @Tags HealthRecommendation
// @Accept json
// @Produce json
// @Param limit query string false "limit" default(10)
// @Param offset query string false "offset" default(0)
// @Success 200 {object} health_sync.GetAllHealthRecommendationsRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health-recommendation/list [get]
func (h *Handler) ListHealthRecommendation(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		limitInt = 10 // default value
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil || offsetInt < 0 {
		offsetInt = 0 // default value
	}

	request := health_sync.Filter{
		Limit:  int32(limitInt),
		Offset: int32(offsetInt),
	}

	res, err := h.HealthRecommendations.List(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
