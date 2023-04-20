package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"github.com/gin-gonic/gin"
)

func (a *Handler) AddTimeRecordEndPoint(c *gin.Context) {
	var addTime models.AddTimeRecord
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
	}
	err = json.Unmarshal(buf, &addTime)
	if err != nil {
		log.Fatalf("unmarshal error: %s\n", err.Error())
	}
	timeRecord, err := a.db.Add(addTime)
	if err != nil {
	}
	c.JSON(200, timeRecord)
}

func (a *Handler) UpdateTimeRecordEndPoint(c *gin.Context) {
	var updateTime models.UpdateTimeRecord
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

	}
	err = json.Unmarshal(buf, &updateTime)
	if err != nil {
		log.Fatalf("unmarshal error: %s\n", err.Error())
	}
	timeRecord, err := a.db.Update(updateTime)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

	}
	c.JSON(200, timeRecord)

}

func (a *Handler) DeleteTimeRecordEndPoint(c *gin.Context) {
	idReq := c.Query("id")
	id, err := strconv.Atoi(idReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	errorD := a.db.Delete(id)
	if errorD != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{})
}

func (a *Handler) GetTimeRecordEndPoint(c *gin.Context) {
	idReq := c.Query("id")
	id, err := strconv.Atoi(idReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	timeRecords, err := a.db.Get(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, timeRecords)
}

func (a *Handler) AllTimeRecordsEndPoint(c *gin.Context) {
	timeRecords, err := a.db.All()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, timeRecords)
}

func (a *Handler) TimeRecordsByEmployeeEndPoint(c *gin.Context) {
	idReq := c.Query("id")
	id, err := strconv.Atoi(idReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	timeRecords, err := a.db.ByEmployeeId(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, timeRecords)
}

type TimeRecordsByDateRequest struct {
	EmployeeId *int           `json:"employee_id"`
	Start      model.DateTime `json:"start"`
	End        model.DateTime `json:"end"`
}

func (a *Handler) TimeRecordsByDateEndPoint(c *gin.Context) {
	var req TimeRecordsByDateRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	timeRecords, err := a.db.ByDate(req.Start.Time(), req.End.Time(), req.EmployeeId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, timeRecords)
}

func (a *Handler) TimeRecordLastByEmployeeEndPoint(c *gin.Context) {
	idReq := c.Query("id")
	id, err := strconv.Atoi(idReq)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	timeRecord, err := a.db.LastByEmployeeId(id)
	if err != nil {
		c.JSON(204, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, timeRecord)
}
