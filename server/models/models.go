package models

import (
	"time"
)

type Item struct {
	Item_Id     string `json:"item_id"`
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_Id    string `json:"order_id"`
}

type Orders struct {
	Order_Id      string `json:"order_id"`
	Customer_Name string `json:"customer_name"`
	Ordered_At    time.Time
}

type Req_Data struct {
	Ordered_At    time.Time
	Customer_Name string `json:"customer_name"`
	Items         []struct {
		Item_Code   string `json:"item_code"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
		Item_Id     string `json:"item_id"`
	} `json:"items"`
}
