package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Exam4/ApiGatawey/genproto/health_sync"
	"github.com/gin-gonic/gin"
)

// @Summary CreateRecords
// @Description CreateRecords
// @Tags Records
// @Accept json
// @Produce json
// @Param request body health_sync.MedicalRecordCreate true "CreateRecords"
// @Success 201 {object} health_sync.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /records/create [post]
func (h *Handler) CreateRecords(c *gin.Context) {
	var request health_sync.MedicalRecordCreate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Records.Create(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Records": &request,
	})
}

// @Summary UpdateRecords
// @Description UpdateRecords updates a records record based on the provided id and user_id.
// @Tags Records
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param body body health_sync.MedicalRecordUpdate true "Records Update Request"
// @Success 200 {object} health_sync.MedicalRecordUpdate
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /records/update/{id} [put]
func (h *Handler) UpdateRecords(c *gin.Context) {
	id := c.Param("id")

	var request health_sync.MedicalRecordUpdate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	request.Id = id
	_, err := h.Records.Update(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Records": &request,
	})
}

// @Summary DeleteRecords
// @Description DeleteRecords deletes a records record based on the provided id and user_id.
// @Tags Records
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} health_sync.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /records/delete/{id} [delete]
func (h *Handler) DeleteRecords(c *gin.Context) {
	id := c.Param("id")

	req := health_sync.GetById{Id: id}

	_, err := h.Records.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Records record deleted successfully",
	})
}

// @Summary GetRecords
// @Description GetRecords
// @Tags Records
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} health_sync.GetById
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /records/get/{id} [get]
func (h *Handler) GetRecords(c *gin.Context) {
	id := c.Query("id")
	req := health_sync.GetById{Id: id}

	res, err := h.Records.Get(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Records": res,
	})
}

// @Summary ListRecords
// @Description ListRecords
// @Tags Records
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} health_sync.Filter
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /records/list [get]
func (h *Handler) ListRecords(c *gin.Context) {
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
	res, err := h.Records.List(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Records": res,
	})

}
