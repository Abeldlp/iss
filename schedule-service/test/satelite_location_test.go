package test

import (
	"testing"

	"github.com/Abeldlp/fullinfo/scheduler-service/config"
	"github.com/Abeldlp/fullinfo/scheduler-service/controller"
	"github.com/Abeldlp/fullinfo/scheduler-service/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func beforeEach() *gorm.DB {
	test_db := config.TestDBInit()
	test_db.AutoMigrate(&model.SateliteLocation{})

	return test_db
}

func TestSavingSateliteLocation(t *testing.T) {
	test_db := beforeEach()
	controller.SaveSateliteLocation()

	var locations []model.SateliteLocation

	test_db.Find(&locations)

	assert.Equal(t, 1, len(locations))

	config.TestDBFree(test_db)
}

func TestSaveTimezone(t *testing.T) {
	test_db := beforeEach()
	controller.SaveSateliteLocation()

	var locations []model.SateliteLocation

	test_db.Find(&locations)

	assert.NotEmpty(t, locations[0].Timezone)

	config.TestDBFree(test_db)
}

func TestConvertDateAndHour(t *testing.T) {
	test_db := beforeEach()
	controller.SaveSateliteLocation()

	var locations []model.SateliteLocation

	test_db.Find(&locations)

	assert.NotEmpty(t, locations[0].Date)
	assert.NotEmpty(t, locations[0].Hour)

	config.TestDBFree(test_db)
}
