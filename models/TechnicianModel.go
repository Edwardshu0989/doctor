package models

import (
	"time"
)

//
type Technician struct {
	Technician_id       int    `form:"Technician_id"			   json:"Technician_id"`
	Technician_name     string `form:"Technician_name"   	   json:"Technician_name"`
	Technician_job      string `form:"Technician_job"   	   json:"Technician_job"`
	Technician_phone    int    `form:"Technician_phone"   	   json:"Technician_phone"`
	Technician_shop     string `form:"Technician_shop"   	   json:"Technician_shop"`
	Technician_account  string `form:"Technician_account"   	   json:"Technician_account"`
	Technician_password string `form:"Technician_password"   	   json:"Technician_password"`
	Technician_time     time.Time
}

func (s Technician) TableName() string {
	return "technician"
}

func (this *Technician) Add() (ID int, err error) {
	res := DB.Save(this)
	if res != nil {
		err = res.Error
	}
	ID = this.Technician_id
	return
}

func (this *Technician) Inquire() (technician []Technician, err error) {
	res := DB.Find(&technician)
	if res != nil {
		err = res.Error
	}
	return
}
