package models

import (
	"database/sql"
	svccourse "english/backend/api/proto/svc-module"
	"english/backend/shared/errorlog"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"net/http"
)

type ListOfCourses struct {
	UserID          int64
	CourseData      []CourseData
	ProtoCourseData []*svccourse.ListOfCoursesResponse_CourseData
}

type CourseData struct {
	CourseID    int64         `json:"courseID,omitempty" db:"id"`
	Name        string        `json:"name,omitempty" db:"name"`
	Language    string        `json:"language,omitempty" db:"language"`
	Hidden      bool          `json:"hidden,omitempty" db:"hidden"`
	Description string        `json:"description,omitempty" db:"description"`
	Price       int64         `json:"price,omitempty" db:"price"`
	MainImage   string        `json:"mainImage,omitempty" db:"main_img"`
	MainVideo   string        `json:"mainVideo,omitempty" db:"main_video"`
	CreatedDate string        `json:"createdDate,omitempty" db:"created_date"`
	EditingDate string        `json:"editingDate,omitempty" db:"editing_date"`
	SubStudents pq.Int64Array `json:"subStudents" db:"sub_students"`
}

func (l *ListOfCourses) GetListOfCourses(db *sqlx.DB) (int, error) {
	if err := db.Select(&l.CourseData, `
	SELECT
		id, name, language, main_img, main_video, description, hidden, price, created_date, editing_date, sub_students
	FROM
	    course
	WHERE 
	    user_id=$1`, l.UserID); err != nil {
		return http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	l.ProtoCourseData = make([]*svccourse.ListOfCoursesResponse_CourseData, len(l.CourseData))
	for i, datum := range l.CourseData {
		courseData := &svccourse.ListOfCoursesResponse_CourseData{}
		if err := copier.Copy(courseData, datum); err != nil {
			return http.StatusInternalServerError, err
		}

		l.ProtoCourseData[i] = courseData
	}

	return http.StatusOK, nil
}

func SelectCourseByID(db *sqlx.DB, courseId int64) (*svccourse.CourseDataResponse, error) {
	data := &svccourse.CourseDataResponse{
		Status: http.StatusOK,
	}
	courseDataCopy := CourseData{}

	if err := db.QueryRow(`
	SELECT
	    name, language, main_img, main_video, description, hidden, price, sub_students
	FROM
	    course
	WHERE
	    id=$1`, courseId).Scan(&courseDataCopy.Name, &courseDataCopy.Language, &courseDataCopy.MainImage,
		&courseDataCopy.MainVideo, &courseDataCopy.Description, &courseDataCopy.Hidden,
		&courseDataCopy.Price, &courseDataCopy.SubStudents); err != nil {
		data.Status = http.StatusInternalServerError
		data.Error = "Неизвестная ошибка"
		return data, errorlog.NewError(err.Error())
	}

	course := svccourse.CreateCourseRequest{}
	if err := copier.Copy(&course, courseDataCopy); err != nil {
		return data, errorlog.NewError(err.Error())
	}

	data.CourseData = &course

	rows, err := db.Query(`
	SELECT
	   id, name, exercise, task_description, lesson_number, deadline_date
	FROM
	   lesson
	WHERE
	   course_id=$1`, courseId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		data.Status = http.StatusInternalServerError
		data.Error = "Неизвестная ошибка"
		return data, errorlog.NewError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		l := &svccourse.LessonItems{}
		if err := rows.Scan(&l.Id, &l.Name, &l.ExerciseJson, &l.TaskDescription, &l.LessonNumber, &l.DeadlineDate); err != nil {
			data.Status = http.StatusInternalServerError
			data.Error = "Неизвестная ошибка"
			return data, errorlog.NewError(err.Error())
		}
		data.LessonDataItems = append(data.LessonDataItems, l)
	}

	return data, nil
}
