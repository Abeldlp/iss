package model

type User struct {
	Email    string `json:"email" gorm:"unique"`
	Timezone string `json:"timezone"`
}
