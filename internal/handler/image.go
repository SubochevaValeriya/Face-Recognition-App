package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UploadEndPoint(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")

	if err != nil {
		c.JSON(400, gin.H{"error": "Wrong image data"})
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	defer file.Close()

	image, err := h.services.UploadImageWithFace(file, header)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not save"})
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(200, image)
	newSuccessResponse("Face found and image saved", image.Path)
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
	id := c.Param("id")
	fileName, imgFile, err := h.services.GetImageAsFile(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "File was not found"})
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	header := make([]byte, 512)
	if _, err := imgFile.Read(header); err != nil {
		c.JSON(500, gin.H{"error": "Failed to read image header"})
		return
	}

	contentType := http.DetectContentType(header)
	c.Writer.Header().Set("Content-Type", contentType)
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	if _, err := io.Copy(c.Writer, imgFile); err != nil {
		c.JSON(500, gin.H{"error": "Failed to stream image"})
	}
	newSuccessResponse("Image found", fileName)
}

func (h *Handler) Recognize(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")

	if err != nil {
		c.JSON(400, gin.H{"error": "Wrong image data"})
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	defer file.Close()

	image, err := h.services.RecognizeImage(file, header)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not process image"})
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(200, image)
	newSuccessResponse("Found faces", image.Path)
}

func (h *Handler) SaveEndPoint(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")

	if err != nil {
		c.JSON(400, gin.H{"error": "Wrong image data"})
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	defer file.Close()

	image, err := h.services.SaveImage(file, header)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not save"})
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(200, image)
	newSuccessResponse("Image saved", image.Path)
}
