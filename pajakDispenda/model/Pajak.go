package model

type Pajak struct {
	NoTransaksi   string `json:"no_transaksi" gorm:"primaryKey;column:no_transaksi"`
	NoBill        string `json:"no_bill" gorm:"column:no_bill"`
	TransDate     string `json:"trans_date" gorm:"column:trans_date"`
	Jam           string `json:"jam" gorm:"column:jam"`
	GrossSales    string `json:"gross_sales" gorm:"column:gross_sales"`
	SalesDiscount string `json:"sales_discount" gorm:"column:sales_discount"`
	ItemDiscount  string `json:"item_discount" gorm:"column:item_discount"`
	Service       string `json:"service" gorm:"column:service"`
	Tax           string `json:"tax" gorm:"column:tax"`
	SubTotal      string `json:"sub_total" gorm:"column:sub_total"`
	Rounding      string `json:"rounding" gorm:"column:rounding"`
	Status        string `json:"status" gorm:"column:status"`
	BranchID      string `json:"branch_id" gorm:"column:branch_id"`
	ValidStatus   string `json:"valid_status" gorm:"column:valid_status"`
	TaCharge      string `json:"tacharge" gorm:"column:tacharge"`
	NMHHB         string `json:"nmhhb" gorm:"column:nmhhb"`
}

func (Pajak) TableName() string {
	return "api_pajak"
}
