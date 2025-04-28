package model

import(
	"database/sql"
)

type Sales_detail struct {
	Seq_no sql.NullString `json:"seq_no" gorm:"column:seq_no"`
	Id sql.NullString `json:"id" gorm:"column:id"`
	Prodcode sql.NullString `json:"prodcode" gorm:"column:prodcode"`
	Prodmain sql.NullString `json:"prodmain" gorm:"column:prodmain"`
	Qty sql.NullString `json:"qty" gorm:"column:qty"`
	Price sql.NullString `json:"price" gorm:"column:price"`
	Disc_item sql.NullString `json:"disc_item" gorm:"column:disc_item"`
	Amount sql.NullString `json:"amount" gorm:"column:amount"`
	Item_type sql.NullString `json:"item_type" gorm:"column:item_type"`
	Gross sql.NullString `json:"gross" gorm:"column:gross"`
	Disc_amount sql.NullString `json:"disc_amount" gorm:"column:disc_amount"`
	Id_index sql.NullString `json:"id_index" gorm:"column:id_index"`
	Prodname sql.NullString `json:"prodname" gorm:"column:prodname"`
	Parentid sql.NullString `json:"parentid" gorm:"column:parentid"`
	Linedesc sql.NullString `json:"linedesc" gorm:"column:linedesc"`
	Promoid sql.NullString `json:"promoid" gorm:"column:promoid"`
	Promodesc sql.NullString `json:"promodesc" gorm:"column:promodesc"`
	Discreason sql.NullString `json:"discreason" gorm:"column:discreason"`
	Disctype sql.NullString `json:"disctype" gorm:"column:disctype"`
	Discvalue sql.NullString `json:"discvalue" gorm:"column:discvalue"`
	Ket sql.NullString `json:"ket" gorm:"column:ket"`
	Discyes sql.NullString `json:"discyes" gorm:"column:discyes"`
	Transrc sql.NullString `json:"transrc" gorm:"column:transrc"`
	Flag_omni sql.NullString `json:"flag_omni" gorm:"column:flag_omni"`
	Branch_id sql.NullString `json:"branch_id" gorm:"column:branch_id"`
	Trx_date sql.NullString `json:"trx_date" gorm:"column:trx_date"`
}

func (Sales_detail) TableName() string {
	return "sales_dtl_bdg"
}