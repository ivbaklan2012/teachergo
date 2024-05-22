package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Student struct {
	ID       uuid.UUID
	Name     string
	Lessons  []*Lesson
	Homework []*Homework
}

func CreateStudent(db *sql.DB, student *Student) error {
	student.ID = uuid.New()
	_, err := db.Exec("INSERT INTO students (id, name) VALUES (?, ?)", student.ID.String(), student.Name)
	return err
}

func GetStudentByID(db *sql.DB, id uuid.UUID) (*Student, error) {
	row := db.QueryRow("SELECT id, name FROM students WHERE id = ?", id.String())

	var student Student
	err := row.Scan(&student.ID, &student.Name)
	if err != nil {
		return nil, err
	}

	student.Lessons, err = GetLessonsByStudentID(db, student.ID)
	if err != nil {
		return nil, err
	}

	student.Homework, err = GetHomeworkByStudentID(db, student.ID)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func UpdateStudent(db *sql.DB, student *Student) error {
	_, err := db.Exec("UPDATE students SET name = ? WHERE id = ?", student.Name, student.ID.String())
	return err
}

func DeleteStudent(db *sql.DB, id uuid.UUID) error {
	_, err := db.Exec("DELETE FROM students WHERE id = ?", id.String())
	return err
}
