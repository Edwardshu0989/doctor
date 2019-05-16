package models

import (
	"fmt"
	"strconv"
	"time"
)

//
type Users struct {
	User_id            int    `form:"User_id"			   json:"User_id"`
	User_nickname      string `form:"User_nickname"   	   json:"User_nickname"`
	User_telphone      string `form:"User_telphone"   	   json:"User_telphone"`
	User_urand         string `form:"User_urand"   	   json:"User_urand"`
	User_sex           string `form:"User_sex"   	   json:"User_sex"`
	User_province      string `form:"User_province"   	   json:"User_province"`
	User_city          string `form:"User_city"   	   json:"User_city"`
	User_profession    string `form:"User_profession"   	   json:"User_profession"`
	User_signature     string `form:"User_signature"   	   json:"User_signature"`
	User_status        string `form:"User_status"   	   json:"User_status"`
	User_start_weight  int    `form:"User_start_weight"   	   json:"User_start_weight"`
	User_belong_stores int    `form:"User_belong_stores"   	   json:"User_belong_stores"`
	User_introducer    string `form:"User_introducer"   	   json:"User_introducer"`
	User_remark        string `form:"User_remark"   	   json:"User_remark"`
	User_reg_time      time.Time
}

func (u Users) TableName() string {
	return "user"
}

type Usermember struct {
	Usermember_mid              int       `form:"Usermember_mid"			   json:"Usermember_mid"`
	Usermember_uid              int       `form:"Usermember_uid"			   json:"Usermember_uid"`
	Usermember_card_no          int64     `form:"Usermember_card_no"   	   json:"Usermember_card_no"`
	Usermember_head_photo       string    `form:"Usermember_head_photo"   	   json:"Usermember_head_photo"`
	Usermember_aesthetic_weight float64   `form:"Usermember_aesthetic_weight"   	   json:"Usermember_aesthetic_weight"`
	Usermember_standard_weight  float64   `form:"Usermember_standard_weight"   	   json:"Usermember_standard_weight"`
	Usermember_start_weight     float64   `form:"Usermember_start_weight"   	   json:"Usermember_start_weight"`
	Usermember_lose_time        time.Time `form:"Usermember_lose_time"   	   json:"Usermember_lose_time"`
	Usermember_card_face        string    `form:"Usermember_card_face"   	   json:"Usermember_card_face"`
	Usermember_handle_time      int64     `form:"Usermember_handle_time"   	   json:"Usermember_handle_time"`
}

func (u Usermember) TableName() string {
	return "usermember"
}
func (this *Users) Add() (ID int, err error) {
	res := DB.Save(this)
	if res.Error != nil {
		err = res.Error
	}
	var usermemberinfo Usermember

	usermemberinfo = Usermember{Usermember_uid: this.User_id, Usermember_card_no: time.Now().Unix(), Usermember_aesthetic_weight: float64(float64(this.User_start_weight) - float64(6.02)), Usermember_standard_weight: float64(float64(this.User_start_weight) - float64(3.08)), Usermember_start_weight: float64(this.User_start_weight), Usermember_lose_time: time.Now(), Usermember_card_face: strconv.FormatInt(time.Now().Unix(), 10) + "hello", Usermember_handle_time: time.Now().Unix()}

	fmt.Println(usermemberinfo)
	ress := DB.Save(&usermemberinfo)
	fmt.Println(ress)
	if ress.Error != nil {
		err = res.Error
		DB.Delete(this)
		return
	}

	ID = this.User_id
	return
}

// func (this *Users) Inquire() (shopinfo []Users, err error) {
// 	res := DB.Find(&shopinfo)
// 	if res != nil {
// 		err = res.Error
// 	}
// 	return
// }
