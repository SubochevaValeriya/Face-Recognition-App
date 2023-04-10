package handler

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Add Staff
// @Tags staff
// @Description add staff to the DB
// @ID add-staff
// @Accept  json
// @Produce  json
// @Param input body Staff
// @Success 200 {object} Staff
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /staff [post]
// AddStaff is made for adding staff information
func (h *Handler) AddStaff(c *gin.Context) {
	var input models.Staff

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	staff, err := h.services.AddStaff(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, staff)

	newSuccessResponse("adding staff", input.Name)
}

// @Summary Update Staff
// @Tags staff
// @Description update staff information
// @ID update-staff
// @Accept  json
// @Produce  json
// @Param input body Staff
// @Success 200 {object} Staff
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /staff [put]
// UpdateStaff is made for updating staff information
func (h *Handler) UpdateStaff(c *gin.Context) {
	var input models.Staff

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	staff, err := h.services.AddStaff(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, staff)

	newSuccessResponse("updating staff", input.Name)
}

// @Summary Get Staff
// @Tags staff
// @Description get staff by ID
// @ID get-staff
// @Accept  json
// @Produce  json
// @Success 200 {object} Staff
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /staff [get]
// GetStaff is used to get staff info by ID
func (h *Handler) GetStaff(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	staff, err := h.services.GetStaff(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, staff)

	newSuccessResponse("getting staff", staff.Name)
}

// @Summary Get AllStaff
// @Tags staff
// @Description get all staff
// @ID get-all-staff
// @Accept  json
// @Produce  json
// @Success 200 {object} Staff
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /staff [get]
// GetAllStaff is used to get all staff
func (h *Handler) GetAllStaff(c *gin.Context) {

	staff, err := h.services.GetAllStaff()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, staff)

	newSuccessResponse("getting all staff", "")
}

// @Summary Recognize Staff
// @Tags staff
// @Description recognize staff
// @ID find-staff
// @Accept  json
// @Produce  json
// @Param input body Staff
// @Success 200 {object} Staff
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /staff [post]
// RecognizeStaff is made for recognize staff
func (h *Handler) RecognizeStaff(c *gin.Context) {
	var meta map[string]any

	if err := c.BindJSON(&meta); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	staff, err := h.services.FindStaff(meta)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, staff)

	newSuccessResponse("finding staff", "")
}

// @Summary Find Staff
// @Tags staff
// @Description find staff by meta
// @ID find-staff
// @Accept  json
// @Produce  json
// @Param input body Staff
// @Success 200 {object} Staff
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /staff [post]
// FindStaff is made for find staff by meta
func (h *Handler) FindStaff(c *gin.Context) {
	var meta map[string]any

	if err := c.BindJSON(&meta); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	staff, err := h.services.FindStaff(meta)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, staff)

	newSuccessResponse("finding staff", "")
}

// @Summary Delete Staff
// @Tags staff
// @Description delete staff
// @ID delete-staff
// @Accept  json
// @Produce  json
// @Success 200 {object} StatusResponse
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /staff [delete]
// DeleteStaff allows to delete staff record
func (h *Handler) DeleteStaff(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.DeleteStaff(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Response{
		Message: "success",
	})

	newSuccessResponse("delete staff", "")
}
