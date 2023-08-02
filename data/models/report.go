package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Report struct {
	ID           string
	AuthorID     string
	ReportedID   string
	Cause        string
	Type         string
	CreateDate   string
	ModifiedDate string
}

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{
		db: db,
	}
}

// Create a new report in the database
func (rr *ReportRepository) CreateReport(report *Report) error {
	_, err := rr.db.Exec("INSERT INTO report (id, authorID, reportedID, cause, type, createDate, modifiedDate) VALUES (?, ?, ?, ?, ?, ?, ?)",
		report.ID, report.AuthorID, report.ReportedID, report.Cause, report.Type, report.CreateDate, report.ModifiedDate)
	return err
}

// Get a report by ID from the database
func (rr *ReportRepository) GetReportByID(reportID string) (*Report, error) {
	var report Report
	row := rr.db.QueryRow("SELECT id, authorID, reportedID, cause, type, createDate, modifiedDate FROM report WHERE id = ?", reportID)
	err := row.Scan(&report.ID, &report.AuthorID, &report.ReportedID, &report.Cause, &report.Type, &report.CreateDate, &report.ModifiedDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Report not found
		}
		return nil, err
	}
	return &report, nil
}

// Update a report in the database
func (rr *ReportRepository) UpdateReport(report *Report) error {
	_, err := rr.db.Exec("UPDATE report SET authorID = ?, reportedID = ?, cause = ?, type = ?, createDate = ?, modifiedDate = ? WHERE id = ?",
		report.AuthorID, report.ReportedID, report.Cause, report.Type, report.CreateDate, report.ModifiedDate, report.ID)
	return err
}

// Delete a report from the database
func (rr *ReportRepository) DeleteReport(reportID string) error {
	_, err := rr.db.Exec("DELETE FROM report WHERE id = ?", reportID)
	return err
}
