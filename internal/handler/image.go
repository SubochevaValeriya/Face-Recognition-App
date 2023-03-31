package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Image struct {
	ID   int   json:"id"
	Path string json:"path"
}
