package handler

import (
	"github.com/SubochevaValeriya/face-recognition-app/internal/middleware"
	"github.com/SubochevaValeriya/face-recognition-app/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", h.Register)
			user.POST("/login", h.Login)
		}

		protected := api.Group("/admin")
		{
			protected.Use(middleware.JwtAuthMiddleware())
			protected.GET("/user", h.CurrentUser)
		}

		staff := api.Group("/staff")
		{
			staff.POST("/add", h.AddStaff)
			staff.PUT("/update", h.UpdateStaff)
			staff.DELETE("/delete", h.DeleteStaff)
			staff.GET("/get", h.GetStaff)
			staff.GET("/all", h.GetAllStaff)
			staff.POST("/recognize", h.RecognizeStaff)
			staff.POST("/find", h.FindStaff)
		}

		// image := api.Group("/image")
		// {
		// 	image.POST("/upload", h.UploadEndPoint)
		// 	image.GET("/data", h.DataEndPoint)
		// 	image.GET("/file", h.FileEndPoint)
		// 	image.POST("/recognize", h.Recognize)
		// 	image.POST("/save", h.SaveEndPoint)
		// }

		// timeRecord := api.Group("/timerecord")
		// {
		// 	timeRecord.POST("/add", h.AddTimeRecord)
		// 	timeRecord.PUT("/update", h.UpdateTimeRecord)
		// 	timeRecord.DELETE("/delete", h.DeleteTimeRecord)
		// 	timeRecord.GET("/get", h.GetTimeRecord)
		// 	timeRecord.GET("/all", h.AllTimeRecords)
		// 	timeRecord.GET("/byemployee", h.TimeRecordsByEmployee)
		// 	timeRecord.POST("/bydate", h.TimeRecordsByDate)
		// 	timeRecord.GET("/lastbyemployee", h.TimeRecordLastByEmployee)
		// }

		// thirdparty := api.Group("/thirdparty")
		// {
		// 	thirdparty.GET("/timerecordStream", h.TimerecordStream)
		// 	thirdparty.POST("/add", h.AddThirdparty)
		// 	thirdparty.DELETE("/delete", h.DeleteThirdparty)
		// 	thirdparty.GET("/all", h.AllThirdparty)
		// 	thirdparty.POST("/check", h.ChekThirdparty)
		// }
	}

	return router
}
