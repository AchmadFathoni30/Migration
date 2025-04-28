package model

import (
	"database/sql"
)

type Sales_void struct {
	Id			 sql.NullString `json:"id" gorm:"primaryKer;column:id"`
	Ent_date	 sql.NullString `json:"ent_date" gorm:"column:ent_date"`
	Ent_user	 sql.NullString `json:"ent_user" gorm:"column:ent_user"`
	Reason		 sql.NullString `json:"reason" gorm:"column:reason"`
	Description	 sql.NullString `json:"description" gorm:"column:description"`
	Branch_id	 sql.NullString `json:"branch_id" gorm:"column:branch_id"`
	Trx_date	 sql.NullString `json:"trx_date" gorm:"column:trx_date"`
}

func (Sales_void) TableName() string {
	return "sales_void_bdg"
}