package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

const (
	sqlDbName = "../resuminute-db"
)

type dbError struct {
	DbName string
	Err    string
}

type ResuminuteDB struct {
	dbContext *gorm.DB
}

func (e *dbError) Error() string {
	return fmt.Sprintf("%s - %s\n", e.Err, e.DbName)
}

func CreateDBSession() (*ResuminuteDB, error) {
	var err error
	db, err := gorm.Open("sqlite3", sqlDbName)
	if err != nil {
		return nil, &dbError{
			DbName: sqlDbName,
			Err:    err.Error(),
		}
	}
	if db == nil {
		fmt.Printf("error connection to sqlite3 instance\n")
		return nil, &dbError{
			DbName: sqlDbName,
			Err:    err.Error(),
		}
	}
	createDBTables(db)
	return &ResuminuteDB{dbContext: db}, nil
}

func createDBTables(db *gorm.DB) error {
	db.AutoMigrate(&Video{})
	return nil
}

func (db *ResuminuteDB) CreateVideo(video Video) (*uuid.UUID, error) {
	id := uuid.Must(uuid.NewV4())
	video.ID = id.String()
	db.dbContext.Create(&video)
	return &id, nil
}
