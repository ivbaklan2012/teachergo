package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Homework struct {
	ID   uuid.UUID
	Name string
	Date time.Time
	Body string
}

func CreateHomework(db *sql.DB, studentID uuid.UUID, homework *Homework) error {
	homework.ID = uuid.New()
	_, err := db.Exec("INSERT INTO homework (id, name, date, body, student_id) VALUES (?, ?, ?, ?, ?)", homework.ID.String(), homework.Name, homework.Date.Format(time.RFC3339), homework.Body, studentID.String())
	return err
}

func GetHomeworkByStudentID(db *sql.DB, studentID uuid.UUID) ([]*Homework, error) {
	rows, err := db.Query("SELECT id, name, date, body FROM homework WHERE student_id = ?", studentID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var homework []*Homework
	for rows.Next() {
		var hw Homework
		var date string
		if err := rows.Scan(&hw.ID, &hw.Name, &date, &hw.Body); err != nil {
			return nil, err
		}
		hw.Date, _ = time.Parse(time.RFC3339, date)
		homework = append(homework, &hw)
	}

	return homework, nil
}

func UpdateHomework(db *sql.DB, homework *Homework) error {
	_, err := db.Exec("UPDATE homework SET name = ?, date = ?, body = ? WHERE id = ?", homework.Name, homework.Date.Format(time.RFC3339), homework.Body, homework.ID.String())
	return err
}

func DeleteHomework(db *sql.DB, id uuid.UUID) error {
	_, err := db.Exec("DELETE FROM homework WHERE id = ?", id.String())
	return err
}
