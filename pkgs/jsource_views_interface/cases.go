package jsource_views_interface

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

var ConnectionString = "server=mmopsftcrmdbp03.jhacorp.com;user id=jFastExport;password=jFastjSource4354!;port=1433;database=CRM92PRD;ApplicationIntent=ReadOnly;"

var jviewDB *sql.DB

func init() {
	var err error
	jviewDB, err = sql.Open("mssql", ConnectionString)
	if err != nil {
		panic(err)
	}
}

type JCase struct {
	CASE_ID            int     `json:"case_id"`
	ASSIGNED_TO        *string `json:"person_id"`
	PROVIDER_GRP_ID    *string `json:"provider_id"`
	STATUS_DESCR       string  `json:"status"`
	RC_SUMMARY         *string `json:"summary"`
	RC_DESCRLONG       *string `json:"description"`
	CONTACT_NAME       string  `json:"contact"`
	RC_CONTACT_INFO    string  `json:"contactMethod"`
	ROW_ADDED_DTTM     string  `json:"openDate"`
	RC_PRIORITY        string  `json:"priority"`
	RC_SEVERITY        string  `json:"severity"`
	ROW_LASTMANT_DTTM  string  `json:"modified"`
	JHA_OLNK_PARTITION *string `json:"prodPartition"`
	JHA_OLNK_BANK_NO   *string `json:"prodBnk"`
	JHA_OLNK_UAT_PARTI *string `json:"uatPartition"`
	JHA_OLNK_UATBNK_NO *string `json:"uatBnk"`
	COMPANYID          string  `json:"company_id"`
	CATEGORY_DESCR     *string `json:"category"`
	RC_TYPE_DESCR      *string `json:"type"`
	RC_DETAIL_DESCR    *string `json:"detail"`
}

var baseQuery = `Select
CASE_ID,
ISNULL(ASSIGNED_TO, '1') AS ASSIGNED_TO,
ISNULL(PROVIDER_GRP_ID, '0') AS PROVIDER_GRP_ID,
STATUS_DESCR,
RC_SUMMARY,
RC_DESCRLONG,
CONTACT_NAME,
RC_CONTACT_INFO,
ROW_ADDED_DTTM,
RC_PRIORITY,
RC_SEVERITY,
ROW_LASTMANT_DTTM,
JHA_OLNK_PARTITION,
JHA_OLNK_BANK_NO,
JHA_OLNK_UAT_PARTI,
JHA_OLNK_UATBNK_NO,
COMPANYID,
CATEGORY_DESCR,
RC_TYPE_DESCR,
RC_DETAIL_DESCR
From ( SELECT *, ROW_NUMBER() OVER (PARTITION BY CASE_ID ORDER BY JHA_OLNK_PARTITION DESC) AS rn
FROM PS_JHA_CASE_EXT_VW WHERE STATUS_DESCR NOT IN ('Canceled', 'Closed', 'Resolved')) AS subquery WHERE rn = 1`

func readCases(rows *sql.Rows) ([]JCase, error) {
	// create a slice to hold the cases
	var cases []JCase
	// iterate over the rows and append the cases to the slice
	for rows.Next() {
		var jcase JCase
		if err := rows.Scan(
			&jcase.CASE_ID,
			&jcase.ASSIGNED_TO,
			&jcase.PROVIDER_GRP_ID,
			&jcase.STATUS_DESCR,
			&jcase.RC_SUMMARY,
			&jcase.RC_DESCRLONG,
			&jcase.CONTACT_NAME,
			&jcase.RC_CONTACT_INFO,
			&jcase.ROW_ADDED_DTTM,
			&jcase.RC_PRIORITY,
			&jcase.RC_SEVERITY,
			&jcase.ROW_LASTMANT_DTTM,
			&jcase.JHA_OLNK_PARTITION,
			&jcase.JHA_OLNK_BANK_NO,
			&jcase.JHA_OLNK_UAT_PARTI,
			&jcase.JHA_OLNK_UATBNK_NO,
			&jcase.COMPANYID,
			&jcase.CATEGORY_DESCR,
			&jcase.RC_TYPE_DESCR,
			&jcase.RC_DETAIL_DESCR); err != nil {
			return nil, err
		}
		cases = append(cases, jcase)
	}
	return cases, nil
}

// get all cases not Canceled Closed or Resolved
func GetAllOpenCases() ([]JCase, error) {
	// connect to the database
	db, err := sql.Open("mssql", ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	fmt.Println(("Getting Cases from Jview..."))
	rows, err := db.Query(baseQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cases, err := readCases(rows)
	if err != nil {
		return nil, err
	}
	return cases, nil
}

func GetModifiedCases() ([]JCase, error) {
	jviewDB, err := sql.Open("mssql", ConnectionString)
	if err != nil {
		return nil, err
	}
	defer jviewDB.Close()
	fmt.Println(("Getting Modified Cases from Jview..."))
	rows, err := jviewDB.Query(baseQuery + " AND ROW_LASTMANT_DTTM > DATEADD(MINUTE, -10, GETDATE())")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cases, err := readCases(rows)
	if err != nil {
		return nil, err
	}
	return cases, nil
}

// GetCasesClosedToday returns cases with status 'Closed' and closed today
func GetCasesClosedToday() ([]JCase, error) {
	db, err := sql.Open("mssql", ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Adjust the field name for close date if different in your schema
	query := `
        SELECT
            CASE_ID,
            ASSIGNED_TO,
            PROVIDER_GRP_ID,
            STATUS_DESCR,
            RC_SUMMARY,
            RC_DESCRLONG,
            CONTACT_NAME,
            RC_CONTACT_INFO,
            ROW_ADDED_DTTM,
            RC_PRIORITY,
            RC_SEVERITY,
            ROW_LASTMANT_DTTM,
            JHA_OLNK_PARTITION,
            JHA_OLNK_BANK_NO,
            JHA_OLNK_UAT_PARTI,
            JHA_OLNK_UATBNK_NO,
            COMPANYID,
            CATEGORY_DESCR,
            RC_TYPE_DESCR,
            RC_DETAIL_DESCR
        FROM PS_JHA_CASE_EXT_VW
        WHERE STATUS_DESCR in ('Closed', 'Canceled')
        AND CAST(ROW_LASTMANT_DTTM AS DATE) = CAST(GETDATE() AS DATE)
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readCases(rows)
}
