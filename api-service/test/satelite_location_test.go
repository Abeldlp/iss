package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Abeldlp/fullinfo/api-service/config"
	"github.com/Abeldlp/fullinfo/api-service/model"
	"github.com/Abeldlp/fullinfo/api-service/route"
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

func TestGetAll(t *testing.T) {
	test_db = config.TestDBInit()
	test_db.AutoMigrate(&model.SateliteLocation{})
	sateliteLocationMocker(2)

	r := config.InitializeServer()
	route.AppendSateliteLocationRoute(r)

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	t.Log(string(responseData))

	config.TestDBFree()
}

// func TestMain(m *testing.M) {
// 	test_db = config.TestDBInit()
// 	test_db.AutoMigrate(&model.SateliteLocation{})
// 	exitVal := m.Run()
// 	// config.TestDBFree(test_db)
// 	os.Exit(exitVal)
// }
