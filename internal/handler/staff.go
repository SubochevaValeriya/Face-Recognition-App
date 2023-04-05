package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Staff struct {
	ID      int                    `json:"id"`
	Name    string                 `json:"name"`
	PhotoId int                    `json:"photo_id"`
	Meta    map[string]interface{} `json:"meta"`
}
