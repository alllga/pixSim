package model

import (
	"time"

	valid "github.com/asaskevich/govalidator"
)

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID string						`json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}