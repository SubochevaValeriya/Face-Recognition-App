package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UploadEndPoint(c *gin.Context) {

}

func (h *Handler) DataEndPoint(c *gin.Context) {
	id := c.Param("id")
	image, err := h.services.GetImage(id)

	if image.ID == 0 {
		c.JSON(404, gin.H{"error": "Image not found"})
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(200, image)
		newSuccessResponse("Image found", image.Path)
	}
}

func (h *Handler) FileEndPoint(c *gin.Context) {

}

func (h *Handler) Recognize(c *gin.Context) {

}

func (h *Handler) SaveEndPoint(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	defer file.Close()

	if err != nil {
		c.JSON(400, gin.H{"error": "Wrong image data"})
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	image, err := h.services.SaveImage(file, header)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not save"})
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(200, image)
	newSuccessResponse("Image found", image.Path)
}
