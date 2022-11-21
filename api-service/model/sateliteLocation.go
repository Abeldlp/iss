package model

import "gorm.io/gorm"

type SateliteLocation struct {
	gorm.Model
	Text string `json:"text"`
}
