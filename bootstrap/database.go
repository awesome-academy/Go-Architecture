package bootstrap

import (
	"Go-Architecture/domain/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func NewPostgresDatabase(env *Env) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		env.DBHost,
		env.DBUser,
		env.DBPass,
		env.DBName,
		env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Fail connect to DB %v", err)
	}

	if env.AppEnv == "development" {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	err = db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	if err != nil {
		return nil
	}
	err = db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	if err != nil {
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("Connected")

	err = db.AutoMigrate(&entity.User{}, &entity.Task{})
	if err != nil {
		return nil
	}

	DB = db
	return DB
}

func CloseDB(ormDb *gorm.DB) {
	db, err := ormDb.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to DB closed.")
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(db *gorm.DB) {
	nowTime := time.Now().Unix()
	if createTimeField := db.Statement.Schema.LookUpField("CreatedOn"); createTimeField != nil {
		if createTimeField.HasDefaultValue {
			err := createTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
			if err != nil {
				return
			}
		}
	}

	if modifyTimeField := db.Statement.Schema.LookUpField("ModifiedOn"); modifyTimeField != nil {
		if modifyTimeField.HasDefaultValue {
			err := modifyTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
			if err != nil {
				return
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if _, ok := db.Get("gorm:update_column"); !ok {
		db.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	}
}
