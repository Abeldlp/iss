package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Abeldlp/fullinfo/api-service/config"
	"github.com/Abeldlp/fullinfo/api-service/model"
	"github.com/Abeldlp/fullinfo/api-service/route"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
)

var test_db *gorm.DB

func newSateliteLocation() model.SateliteLocation {
	return model.SateliteLocation{
		Name:       "iss",
		Latitude:   111.11111,
		Longitude:  222.22222,
		Altitude:   333.3333,
		Velocity:   444.44444,
		Visibility: "555.55555",
		Footprint:  666.66666,
		Timestamp:  time.Now().Unix(),
		Daynum:     777.7777,
		SolarLat:   111.1111,
		SolarLon:   222.2222,
		Units:      "kilometers",
		Timezone:   "Europe/Amsterdam",
		Date:       "26-08-1991",
		Hour:       "15",
		DateISO:    time.Now().Local(),
	}
}

func sateliteLocationMocker(n int64) []model.SateliteLocation {
	var offset int64
	test_db.Model(&model.SateliteLocation{}).Count(&offset)
	var ret []model.SateliteLocation
	for i := offset + 1; i <= offset+n; i++ {
		sateliteLocation := newSateliteLocation()
		test_db.Create(&sateliteLocation)
		ret = append(ret, sateliteLocation)
	}
	return ret
}

func beforeEach() (*gorm.DB, []model.SateliteLocation, *gin.Engine) {
	test_db = config.TestDBInit()
	test_db.AutoMigrate(&model.SateliteLocation{})
	mockedLocations := sateliteLocationMocker(1)

	r := config.InitializeServer()
	route.AppendSateliteLocationRoute(r)

	return test_db, mockedLocations, r
}

func TestGetAll(t *testing.T) {

	test_db, mockedLocations, r := beforeEach()

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var locations []model.SateliteLocation

	json.NewDecoder(w.Body).Decode(&locations)

	assert.Equal(t, mockedLocations, locations)

	config.TestDBFree(test_db)
}

func TestGetAggregated(t *testing.T) {
	test_db, _, r := beforeEach()

	req, _ := http.NewRequest("GET", "/grouped", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var locations []model.AggregatedResult

	json.NewDecoder(w.Body).Decode(&locations)

	expected := []model.AggregatedResult{
		{
			Timezone: "Europe/Amsterdam",
			Date:     "26-08-1991",
			Hour:     "15",
			Minutes:  1,
		},
	}
	assert.Equal(t, expected, locations)

	config.TestDBFree(test_db)
}

func TestFilterWithoutResults(t *testing.T) {
	test_db, _, r := beforeEach()

	req, _ := http.NewRequest("GET", "/grouped?timezone[]=Europe/Berlin", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var locations []model.AggregatedResult

	json.NewDecoder(w.Body).Decode(&locations)

	expected := []model.AggregatedResult{}
	assert.Equal(t, expected, locations)

	config.TestDBFree(test_db)
}
