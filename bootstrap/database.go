package bootstrap

import (
	"Go-Architecture/domain/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
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
