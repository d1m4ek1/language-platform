package models

import (
	"database/sql"
	svccourse "english/backend/api/proto/svc-course"
	"english/backend/shared/errorlog"
	"errors"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type ListOfCourses struct {
	UserID          int64
	CourseData      []CourseData
	ProtoCourseData []*svccourse.ListOfCoursesResponse_CourseData
}

type CourseData struct {
	CourseID    int64  `json:"courseID,omitempty" db:"id"`
	Name        string `json:"name,omitempty" db:"name"`
	Language    string `json:"language,omitempty" db:"language"`
	Hidden      bool   `json:"hidden,omitempty" db:"hidden"`
	Description string `json:"description,omitempty" db:"description"`
	Price       int64  `json:"price,omitempty" db:"price"`
	MainImage   string `json:"mainImage,omitempty" db:"main_img"`
	MainVideo   string `json:"mainVideo,omitempty" db:"main_video"`
	CreatedDate string `json:"createdDate,omitempty" db:"created_date"`
	EditingDate string `json:"editingDate,omitempty" db:"editing_date"`
}

func (l *ListOfCourses) GetListOfCourses(db *sqlx.DB) (int, error) {
	if err := db.Select(&l.CourseData, `
	SELECT
		id, name, language, main_img, main_video, description, hidden, price, created_date, editing_date
	FROM
	    course
	WHERE 
	    user_id=$1`, l.UserID); err != nil {
		return http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	l.ProtoCourseData = make([]*svccourse.ListOfCoursesResponse_CourseData, len(l.CourseData))
	for i, datum := range l.CourseData {
		l.ProtoCourseData[i] = &svccourse.ListOfCoursesResponse_CourseData{
			CourseID:    datum.CourseID,
			Name:        datum.Name,
			Language:    datum.Language,
			Hidden:      datum.Hidden,
			Description: datum.Description,
			Price:       datum.Price,
			MainImage:   datum.MainImage,
			MainVideo:   datum.MainVideo,
			CreatedDate: datum.CreatedDate,
			EditingDate: datum.EditingDate,
		}
	}

	return http.StatusOK, nil
}

func SelectCourseByID(db *sqlx.DB, courseId int64) (*svccourse.CourseDataResponse, error) {
	data := &svccourse.CourseDataResponse{
		Status: http.StatusOK,
	}
	course := svccourse.CreateCourseRequest{}

	if err := db.QueryRow(`
	SELECT
	    name, language, main_img, main_video, description, hidden, price
	FROM
	    course
	WHERE
	    id=$1`, courseId).Scan(&course.Name, &course.Language, &course.MainImg,
		&course.MainVideo, &course.Description, &course.Hidden,
		&course.Price); err != nil {

		data.Status = http.StatusInternalServerError
		data.Error = "Неизвестная ошибка"
		return data, errorlog.NewError(err.Error())
	}

	data.CourseData = &course

	rows, err := db.Query(`
	SELECT
	   id, name, deadline, exercise, task_description, lesson_number
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
		if err := rows.Scan(&l.Id, &l.Name, &l.Deadline, &l.ExerciseJson, &l.TaskDescription, &l.LessonNumber); err != nil {
			data.Status = http.StatusInternalServerError
			data.Error = "Неизвестная ошибка"
			return data, errorlog.NewError(err.Error())
		}
		data.LessonDataItems = append(data.LessonDataItems, l)
	}

	return data, nil
}
