package database

import (
	"fmt"
	"mmoallapps/mmo-pg-manager/pkgs/jsource_views_interface"
	"strconv"
	"strings"
)

var fields = `
	"case_id",
	"person_id",
	"provider_id",
	"status",
	"summary",
	"description",
	"contact",
	"contactMethod",
	"openDate",
	"priority",
	"severity",
	"modified",
	"prodPartition",
	"prodBnk",
	"uatPartition",
	"uatBnk",
	"company_id",
	"category",
	"type",
	"detail"
`

func InsertCase(caseData MMOCase) {
	// connect to the database

	query := `INSERT INTO "jCase" (`
	query += fields
	query += `) VALUES (`
	query += fmt.Sprintf(`%d,`, caseData.case_id)
	query += fmt.Sprintf(`%d,`, personIDToInt(caseData.person_id))
	query += fmt.Sprintf(`'%s',`, safeString(caseData.provider_id))
	query += fmt.Sprintf(`'%s',`, caseData.status)
	query += fmt.Sprintf(`'%s',`, safeString(caseData.summary))
	query += fmt.Sprintf(`'%s',`, safeString(caseData.description))
	query += fmt.Sprintf(`'%s',`, safeString(&caseData.contact))
	query += fmt.Sprintf(`'%s',`, caseData.contactMethod)
	query += fmt.Sprintf(`'%s',`, caseData.openDate)
	query += fmt.Sprintf(`'%s',`, caseData.priority)
	query += fmt.Sprintf(`'%s',`, caseData.severity)
	query += fmt.Sprintf(`'%s',`, caseData.modified)
	query += fmt.Sprintf(`'%s',`, safeString(caseData.prodPartition))
	query += fmt.Sprintf(`'%s',`, safeString(caseData.prodBnk))
	query += fmt.Sprintf(`'%s',`, safeString(caseData.uatPartition))
	query += fmt.Sprintf(`'%s',`, safeString(caseData.uatBnk))
	query += fmt.Sprintf(`'%s',`, caseData.company_id)
	query += fmt.Sprintf(`'%s',`, safeString(caseData.category))
	query += fmt.Sprintf(`'%s',`, safeString(caseData.case_type))
	query += fmt.Sprintf(`'%s'`, safeString(caseData.detail))
	query += `);`
	_, err := Db.Exec(query)
	if err != nil {
		fmt.Println("Error inserting case:", err, "company_id:", caseData.company_id, "case_id:", caseData.case_id)
	}
}

func InsertCases(cases []MMOCase) {
	fmt.Println("Inserting cases...:", len(cases))
	chunkSize := 10000 // Insert in chunks to avoid large queries
	for i := 0; i < len(cases); i += chunkSize {
		end := i + chunkSize
		if end > len(cases) {
			end = len(cases)
		}

		// Start building the query
		query := `INSERT INTO "jCase" (`
		query += fields
		query += `) VALUES `

		// Add rows of values
		values := []string{}
		for j := i; j < end; j++ {
			row := fmt.Sprintf(`(
                %d,
                '%d',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s'
            )`,
				cases[j].case_id,
				personIDToInt(cases[j].person_id),
				safeString(cases[j].provider_id),
				cases[j].status,
				safeString(cases[j].summary),
				safeString(cases[j].description),
				safeString(&cases[j].contact),
				cases[j].contactMethod,
				cases[j].openDate,
				cases[j].priority,
				cases[j].severity,
				cases[j].modified,
				safeString(cases[j].prodPartition),
				safeString(cases[j].prodBnk),
				safeString(cases[j].uatPartition),
				safeString(cases[j].uatBnk),
				cases[j].company_id,
				safeString(cases[j].category),
				safeString(cases[j].case_type),
				safeString(cases[j].detail),
			)
			values = append(values, row)
		}

		// Join all rows with commas
		query += strings.Join(values, ",")
		query += `;`

		// Execute the queryF
		_, err = Db.Exec(query)
		if err != nil {
			fmt.Println("Error inserting cases:", err)
		}
		fmt.Println("Inserted", len(values), "cases")
	}
}

// safeString handles nil pointers and returns a default value (empty string) if the pointer is nil.
func safeString(s *string) string {
	if s == nil {
		return ""
	}
	// we can get all sorts of characters in the string, so we need to escape it
	str := *s
	str = strings.ReplaceAll(str, "'", "''")
	str = strings.ReplaceAll(str, "\"", "\\\"")
	str = strings.ReplaceAll(str, "\\", "\\\\")
	str = strings.ReplaceAll(str, "\n", "\\n")
	str = strings.ReplaceAll(str, "\r", "\\r")

	return str
}

