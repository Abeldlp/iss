package main

import (
	"fmt"

	"github.com/Abeldlp/fullinfo/scheduler-service/model"
	"github.com/Abeldlp/fullinfo/scheduler-service/util"
)

func GetSateliteLocationData() {
	info := &model.SateliteLocation{}

	util.GetJson("https://api.wheretheiss.at/v1/satellites/25544", &info)

	fmt.Println(info)
}

func main() {
	GetSateliteLocationData()
}
