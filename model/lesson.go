package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID   uuid.UUID
	Name string
	Date time.Time
	Body string
}

func CreateLesson(db *sql.DB, studentID uuid.UUID, lesson *Lesson) error {
	lesson.ID = uuid.New()
	_, err := db.Exec("INSERT INTO lessons (id, name, date, body, student_id) VALUES (?, ?, ?, ?, ?)", lesson.ID.String(), lesson.Name, lesson.Date.Format(time.RFC3339), lesson.Body, studentID.String())
	return err
}

func GetLessonsByStudentID(db *sql.DB, studentID uuid.UUID) ([]*Lesson, error) {
	rows, err := db.Query("SELECT id, name, date, body FROM lessons WHERE student_id = ?", studentID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []*Lesson
	for rows.Next() {
		var lesson Lesson
		var date string
		if err := rows.Scan(&lesson.ID, &lesson.Name, &date, &lesson.Body); err != nil {
			return nil, err
		}
		lesson.Date, _ = time.Parse(time.RFC3339, date)
		lessons = append(lessons, &lesson)
	}

	return lessons, nil
}

func UpdateLesson(db *sql.DB, lesson *Lesson) error {
	_, err := db.Exec("UPDATE lessons SET name = ?, date = ?, body = ? WHERE id = ?", lesson.Name, lesson.Date.Format(time.RFC3339), lesson.Body, lesson.ID.String())
	return err
}

func DeleteLesson(db *sql.DB, id uuid.UUID) error {
	_, err := db.Exec("DELETE FROM lessons WHERE id = ?", id.String())
	return err
}
