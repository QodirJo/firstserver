package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/Exam4/ApiGatawey/genproto/health_sync"
	"github.com/gin-gonic/gin"
)

// @Summary CreateGenetic
// @Description CreateGenetic
// @Tags Genetic
// @Accept json
// @Produce json
// @Param request body health_sync.GeneticCreate true "CreateGenetic"
// @Success 201 {object} health_sync.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /genetic/create [post]
func (h *Handler) CreateGenetic(c *gin.Context) {
	var request health_sync.GeneticCreate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Generic.CreateGenetic(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Genetic": &request,
	})

}

// @Summary UpdateGenetic
// @Description UpdateGenetic updates a genetic record based on the provided id and user_id.
// @Tags Genetic
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param body body health_sync.GeneticUpdate true "Genetic Update Request"
// @Success 200 {object} health_sync.GeneticUpdate
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /genetic/update/{id} [put]
func (h *Handler) UpdateGenetic(c *gin.Context) {
	id := c.Param("id")

	var request health_sync.GeneticUpdate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	request.Id = id

	_, err := h.Generic.UpdateGenetic(context.Background(), &request)
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
		"message": "Genetic record updated successfully",
		"Genetic": &request,
	})
}

// @Summary DeleteGenetic
// @Description DeleteGenetic
// @Tags Genetic
// @Accept json
// @Produce json
// @Param id path string true "Genetic ID"
// @Success 200 {object} health_sync.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /genetic/delete/{id} [delete]
func (h *Handler) DeleteGenetic(c *gin.Context) {
	id := c.Param("id")
	req := health_sync.GetById{Id: id}

	_, err := h.Generic.DeleteGenetic(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Genetic deleted successfully")
}

// @Summary GetGenetic
// @Description GetGenetic
// @Tags Genetic
// @Accept json
// @Produce json
// @Param id query string false "id"
// @Success 200 {object} health_sync.GeneticRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /genetic/get/{id} [get]
func (h *Handler) GetGenetic(c *gin.Context) {
	id := c.Query("id")
	req := health_sync.GetById{Id: id}

	res, err := h.Generic.GetGenetic(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Genetic": res,
	})

}

// @Summary ListGenetic
// @Description ListGenetic
// @Tags Genetic
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} health_sync.GeneticList
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /genetic/list [get]
func (h *Handler) ListGenetic(c *gin.Context) {
	// Get limit and offset from query parameters
	limit := c.Query("limit")
	offset := c.Query("offset")

	// Convert limit and offset to integers
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		limitInt = 10 // default value
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil || offsetInt < 0 {
		offsetInt = 0 // default value
	}

	// Convert int to int32 for the Filter struct
	request := health_sync.Filter{
		Limit:  int32(limitInt),
		Offset: int32(offsetInt),
	}

	res, err := h.Generic.ListGenetic(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
