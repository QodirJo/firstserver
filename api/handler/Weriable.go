package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Exam4/ApiGatawey/genproto/health_sync"
	"github.com/gin-gonic/gin"
)

// @Summary CreateWearable
// @Description CreateWearable
// @Tags Wearable
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body health_sync.WearableCreate true "CreateWearable"
// @Success 201 {object} health_sync.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /wearable/create [post]
func (h *Handler) CreateWearable(c *gin.Context) {
	var request health_sync.WearableCreate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req, errr := json.Marshal(&request)
	if errr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errr.Error(),
		})
		return
	}

	err := h.Producer.ProduceMessages("wearables", req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Wearable": &request,
	})
}

// @Summary UpdateWearable
// @Description UpdateWearable updates a wearable record based on the provided id and user_id.
// @Tags Wearable
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Param body body health_sync.WearableUpdate true "Wearable Update Request"
// @Success 200 {object} health_sync.WearableUpdate
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /wearable/update/{id} [put]
func (h *Handler) UpdateWearable(c *gin.Context) {
	id := c.Param("id")

	var request health_sync.WearableUpdate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	request.Id = id

	_, err := h.Wearable.Update(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Wearable": &request,
	})
}

// @Summary DeleteWearable
// @Description DeleteWearable deletes a wearable record based on the provided id and user_id.
// @Tags Wearable
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} health_sync.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /wearable/delete/{id} [delete]
func (h *Handler) DeleteWearable(c *gin.Context) {
	id := c.Param("id")

	req := health_sync.GetById{Id: id}

	_, err := h.Wearable.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wearable record deleted successfully",
	})
}

// @Summary GetWearable
// @Description GetWearable retrieves a wearable record based on the provided id.
// @Tags Wearable
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} health_sync.GetById
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /wearable/get/{id} [get]
func (h *Handler) GetWearable(c *gin.Context) {
	id := c.Query("id")

	req := health_sync.GetById{Id: id}

	res, err := h.Wearable.Get(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Wearable": res,
	})
}

// @Summary ListWearable
// @Description ListWearable retrieves a list of wearable records with optional pagination.
// @Tags Wearable
// @Accept json
// @Produce json
// @Param limit query string false "limit" default(10)
// @Param offset query string false "offset" default(0)
// @Success 200 {object} health_sync.WearableList
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /wearable/list [get]
func (h *Handler) ListWearable(c *gin.Context) {
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

	req := health_sync.Filter{
		Limit:  int32(limitInt),
		Offset: int32(offsetInt),
	}

	res, err := h.Wearable.List(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

