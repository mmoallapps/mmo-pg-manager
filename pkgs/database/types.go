package database

type JCase struct {
	CASE_ID            int     `json:"case_id" alias:"case_id"`
	ASSIGNED_TO        *string `json:"person_id" alias:"person_id"`
	PROVIDER_GRP_ID    *string `json:"provider_id" alias:"provider_id"`
	STATUS_DESCR       string  `json:"status" alias:"status"`
	RC_SUMMARY         *string `json:"summary" alias:"summary"`
	RC_DESCRLONG       *string `json:"description" alias:"description"`
	CONTACT_NAME       string  `json:"contact" alias:"contact"`
	RC_CONTACT_INFO    string  `json:"contactMethod" alias:"contactMethod"`
	ROW_ADDED_DTTM     string  `json:"openDate" alias:"openDate"`
	RC_PRIORITY        string  `json:"priority" alias:"priority"`
	RC_SEVERITY        string  `json:"severity" alias:"severity"`
	ROW_LASTMANT_DTTM  string  `json:"modified" alias:"modified"`
	JHA_OLNK_PARTITION *string `json:"prodPartition" alias:"prodPartition"`
	JHA_OLNK_BANK_NO   *string `json:"prodBnk" alias:"prodBnk"`
	JHA_OLNK_UAT_PARTI *string `json:"uatPartition" alias:"uatPartition"`
	JHA_OLNK_UATBNK_NO *string `json:"uatBnk" alias:"uatBnk"`
	COMPANYID          string  `json:"company_id" alias:"company_id"`
	CATEGORY_DESCR     *string `json:"category" alias:"category"`
	RC_TYPE_DESCR      *string `json:"type" alias:"type"`
	RC_DETAIL_DESCR    *string `json:"detail" alias:"detail"`
}

type MMOCase struct {
	case_id       int     `json:"case_id"`
	person_id     *string `json:"person_id"`
	provider_id   *string `json:"provider_id"`
	status        string  `json:"status"`
	summary       *string `json:"summary"`
	description   *string `json:"description"`
	contact       string  `json:"contact"`
	contactMethod string  `json:"contactMethod"`
	openDate      string  `json:"openDate"`
	priority      string  `json:"priority"`
	severity      string  `json:"severity"`
	modified      string  `json:"modified"`
	prodPartition *string `json:"prodPartition"`
	prodBnk       *string `json:"prodBnk"`
	uatPartition  *string `json:"uatPartition"`
	uatBnk        *string `json:"uatBnk"`
	company_id    string  `json:"company_id"`
	category      *string `json:"category"`
	case_type     *string `json:"type"`
	detail        *string `json:"detail"`
}
