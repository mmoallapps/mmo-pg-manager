package jsource_views_interface

import (
	"database/sql"
	"fmt"
	"time"
)

// JNote represents a note record in the system
type JNote struct {
	CASE_ID        int       `json:"case_id"`
	NOTE_SEQ_NBR   int       `json:"noteSeq"`
	RC_SUMMARY     string    `json:"summary"`
	RC_NOTE_TYPE   string    `json:"type"`
	RC_VISIBILITY  string    `json:"visibility"`
	ROW_ADDED_DTTM time.Time `json:"created"`
	BO_NAME        string    `json:"addedBy"`
	RC_DESCRLONG   *string   `json:"note"`
}

func GetNotesByCases(cases []int) ([]JNote, error) {
	// Create a slice to hold the notes
	notes := []JNote{}
	// Create a string to hold the query
	query := `SELECT 
	CASE_ID, 
	NOTE_SEQ_NBR, 
	RC_SUMMARY, 
	RC_NOTE_TYPE, 
	RC_VISIBILITY, 
	ROW_ADDED_DTTM, 
	BO_NAME, 
	RC_DESCRLONG 
	FROM PS_JHA_CSNT_EXT_VW WHERE CASE_ID IN (`
	// Create a string to hold the case IDs
	caseIDs := ""
	// Loop through the cases and add them to the caseIDs string
	for i, c := range cases {
		caseIDs += fmt.Sprintf("%d", c)
		if i < len(cases)-1 {
			caseIDs += ","
		}
	}
	// Add the case IDs to the query
	query += caseIDs + `)`
	// Execute the query
	rows, err := jviewDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through the rows and add them to the notes slice
	for rows.Next() {
		var note JNote
		var nullNote sql.NullString

		if err := rows.Scan(&note.CASE_ID, &note.NOTE_SEQ_NBR, &note.RC_SUMMARY, &note.RC_NOTE_TYPE,
			&note.RC_VISIBILITY, &note.ROW_ADDED_DTTM, &note.BO_NAME, &nullNote); err != nil {
			return nil, err
		}

		// Convert sql.NullString to *string
		if nullNote.Valid {
			note.RC_DESCRLONG = &nullNote.String
		} else {
			note.RC_DESCRLONG = nil
		}

		notes = append(notes, note)
	}
	return notes, nil
}
