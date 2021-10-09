package models

import "encoding/json"

//type Product struct {
//	tableName      struct{} `pg:"products"`
//	Id             int64    `pg:"id,pk"`
//	VendorCode     string   `pg:"vendor_code"`
//	Name           string   `pg:"name"`
//	Category       string   `pg:"category"`
//	Type           string
//	Description    string   `pg:"description"`
//	StockBalance   int32    `pg:"stock_balance"`
//	Specifications json.RawMessage
//}

type Product struct {
	tableName      struct{} `pg:"products"`
	Id             int      `pg:"id,pk"`
	Name           string
	VendorCode     float64
	Price          float64
	StockBalance   int
	Description    string
	Specifications json.RawMessage
	CategoryId     int
	TypeId         int
}

type ProductCreateRequest struct {
	Name           string          `json:"name"`
	VendorCode     float64         `json:"vendor_code"`
	Price          float64         `json:"price"`
	StockBalance   int             `json:"stock_balance"`
	Description    string          `json:"description"`
	Specifications json.RawMessage `json:"specifications"`
	CategoryId     int             `json:"category_id"`
	TypeId         int             `json:"type_id"`
}
