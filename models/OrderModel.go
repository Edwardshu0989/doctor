package models

import (
	"time"
)

type Medical struct {
	Medical_id         int     `form:"Medical_id"			   json:"Medical_id"`
	Medical_name       string  `form:"Medical_name"   	   json:"Medical_name"`
	Medical_price      float32 `form:"Medical_price"   	   json:"Medical_price"`
	Medical_num        int     `form:"Medical_num"   	   json:"Medical_num"`
	Medical_days       int     `form:"Medical_days"   	   json:"Medical_days"`
	Medical_last_days  int     `form:"Medical_last_days"   	   json:"Medical_last_days"`
	Medical_is_status  uint    `form:"Medical_is_status"   	   json:"Medical_is_status"`
	Medical_instrument int     `form:"Medical_instrument"   	   json:"Medical_instrument"`
	Medical_point      int     `form:"Medical_point"   	   json:"Medical_point"`
	Medical_time       time.Time
}
