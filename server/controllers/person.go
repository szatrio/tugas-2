package controllers

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"tugas-2/server/models"
	"tugas-2/server/views"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderControllers struct {
	db *gorm.DB
}

func NewPersonController(db *gorm.DB) *OrderControllers {
	return &OrderControllers{
		db: db,
	}
}

func randomString(min int, max int) string {

	rand.Seed(time.Now().UnixNano())
	ord_id := strconv.Itoa(rand.Intn(max-min+1) + min)

	return ord_id
}

func (p *OrderControllers) POST_Orders(ctx *gin.Context) {

	rs := "OrderID_" + randomString(1000000000000, 2000000000000)
	it_id := strconv.Itoa(1)

	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)

	println(body_string)

	var key_data models.Req_Data

	err := json.Unmarshal([]byte(body_string), &key_data)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_ORDER_FAILURE",
			Error:   err.Error(),
		})
		return
	}

	p.db.Create(&models.Orders{
		Order_Id:      rs,
		Customer_Name: key_data.Customer_Name})

	p.db.Create(&models.Item{
		Item_Id:     it_id,
		Item_Code:   key_data.Items[0].Item_Code,
		Description: key_data.Items[0].Description,
		Quantity:    key_data.Items[0].Quantity,
		Order_Id:    rs})

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_ORDER_SUCCESS",
		Payload: key_data,
	})
}

func (p *OrderControllers) GET_Orders(ctx *gin.Context) {

	var item_data []models.Item
	err := p.db.Find(&item_data).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: item_data,
	})
}

func (p *OrderControllers) PUT_Orders(ctx *gin.Context) {

	param := ctx.Param("order_id")
	println(param)

	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)

	println(body_string)

	var key_data models.Req_Data

	err := json.Unmarshal([]byte(body_string), &key_data)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_ORDER_FAILURE",
			Error:   err.Error(),
		})
		return
	}

	p.db.Model(&models.Orders{
		Order_Id:      param,
		Customer_Name: key_data.Customer_Name,
		Ordered_At:    time.Time{},
	}).Where("order_id = ?", param).Update("customer_name", key_data.Customer_Name)

	p.db.Model(&models.Item{
		Item_Id:     key_data.Items[0].Item_Id,
		Item_Code:   key_data.Items[0].Item_Code,
		Description: key_data.Items[0].Description,
		Quantity:    key_data.Items[0].Quantity,
		Order_Id:    param}).Where("order_id = ?", param).Updates(models.Item{
		Item_Id:     key_data.Items[0].Item_Id,
		Item_Code:   key_data.Items[0].Item_Code,
		Description: key_data.Items[0].Description,
		Quantity:    key_data.Items[0].Quantity,
		Order_Id:    param,
	})

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_ORDER_SUCCESS",
		Payload: key_data,
	})
}

func (p *OrderControllers) DELETE_Orders(ctx *gin.Context) {
	param := ctx.Param("order_id")
	println(param)

	find := p.db.First(&models.Orders{}, "order_id = ?", param).Error

	if find != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "ORDER ID ALREADY DELETED",
			Error:   find.Error(),
		})
		return
	}

	delete_order := p.db.Where("order_id = ?", param).Delete(&models.Orders{}).Error
	delete_item := p.db.Where("order_id = ?", param).Delete(&models.Item{}).Error

	if delete_item != nil && delete_order != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_ORDER_FAILED",
			Error:   delete_order.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "DELETE ORDER SUCCESSFULL",
		Payload: "Order_ID Has Deleted",
	})
}
