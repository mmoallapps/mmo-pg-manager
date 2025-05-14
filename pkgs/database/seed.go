package database

import (
	"fmt"
	jview "mmoallapps/mmo-pg-manager/pkgs/jsource_views_interface"
)

// SeedCases seeds the jcases table with data from the jsource view

func SeedCases() {
	// clear the jcases table
	ClearTable("jCase")
	newCases, err := jview.GetAllOpenCases()
	if err != nil {
		fmt.Println("Error getting cases:", err)
		return
	}

	// check if the cases are empty
	if len(newCases) == 0 {
		fmt.Println("No cases to insert")
		return
	}

	// print the number of cases to be inserted
	fmt.Println("Number of cases to be inserted:", len(newCases))
	// convert the cases to the database format
	casesToInsert := make([]MMOCase, len(newCases))
	for i, caseData := range newCases {
		// convert the case data to the database format
		casesToInsert[i] = JcasetoMMOCase(caseData)
	}
	// insert the cases into the database
	InsertCases(casesToInsert)
	// insert cases individually to avoid errors
	// for _, caseData := range casesToInsert {
	// 	// insert the case into the database
	// 	InsertCase(caseData)
	// }
	fmt.Println("Cases seeded successfully")
	// print the number of cases inserted
	fmt.Println("Number of cases inserted:", len(casesToInsert))
}

func UpdateDBCases() {
	cases, err := jview.GetModifiedCases()
	if err != nil {
		fmt.Println("Error getting cases:", err)
		return
	}
	if len(cases) == 0 {
		fmt.Println("No cases to update")
		return
	}
	// print the number of cases to be updated
	fmt.Println("Number of cases to be updated:", len(cases))
	// convert the cases to the database format
	casesToUpdate := make([]MMOCase, len(cases))
	for i, caseData := range cases {
		// convert the case data to the database format
		casesToUpdate[i] = JcasetoMMOCase(caseData)
	}
	// update the cases in the database
	UpdateCases(casesToUpdate)

}
