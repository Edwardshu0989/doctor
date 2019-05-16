package models

import (
	"time"
)

//
type Shop struct {
	Shop_id             int    `form:"Shop_id"			   json:"Shop_id"`
	Shop_status         uint   `form:"Shop_status"   	   json:"Shop_status"`
	Shop_province       string `form:"Shop_province" 	   json:"Shop_province"`
	Shop_city           string `form:"Shop_city"      	   json:"Shop_city"`
	Shop_detail_address string `form:"Shop_detail_address" json:"Shop_detail_address"`
	Shop_coordinate     string `form:"Shop_coordinate"     json:"Shop_coordinate"`
	Shop_done           string `form:"Shop_done"           json:"Shop_done"`
	Shop_point          string `form:"Shop_point"          json:"Shop_point"`
	Shop_telphone       string `form:"Shop_telphone"       json:"Shop_telphone"`
	Shop_account        string `form:"Shop_account" 	   json:"Shop_account"`
	Shop_password       string `form:"Shop_password" 	   json:"Shop_password"`
	Shop_time           time.Time
}

func (s Shop) TableName() string {
	return "shop"
}

func (this *Shop) Add() (ID int, err error) {
	res := DB.Save(this)
	if res != nil {
		err = res.Error
	}
	ID = this.Shop_id
	return
}

func (this *Shop) Inquire() (shopinfo []Shop, err error) {
	res := DB.Find(&shopinfo)
	if res != nil {
		err = res.Error
	}
	return
}
