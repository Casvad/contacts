package models

import (
	"contacts/utils/date"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type User struct {
	Id        int64             `json:"id" gorm:"primaryKey"`
	Email     string            `gorm:"column:email" json:"email"`
	Password  string            `gorm:"column:password" json:"-"`
	Name      string            `gorm:"column:name" json:"name"`
	CreatedAt date.ContactDate  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt date.ContactDate  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *date.ContactDate `gorm:"column:deleted_at" json:"deleted_at"`
}

type UserContact struct {
	Id          int64             `json:"id" gorm:"primaryKey"`
	UserId      int64             `gorm:"column:user_id" json:"user_id"`
	Email       string            `gorm:"column:email" json:"email"`
	Name        string            `gorm:"column:name" json:"name"`
	PrefixPhone string            `gorm:"column:prefix_phone" json:"prefix_phone"`
	Phone       string            `gorm:"column:phone" json:"phone"`
	CreatedAt   date.ContactDate  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   date.ContactDate  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *date.ContactDate `gorm:"column:deleted_at" json:"deleted_at"`
}

type UserContactModification struct {
	Id            int64            `json:"id" gorm:"primaryKey"`
	UserContactId int64            `gorm:"column:user_contact_id" json:"user_contact_id"`
	Modification  UserContactMod   `gorm:"column:modification" json:"modification"`
	CreatedAt     date.ContactDate `gorm:"column:created_at" json:"created_at"`
}

type UserContactMod UserContact

func (c UserContactMod) Value() (driver.Value, error) {
	return json.Marshal(c)
}
func (c *UserContactMod) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("UserContactMod type assertion to []byte failed")
	}
	return json.Unmarshal(b, &c)
}
