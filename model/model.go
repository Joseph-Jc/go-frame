package model

import (
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	zone       = "Asia/Shanghai"
)

type Model struct {
	ID        uint    `gorm:"primary_key" json:"id"`
	CreatedAt string  `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt string  `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt *string `gorm:"type:timestamp" sql:"index" json:"deleted_at"`
}

func (m Model) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("created_at", time.Now().Format(timeFormat)); err != nil {
		return err
	}
	if err := scope.SetColumn("updated_at", time.Now().Format(timeFormat)); err != nil {
		return err
	}
	return nil
}

func (m Model) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("updated_at", time.Now().Format(timeFormat)); err != nil {
		return err
	}
	return nil
}

type Time time.Time

// UnmarshalJSON implements json unmarshal interface.
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

// MarshalJSON implements json marshal interface.
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t Time) local() time.Time {
	loc, _ := time.LoadLocation(zone)
	return time.Time(t).In(loc)
}

// Value ...
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time 注意是指针类型 method
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
