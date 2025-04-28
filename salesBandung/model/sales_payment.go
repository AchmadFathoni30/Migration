package model

import (
	"database/sql"
)

type Sales_payment struct {
	Id 					sql.NullString `json:"id" gorm:"primaryKey;column:id"`
	Gross_sales 		sql.NullString `json:"gross_sales" gorm:"column:gross_sales"`
	Gross 				sql.NullString `json:"gross" gorm:"column:gross"`
	Sales_discount 		sql.NullString `json:"sales_discount" gorm:"column:sales_discount"`
	Total_cost 			sql.NullString `json:"total_cost" gorm:"column:total_cost"`
	Tax 				sql.NullString `json:"tax" gorm:"column:tax"`
	Sub_total 			sql.NullString `json:"sub_total" gorm:"column:sub_total"`
	Net 				sql.NullString `json:"net" gorm:"column:net"`
	Cash 				sql.NullString `json:"cash" gorm:"column:cash"`
	Non_cash 			sql.NullString `json:"non_cash" gorm:"column:non_cash"`
	Rounding 			sql.NullString `json:"rounding" gorm:"column:rounding"`
	Refund 				sql.NullString `json:"refund" gorm:"column:refund"`
	Refund_vch 			sql.NullString `json:"refund_vch" gorm:"column:refund_vch"`
	Item_discount 		sql.NullString `json:"item_discount" gorm:"column:item_discount"`
	Id_dp 				sql.NullString `json:"id_dp" gorm:"column:id_dp"`
	Dp 					sql.NullString `json:"dp" gorm:"column:dp"`
	Net_dp 				sql.NullString `json:"net_dp" gorm:"column:net_dp"`
	Fbca 				sql.NullString `json:"fbca" gorm:"column:fbca"`
	Card_amt 			sql.NullString `json:"card_amt" gorm:"column:card_amt"`
	Card_no 			sql.NullString `json:"card_no" gorm:"column:card_no"`
	Edc_vdr 			sql.NullString `json:"edc_vdr" gorm:"column:edc_vdr"`
	Promo_discount 		sql.NullString `json:"promo_discount" gorm:"column:promo_discount"`
	Note 				sql.NullString `json:"note" gorm:"column:note"`
	Promoid 			sql.NullString `json:"promoid" gorm:"column:promoid"`
	Promodesc 			sql.NullString `json:"promodesc" gorm:"column:promodesc"`
	Discreason 			sql.NullString `json:"discreason" gorm:"column:discreason"`
	Disctype 			sql.NullString `json:"disctype" gorm:"column:disctype"`
	Discvalue 			sql.NullString `json:"discvalue" gorm:"column:discvalue"`
	Item_type 			sql.NullString `json:"item_type" gorm:"column:item_type"`
	Discmax 			sql.NullString `json:"discmax" gorm:"column:discmax"`
	Gross_sales_disc 	sql.NullString `json:"gross_sales_disc" gorm:"column:gross_sales_disc"`
	Transrc 			sql.NullString `json:"transrc" gorm:"column:transrc"`
	Tacharge 			sql.NullString `json:"tacharge" gorm:"column:tacharge"`
	Flag_omni 			sql.NullString `json:"flag_omni" gorm:"column:flag_omni"`
	Branch_id 			sql.NullString `json:"branch_id" gorm:"column:branch_id"`
	Trx_date 			sql.NullString `json:"trx_date" gorm:"column:trx_date"`
}

func (Sales_payment) TableName() string {
	return "sales_payment_bdg"
}