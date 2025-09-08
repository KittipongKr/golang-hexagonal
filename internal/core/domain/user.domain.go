package domain

import (
	m "csat-servay/internal/adapter/mongo/model"
)

type User struct {
	m.BaseModel `json:",inline"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Status      string `json:"status"`
}
