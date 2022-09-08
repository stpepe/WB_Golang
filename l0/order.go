package testapp

import (
	"github.com/go-playground/validator/v10"
)

type Order struct {
	OrderUid          string   `json:"order_uid" validate:"required"`         
	TrackNumber       string   `json:"track_number" validate:"required"`      
	Entry             string   `json:"entry" validate:"required"`             
	Delivery          Delivery `json:"delivery" validate:"required,dive,required"`          
	Payment           Payment  `json:"payment" validate:"required,dive,required"`           
	Items             []Item   `json:"items" validate:"required,dive,required"`             
	Locale            string   `json:"locale" validate:"required"`            
	InternalSignature string   `json:"internal_signature"`
	CustomerID        string   `json:"customer_id" validate:"required"`       
	DeliveryService   string   `json:"delivery_service" validate:"required"`  
	Shardkey          string   `json:"shardkey" validate:"required"`          
	SmID              int64    `json:"sm_id" validate:"gte=0"`             
	DateCreated       string   `json:"date_created" validate:"required"`      
	OofShard          string   `json:"oof_shard" validate:"required"`         
  }
  
  type Delivery struct {
	Name    string `json:"name" validate:"required"`   
	Phone   string `json:"phone" validate:"required"`  
	Zip     string `json:"zip" validate:"required"`    
	City    string `json:"city" validate:"required"`   
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"` 
	Email   string `json:"email" validate:"required"`  
  }
  
  type Item struct {
	ChrtID      int64  `json:"chrt_id" validate:"gte=0"`     
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int64  `json:"price" validate:"gte=0"`       
	Rid         string `json:"rid" validate:"required"`         
	Name        string `json:"name" validate:"required"`        
	Sale        int64  `json:"sale" validate:"gte=0"`        
	Size        string `json:"size" validate:"required"`        
	TotalPrice  int64  `json:"total_price" validate:"gte=0"` 
	NmID        int64  `json:"nm_id" validate:"gte=0"`       
	Brand       string `json:"brand" validate:"required"`       
	Status      int64  `json:"status" validate:"gte=0"`      
  }
  
  type Payment struct {
	Transaction  string `json:"transaction" validate:"required"`  
	RequestID    string `json:"request_id"`   
	Currency     string `json:"currency" validate:"required"`     
	Provider     string `json:"provider" validate:"required"`     
	Amount       int64  `json:"amount" validate:"gte=0"`       
	PaymentDt    int64  `json:"payment_dt" validate:"gte=0"`   
	Bank         string `json:"bank" validate:"required"`         
	DeliveryCost int64  `json:"delivery_cost" validate:"gte=0"`
	GoodsTotal   int64  `json:"goods_total" validate:"gte=0"`  
	CustomFee    int64  `json:"custom_fee" validate:"gte=0"`   
  }
  
func (order *Order) Validate() (error){
	validate := validator.New()
	return validate.Struct(order)
}