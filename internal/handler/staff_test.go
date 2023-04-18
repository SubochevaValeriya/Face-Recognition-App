package handler

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/service"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestHandler_AddStaff(t *testing.T) {
	type fields struct {
		services *service.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "", fields: fields{services: }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
			}
			h.AddStaff(tt.args.c)
		})
	}
}

func TestHandler_DeleteStaff(t *testing.T) {
	type fields struct {
		services *service.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
			}
			h.DeleteStaff(tt.args.c)
		})
	}
}

func TestHandler_FindStaff(t *testing.T) {
	type fields struct {
		services *service.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
			}
			h.FindStaff(tt.args.c)
		})
	}
}

func TestHandler_GetAllStaff(t *testing.T) {
	type fields struct {
		services *service.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
			}
			h.GetAllStaff(tt.args.c)
		})
	}
}

func TestHandler_GetStaff(t *testing.T) {
	type fields struct {
		services *service.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
			}
			h.GetStaff(tt.args.c)
		})
	}
}

func TestHandler_RecognizeStaff(t *testing.T) {
	type fields struct {
		services *service.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
			}
			h.RecognizeStaff(tt.args.c)
		})
	}
}

func TestHandler_UpdateStaff(t *testing.T) {
	type fields struct {
		services *service.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services: tt.fields.services,
			}
			h.UpdateStaff(tt.args.c)
		})
	}
}
