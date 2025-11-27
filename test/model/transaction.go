package model 

import "time"


type Transaction struct {
	TransactionId string    `json:"transactionId" gorm:"column:transaction_id;primaryKey"` 
	OrderId       string    `json:"orderId" gorm:"column:order_id"`                     
	Title         string    `json:"title" gorm:"column:title"`                           
	Description   string    `json:"description" gorm:"column:description"`             
	Value         int       `json:"value" gorm:"column:value"`                           
	Type          string    `json:"type" gorm:"column:type"`                             // EARNED / USED
	CreatedTime   time.Time `json:"createdTime" gorm:"column:created_time"`              
}

