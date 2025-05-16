package controller

import (
	"log"
	"net/http"

	"Migration/config"
	"Migration/pajakDispenda/model"

	"github.com/gin-gonic/gin"
)

func MigrateData(c *gin.Context) {
	query := `
	SELECT
		a.ID AS NO_TRANSAKSI,
		a.ID AS NO_BILL,
		a.TRX_DATE,
		a.VALID_DATE,
		c.GROSS_SALES,
		c.SALES_DISCOUNT,
		c.ITEM_DISCOUNT,
		c.TOTAL_COST,
		c.TAX,
		c.SUB_TOTAL,
		c.ROUNDING,
		a.BRANCH_ID,
		a.STATUS,
		a.VALID_STATUS,
		c.TACHARGE,
		b.NMHHB
	FROM POS_SALES_HDR a
	JOIN POS_SALES_PAYMENT c ON a.ID = c.ID
	JOIN GHHB b ON a.BRANCH_ID = b.NOHHB
	WHERE a.STATUS = '1' AND a.TRX_DATE BETWEEN '2025-04-09' AND '2025-05-09' AND a.VALID_STATUS = '1'`

	rows, err := config.MSSQL.Query(query)
	if err != nil {
		log.Println("Error executing SQL Server query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query SQL Server"})
		return
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		var (
			p         model.Pajak
		)

		err := rows.Scan(
			&p.NoTransaksi, &p.NoBill, &p.TransDate, &p.Jam,
			&p.GrossSales, &p.SalesDiscount, &p.ItemDiscount,
			&p.Service, &p.Tax, &p.SubTotal, &p.Rounding,
			&p.BranchID, &p.Status, &p.ValidStatus, &p.TaCharge, &p.NMHHB,
		)

		if err != nil {
			log.Println("Row scan error:", err)
			continue
		}

		if err := config.DB.Create(&p).Error; err != nil {
			log.Println("Insert error:", err)
			continue
		}
		count++
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Data migration completed",
		"recordsInserted": count,
	})
}