func safeNumber(n *int) int {
	if n == nil {
		return 0
	}
	return *n
}

func personIDToInt(personID *string) int {
	if personID == nil {
		return 0
	}
	id, err := strconv.Atoi(*personID)
	if err != nil {
		return 0
	}
	return id
}

func JcasetoMMOCase(jcase jsource_views_interface.JCase) MMOCase {
	return MMOCase{
		case_id:       jcase.CASE_ID,
		person_id:     jcase.ASSIGNED_TO,
		provider_id:   jcase.PROVIDER_GRP_ID,
		status:        jcase.STATUS_DESCR,
		summary:       jcase.RC_SUMMARY,
		description:   jcase.RC_DESCRLONG,
		contact:       jcase.CONTACT_NAME,
		contactMethod: jcase.RC_CONTACT_INFO,
		openDate:      jcase.ROW_ADDED_DTTM,
		priority:      jcase.RC_PRIORITY,
		severity:      jcase.RC_SEVERITY,
		modified:      jcase.ROW_LASTMANT_DTTM,
		prodPartition: jcase.JHA_OLNK_PARTITION,
		prodBnk:       jcase.JHA_OLNK_BANK_NO,
		uatPartition:  jcase.JHA_OLNK_UAT_PARTI,
		uatBnk:        jcase.JHA_OLNK_UATBNK_NO,
		company_id:    jcase.COMPANYID,
		category:      jcase.CATEGORY_DESCR,
		case_type:     jcase.RC_TYPE_DESCR,
		detail:        jcase.RC_DETAIL_DESCR,
	}
}

func UpdateCases(cases []MMOCase) {
	fmt.Println("Updating cases...:", len(cases))
	chunkSize := 10000 // Update in chunks to avoid large queries
	for i := 0; i < len(cases); i += chunkSize {
		end := i + chunkSize
		if end > len(cases) {
			end = len(cases)
		}

		// Start building the query
		query := `INSERT INTO "jCase" (`
		query += fields
		query += `) VALUES `

		// Add rows of values
		values := []string{}
		for j := i; j < end; j++ {
			row := fmt.Sprintf(`(
                %d,
                '%d',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s',
                '%s'
            )`,
				cases[j].case_id,
				personIDToInt(cases[j].person_id),
				safeString(cases[j].provider_id),
				cases[j].status,
				safeString(cases[j].summary),
				safeString(cases[j].description),
				safeString(&cases[j].contact),
				cases[j].contactMethod,
				cases[j].openDate,
				cases[j].priority,
				cases[j].severity,
				cases[j].modified,
				safeString(cases[j].prodPartition),
				safeString(cases[j].prodBnk),
				safeString(cases[j].uatPartition),
				safeString(cases[j].uatBnk),
				cases[j].company_id,
				safeString(cases[j].category),
				safeString(cases[j].case_type),
				safeString(cases[j].detail),
			)
			values = append(values, row)
		}

		// Join all rows with commas
		query += strings.Join(values, ",")
		query += ` ON CONFLICT (case_id) DO UPDATE SET`
		query += ` person_id = EXCLUDED.person_id,`
		query += ` provider_id = EXCLUDED.provider_id,`
		query += ` status = EXCLUDED.status,`
		query += ` summary = EXCLUDED.summary,`
		query += ` description = EXCLUDED.description,`
		query += ` contact = EXCLUDED.contact,`
		query += ` "contactMethod" = EXCLUDED."contactMethod",`
		query += ` "openDate" = EXCLUDED."openDate",`
		query += ` priority = EXCLUDED.priority,`
		query += ` severity = EXCLUDED.severity,`
		query += ` modified = EXCLUDED.modified,`
		query += ` "prodPartition" = EXCLUDED."prodPartition",`
		query += ` "prodBnk" = EXCLUDED."prodBnk",`
		query += ` "uatPartition" = EXCLUDED."uatPartition",`
		query += ` "uatBnk" = EXCLUDED."uatBnk",`
		query += ` company_id = EXCLUDED.company_id,`
		query += ` category = EXCLUDED.category,`
		query += ` type = EXCLUDED.type,`
		query += ` detail = EXCLUDED.detail`
		query += `;`
		_, err := Db.Exec(query)
		if err != nil {
			fmt.Printf("Query: %s\n", query)
			fmt.Println("Error updating case:", err)

		}
		// Execute the query

		fmt.Println("Inserted", len(values), "cases")
	}
}
