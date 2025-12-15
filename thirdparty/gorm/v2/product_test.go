package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db?_loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	if err = db.AutoMigrate(&Product{}); err != nil {
		panic("failed to migrate schema")
	}
}

func teardown() {
	if err := os.Remove("test.db"); err != nil {
		panic("failed to remove test.db")
	}
}

func TestCreate(t *testing.T) {
	productItem := Product{Code: "D42", Price: 100}
	tx := db.Create(&productItem)
	assert.NoError(t, tx.Error)
	db.Delete(&productItem)
}

func TestSelect(t *testing.T) {
	productItem := Product{Code: "D42", Price: 100}
	tx := db.Create(&productItem)
	assert.NoError(t, tx.Error)

	var product Product
	tx = db.Where("price = ?", 200).First(&product)
	assert.ErrorIs(t, tx.Error, gorm.ErrRecordNotFound)

	tx = db.Model(&Product{}).Where("price = ?", 100).First(&product)
	assert.NoError(t, tx.Error)
	assert.Equal(t, productItem, product)
}
