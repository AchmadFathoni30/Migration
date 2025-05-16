package controllers

import (
	"log"
	"net/http"

	"Migration/config"
	"Migration/salesBandung/model"

	"github.com/gin-gonic/gin"
)

func GetSalesHeader(c *gin.Context) {
	query := `
	SELECT 
		a.ID, 
		a.RESET_NO, 
		a.COUNTER_NO, 
		a.TRX_ID, 
		a.TRX_DATE, 
		a.BRANCH_ID, 
		a.CLIENT_ID, 
		a.SESSION_ID, 
		a.POS_NO,
        a.POS_ID, 
		a.ORDER_NO, 
		a.ENT_DATE, 
		a.ENT_USER, 
		a.TABLE_NO, 
		a.TAX_RATE, 
		a.SVC_CHARGE_RATE, 
		a.SDISC_RATE, 
		a.GUEST, 
		a.STATUS, 
		a.VALID_STATUS, 
		a.VALID_DATE, 
		a.PRN_REV_NO, 
		a.REFUND_VCH, 
		a.DELIVERY_NO, 
		a.DP_STATUS, 
		a.FLAG_OMNI, 
		a.KD_PEMAKAI, 
		a.HREF, 
		a.TRX_STATUS
	FROM POS_SALES_HDR a
	JOIN GHHB b ON a.BRANCH_ID = b.NOHHB
	WHERE a.STATUS = '1'
		AND a.TRX_DATE BETWEEN '2025-03-25' AND '2025-03-25'
		AND a.BRANCH_ID IN ('801','802','803','804','805','806','807','808','809','810','811','812','813','814','815','816','817','818','819','820','821','822','823',
        '824','825','826','827','828','829','830','831','833','834','835','836','837','838','839','840','841','842','843','844','845','846','847','848','849','850')
		AND a.VALID_STATUS='1'`

	rows, err := config.MSSQL.Query(query)
	if err != nil {
		log.Println("Error executing SQL Server query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query SQL Server"})
		return
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		var p model.Sales_hdr

		err := rows.Scan(
			&p.Id, &p.Reset_no, &p.Counter_no, &p.Trx_id, &p.Trx_date,
			&p.Branch_id, &p.Client_id, &p.Session_id, &p.Pos_no, &p.Pos_id,
			&p.Order_no, &p.Ent_date, &p.Ent_user, &p.Table_no, &p.Tax_rate,
			&p.Svc_charge_rate, &p.Sdisc_rate, &p.Guest, &p.Status, &p.Valid_status,
			&p.Valid_date, &p.Prn_rev_no, &p.Refund_vch, &p.Delivery_no, &p.Dp_status,
			&p.Flag_omni, &p.Kd_pemakai, &p.Href, &p.Trx_status,
		)

		if err != nil {
			log.Println("Rows scan error:", err)
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

func GetSalesDetail(c *gin.Context) {
    query := `
    SELECT
        a.SEQ_NO, 
        a.ID, 
        a.PRODCODE, 
        a.PRODMAIN, 
        a.QTY, 
        a.PRICE, 
        a.DISC_ITEM, 
        a.AMOUNT,
        a.ITEM_TYPE, 
        a.GROSS, 
        a.DISC_AMOUNT, 
        a.ID_INDEX, 
        a.PRODNAME, 
        a.PARENTID, 
        a.LINEDESC,
        a.PROMOID, 
        a.PROMODESC, 
        a.DISCREASON, 
        a.DISCTYPE, 
        a.DISCVALUE, 
        a.KET, 
        a.DISCYES,
        a.TRANSRC, 
        a.FLAG_OMNI, 
        b.BRANCH_ID, 
        b.TRX_DATE
    FROM POS_SALES_DTL a
    JOIN POS_SALES_HDR b ON a.ID = b.ID
    JOIN GHHB c ON b.BRANCH_ID = c.NOHHB
    WHERE b.STATUS = '1'
    AND b.TRX_DATE BETWEEN '2025-03-25' AND '2025-03-25'
    AND b.BRANCH_ID IN ('801','802','803','804','805','806','807','808','809','810','811','812','813','814','815','816','817','818','819','820','821','822','823',
        '824','825','826','827','828','829','830','831','833','834','835','836','837','838','839','840','841','842','843','844','845','846','847','848','849','850')
    AND b.VALID_STATUS='1'`

    rows, err := config.MSSQL.Query(query)
    if err != nil {
        log.Println("Error executing SQL Server query:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query SQL Server"})
        return
    }
    defer rows.Close()

    var count int
    for rows.Next() {

        var p model.Sales_detail

        err := rows.Scan(
            &p.Seq_no, &p.Id, &p.Prodcode, &p.Prodmain, &p.Qty,
            &p.Price, &p.Disc_item, &p.Amount, &p.Item_type,
            &p.Gross, &p.Disc_amount, &p.Id_index, &p.Prodname,
            &p.Parentid, &p.Linedesc, &p.Promoid, &p.Promodesc,
            &p.Discreason, &p.Disctype, &p.Discvalue, &p.Ket,
            &p.Discyes, &p.Transrc, &p.Flag_omni, &p.Branch_id, &p.Trx_date,
        )

        if err != nil {   
            log.Println("Rows scan error:", err)
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
        "ReecordsInserted": count,
    })
}

func GetSalesPayment(c *gin.Context) {
    query := `
    SELECT 
        a.ID, 
        a.GROSS_SALES, 
        a.GROSS, 
        a.SALES_DISCOUNT, 
        a.TOTAL_COST, 
        a.TAX, 
        a.SUB_TOTAL, 
        a.NET, 
        a.CASH, 
        a.NON_CASH,
        a.ROUNDING, 
        a.REFUND, 
        a.REFUND_VCH, 
        a.ITEM_DISCOUNT,
        a.ID_DP, 
        a.DP, 
        a.NET_DP,
        a.FBCA,
        a.CARD_AMT,
        a.CARD_NO,
        a.EDC_VDR,
        a.PROMO_DISCOUNT,
        a.NOTE,
        a.PROMOID,
        a.PROMODESC,
        a.DISCREASON,
        a.DISCTYPE,
        a.DISCVALUE,
        a.ITEM_TYPE,
        a.DISCMAX,
        a.GROSS_SALES_DISC,
        a.TRANSRC,
        a.TACHARGE,
        a.FLAG_OMNI,
        b.BRANCH_ID,
        b.TRX_DATE
    FROM POS_SALES_PAYMENT a
    JOIN POS_SALES_HDR b ON a.ID = b.ID
    JOIN GHHB c ON b.BRANCH_ID = c.NOHHB
    WHERE b.STATUS = '1'
    AND b.TRX_DATE BETWEEN '2025-03-25' AND '2025-04-28'
    AND b.BRANCH_ID IN ('801','802','803','804','805','806','807','808','809','810','811','812','813','814','815','816','817','818','819','820','821','822','823',
        '824','825','826','827','828','829','830','831','833','834','835','836','837','838','839','840','841','842','843','844','845','846','847','848','849','850')`

    rows, err := config.MSSQL.Query(query)
    if err != nil {
        log.Println("Error executing SQL Server query:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query SQL Server"})
        return
    }
    defer rows.Close()

    var count int
    for rows.Next() {
        var p model.Sales_payment

        err := rows.Scan(
            &p.Id, &p.Gross_sales, &p.Gross, &p.Sales_discount,
            &p.Total_cost, &p.Tax, &p.Sub_total, &p.Net, &p.Cash,
            &p.Non_cash, &p.Rounding, &p.Refund, &p.Refund_vch, &p.Item_discount,
            &p.Id_dp, &p.Dp, &p.Net_dp, &p.Fbca, &p.Card_amt,
            &p.Card_no, &p.Edc_vdr, &p.Promo_discount, &p.Note,
            &p.Promoid, &p.Promodesc, &p.Discreason, &p.Disctype,
            &p.Discvalue, &p.Item_type, &p.Discmax, &p.Gross_sales_disc,
            &p.Transrc, &p.Tacharge, &p.Flag_omni, &p.Branch_id, &p.Trx_date,
        )
        
        if err != nil {
            log.Println("Rows scan error:", err)
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


func GetSalesVoid(c *gin.Context) {
    query :=`
        SELECT 
            a.ID, 
            a.ENT_DATE,
            a.ENT_USER,
            a.REASON, 
            a.DESCRIPTION, 
            b.BRANCH_ID,
            b.TRX_DATE
        FROM POS_SALES_VOID a
        JOIN POS_SALES_HDR b ON a.ID = b.ID
        JOIN GHHB c ON b.BRANCH_ID = c.NOHHB
        WHERE b.STATUS = '1'
        AND b.TRX_DATE BETWEEN '2025-03-25' AND '2025-03-25'
        AND b.BRANCH_ID IN ('801','802','803','804','805','806','807','808','809','810','811','812','813','814','815','816','817','818','819','820','821','822','823',
        '824','825','826','827','828','829','830','831','833','834','835','836','837','838','839','840','841','842','843','844','845','846','847','848','849','850')`

    rows, err := config.MSSQL.Query(query)
    if err != nil {
        log.Println("Error executing SQL Server query")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query SQL Server"})
        return
    }
    defer rows.Close()

    var count int
    for rows.Next() {
        var a model.Sales_void

        err := rows.Scan(
            &a.Id, &a.Ent_date, &a.Ent_user, &a.Reason,
            &a.Description, &a.Branch_id, &a.Trx_date,
        )

        if err != nil {
            log.Println("Rows scan error:", err)
            continue
        }

        if err := config.DB.Create(&a).Error; err != nil {
            log.Println("Insert error:", err)
            continue
        }
        count++
    }

    c.JSON(http.StatusOK, gin.H{
        "message":      "Data migration completed",
        "recordsInserted": count,
    })
}