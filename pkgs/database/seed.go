package database

import (
	"fmt"
	jview "mmoallapps/mmo-pg-manager/pkgs/jsource_views_interface"
)

// SeedCases seeds the jcases table with data from the jsource view

func SeedCases() {
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
		casesToInsert[i] = JcasetoMMOCase(caseData)
	}
	// insert cases individually to avoid errors (testing)
	for _, caseData := range casesToInsert {
		// insert the case into the database
		InsertCase(caseData)
	}
	fmt.Println("Cases seeded successfully")

	fmt.Println("Seeding case notes...")
	// get the case ids from the cases to insert
	caseIDs := make([]int, len(casesToInsert))
	for i, caseData := range casesToInsert {
		caseIDs[i] = caseData.case_id
	}

	// Process notes in chunks of 500 cases at a time
	chunkSize := 500
	for i := 0; i < len(caseIDs); i += chunkSize {
		end := i + chunkSize
		if end > len(caseIDs) {
			end = len(caseIDs)
		}

		// Get the chunk of case IDs
		caseIDsChunk := caseIDs[i:end]

		// Get notes for the current chunk of cases
		notes, err := jview.GetNotesByCases(caseIDsChunk)
		if err != nil {
			fmt.Println("Error getting notes for cases", caseIDsChunk, ":", err)
			continue
		}

		// Check if the notes are empty for this chunk
		if len(notes) == 0 {
			fmt.Println("No notes to insert for cases", caseIDsChunk)
			continue
		}

		// Insert all notes for this chunk
		InsertNotes(notes)
		fmt.Printf("Processed notes for cases %d to %d\n", i+1, end)
	}

	fmt.Println("Case notes seeded successfully")
	fmt.Println("Seeding completed successfully")
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
	// update notes
	// get the case ids from the cases to update
	caseIDs := make([]int, len(casesToUpdate))
	for i, caseData := range casesToUpdate {
		caseIDs[i] = caseData.case_id
	}
	// get the notes for the cases
	notes, err := jview.GetNotesByCases(caseIDs)
	if err != nil {
		fmt.Println("Error getting notes:", err)
		return
	}
	// check if the notes are empty
	if len(notes) == 0 {
		fmt.Println("No notes to update")
		return
	}
	// insert the notes into the database
	InsertNotes(notes)
	// delete cases that are closed
	DeleteCasesClosedToday()
}

func DeleteCasesClosedToday() {
	closedCases, err := jview.GetCasesClosedToday()
	if err != nil {
		fmt.Println("Error getting closed cases:", err)
		return
	}
	if len(closedCases) == 0 {
		fmt.Println("No cases closed today to delete.")
		return
	}
	caseIDs := make([]int, len(closedCases))
	for i, c := range closedCases {
		caseIDs[i] = c.CASE_ID
	}
	DeleteCases(caseIDs)
	fmt.Printf("Deleted %d cases closed today from the database.\n", len(caseIDs))
}
