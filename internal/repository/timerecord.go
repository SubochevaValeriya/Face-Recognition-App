package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
)

type TimeRecordDBMem struct {
	TimeRecords []models.TimeRecord                        `json:"time_records"`
	c           chan models.StreamModel[models.TimeRecord] `json:"-"`
	m           sync.Mutex                                 `json:"-"`
	path        string                                     `json:"-"`
}

func NewTimeRecordDBMem() *TimeRecordDBMem {
	return &TimeRecordDBMem{
		TimeRecords: []models.TimeRecord{},
		c:           make(chan models.StreamModel[models.TimeRecord], 10),
		m:           sync.Mutex{},
	}
}

func (db *TimeRecordDBMem) Stream(ctx context.Context) (chan models.StreamModel[models.TimeRecord], error) {
	return db.c, nil
}

func (db *TimeRecordDBMem) Add(timeRecord models.AddTimeRecord) (models.TimeRecord, error) {
	db.m.Lock()
	defer db.m.Unlock()
	id := len(db.TimeRecords) + 1
	record := timeRecord.ToTimeRecord(id)
	db.TimeRecords = append(db.TimeRecords, record)

	go func() {
		db.c <- models.StreamModel{Model: record, IsDeleted: false}

	}()
	return record, nil
}

func (db *TimeRecordDBMem) Update(timeRecord models.UpdateTimeRecord) (models.TimeRecord, error) {
	db.m.Lock()
	defer db.m.Unlock()
	for i, record := range db.TimeRecords {
		if record.Id == timeRecord.Id {
			db.TimeRecords[i] = timeRecord.ToUpdateTimeRecord(record.Id)
			go func() {
				db.c <- models.StreamModel{Model: db.TimeRecords[i], IsDeleted: false}
			}()
			return db.TimeRecords[i], nil
		}
	}
	return models.TimeRecord{}, fmt.Errorf("Time record not found")
}

func (db *TimeRecordDBMem) Delete(id int) error {
	db.m.Lock()
	defer db.m.Unlock()
	for i, record := range db.TimeRecords {
		if record.Id == id {
			db.TimeRecords = append(db.TimeRecords[:i], db.TimeRecords[i+1:]...)
			go func() {
				db.c <- models.StreamModel{Model: record, IsDeleted: true}
			}()
			return nil
		}
	}
	return fmt.Errorf("Time record not found")

}

func (db *TimeRecordDBMem) Get(id int) (models.TimeRecord, error) {
	db.m.Lock()
	defer db.m.Unlock()
	for _, record := range db.TimeRecords {
		if record.Id == id {
			return record, nil
		}
	}
	return models.TimeRecord{}, fmt.Errorf("Time record not found")
}

func (db *TimeRecordDBMem) All() ([]models.TimeRecord, error) {
	db.m.Lock()
	defer db.m.Unlock()
	// allRecords := make([]model.TimeRecord, 0, len(db.TimeRecords))
	// for _, rec := range db.TimeRecords {
	// 	allRecords = append(allRecords, rec)
	// }
	return db.TimeRecords, nil
}

func (db *TimeRecordDBMem) ByEmployeeId(id int) ([]models.TimeRecord, error) {
	db.m.Lock()
	defer db.m.Unlock()
	var records []models.TimeRecord
	for _, record := range db.TimeRecords {
		if *record.Employee == id {
			records = append(records, record)
		}
	}

	return records, nil
}

func (db *TimeRecordDBMem) ByDate(start time.Time, end time.Time, employeeId *int) ([]models.TimeRecord, error) {
	db.m.Lock()
	defer db.m.Unlock()
	var records []models.TimeRecord

	for _, record := range db.TimeRecords {
		entry := record.EntryTime.StartTime
		employee := record.Employee
		exit := record.ExitTime.EndTime
		if entry == start && exit == end && employee == employeeId {
			records = append(records, record)
		}
	}
	return records, nil
}

func (db *TimeRecordDBMem) LastByEmployeeId(id int) (models.TimeRecord, error) {
	db.m.Lock()
	defer db.m.Unlock()
	var result models.TimeRecord

	for _, record := range db.TimeRecords {
		if *record.Employee == id && record.ExitTime == nil {
			result = record
		}
	}
	if result.Id == 0 {
		return result, fmt.Errorf("No time records found for employee id %v", id)
	}
	return result, nil

}
