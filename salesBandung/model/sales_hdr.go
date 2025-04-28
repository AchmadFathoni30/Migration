package model

import (
	"database/sql"
)

type Sales_hdr struct {
	Id 				sql.NullString			 `json:"id" gorm:"primaryKey;column:id"`
	Reset_no		sql.NullString			 `json:"reset_no" gorm:"column:reset_no"`
	Counter_no		sql.NullString			 `json:"counter_no" gorm:"column:counter_no"`
	Trx_id 			sql.NullString			 `json:"trx_id" gorm:"column:trx_id"`
	Trx_date 		sql.NullString			 `json:"trx_date" gorm:"column:trx_date"`
	Branch_id 		sql.NullString			 `json:"branch_id" gorm:"column:branch_id"`
	Client_id 		sql.NullString	 		 `json:"client_id" gorm:"column:client_id"`
	Session_id 		sql.NullString			 `json:"session_id" gorm:"column:session_id"`
	Pos_no 			sql.NullString	 		 `json:"pos_no" gorm:"column:pos_no"`
	Pos_id			sql.NullString			 `json:"pos_id" gorm:"column:pos_id"`
	Order_no 		sql.NullString	 		 `json:"order_no" gorm:"column:order_no"`
	Ent_date 		sql.NullString			 `json:"ent_date" gorm:"column:ent_date"`
	Ent_user 		sql.NullString			 `json:"ent_user" gorm:"column:ent_user"`
	Table_no 		sql.NullString	 		 `json:"table_no" gorm:"column:table_no"`
	Tax_rate 		sql.NullString			 `json:"tax_rate" gorm:"column:tax_rate"`
	Svc_charge_rate sql.NullString	 		 `json:"svc_charge_rate" gorm:"column:svc_charge_rate"`
	Sdisc_rate 		sql.NullString			 `json:"sdisc_rate" gorm:"column:sdisc_rate"`
	Guest 			sql.NullString			 `json:"guest" gorm:"column:guest"`
	Status 			sql.NullString			 `json:"status" gorm:"column:status"`
	Valid_status 	sql.NullString			 `json:"valid_status" gorm:"column:valid_status"`
	Valid_date 		sql.NullString			 `json:"valid_date" gorm:"column:valid_date"`
	Prn_rev_no 		sql.NullString			 `json:"prn_rev_no" gorm:"column:prn_rev_no"`
	Refund_vch 		sql.NullString			 `json:"refund_vch" gorm:"column:refund_vch"`
	Delivery_no 	sql.NullString	 		 `json:"delivery_no" gorm:"column:delivery_no"`
	Dp_status 		sql.NullString	 		 `json:"dp_status" gorm:"column:dp_status"`
	Kd_pemakai 		sql.NullString			 `json:"kd_pemakai" gorm:"column:kd_pemakai"`
	Href 			sql.NullString			 `json:"href" gorm:"column:href"`
	Trx_status 		sql.NullString	 		 `json:"trx_status" gorm:"column:trx_status"`
	Flag_omni 		sql.NullString	 		 `json:"flag_omni" gorm:"column:flag_omni"`
}

func (Sales_hdr) TableName() string {
	return "sales_hdr_bdg"
}