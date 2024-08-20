package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/Exam4/ApiGatawey/genproto/health_sync"
	"github.com/gin-gonic/gin"
)

// @Summary CreateLifeStyle
// @Description Create a new Lifestyle record
// @Tags Lifestyle
// @Accept json
// @Produce json
// @Param request body health_sync.LifestyleCreate true "Create Lifestyle"
// @Success 201 {object} health_sync.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /lifestyle/create [post]
func (h *Handler) CreateLifeStyle(c *gin.Context) {
	var request health_sync.LifestyleCreate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Lifestyle.Create(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Lifestyle": &request,
	})
}

// @Summary UpdateLifeStyle
// @Description Update a Lifestyle record based on the provided id and user_id
// @Tags Lifestyle
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param body body health_sync.LifestyleUpdate true "Lifestyle Update Request"
// @Success 200 {object} health_sync.LifestyleUpdate
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /lifestyle/update/{id} [put]
func (h *Handler) UpdateLifeStyle(c *gin.Context) {
	id := c.Param("id")

	var request health_sync.LifestyleUpdate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	request.Id = id

	_, err := h.Lifestyle.Update(context.Background(), &request)
	if err != nil {
		if strings.Contains(err.Error(), "no records matched") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lifestyle record updated successfully",
		"Lifestyle": &request,
	})
}

// @Summary DeleteLifeStyle
// @Description Delete a Lifestyle record based on the provided id
// @Tags Lifestyle
// @Accept json
// @Produce json
// @Param id path string true "Lifestyle ID"
// @Success 200 {object} health_sync.GetById
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /lifestyle/delete/{id} [delete]
func (h *Handler) DeleteLifeStyle(c *gin.Context) {
	id := c.Param("id")
	req := health_sync.GetById{Id: id}

	_, err := h.Lifestyle.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Lifestyle deleted successfully")
}

// @Summary GetLifeStyle
// @Description Retrieve a Lifestyle record based on the provided id
// @Tags Lifestyle
// @Accept json
// @Produce json
// @Param id query string false "id"
// @Success 200 {object} health_sync.GetById
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /lifestyle/get/{id} [get]
func (h *Handler) GetLifeStyle(c *gin.Context) {
	id := c.Query("id")
	req := health_sync.GetById{Id: id}

	res, err := h.Lifestyle.Get(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Lifestyle": res,
	})
}

// @Summary ListLifeStyle
// @Description List Lifestyle records with pagination
// @Tags Lifestyle
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} health_sync.LifestyleList
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /lifestyle/list [get]
func (h *Handler) ListLifeStyle(c *gin.Context) {
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

	res, err := h.Lifestyle.List(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
