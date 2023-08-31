package models

import (
	"encoding/json"
	svccourse "english/backend/api/proto/svc-course"
	"english/backend/shared/errorlog"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Course struct {
	CourseId    int64  `json:"courseId" db:"id"`
	UserID      int64  `json:"userId" db:"user_id"`
	Name        string `json:"name"`
	Language    string `json:"language"`
	Hidden      bool   `json:"hidden"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price"`
}

type AddLessonToCreateCourse struct {
	UserID            int64         `json:"userId"`
	CourseID          int64         `json:"courseId"`
	LessonItems       []LessonItems `json:"lessonItems"`
	DeleteLessonItems []int64       `json:"deleteLessonItems"`
}

type LessonItems struct {
	Id              int64           `json:"id" db:"id"`
	CourseID        int             `json:"-" db:"course_id"`
	Name            string          `json:"name"`
	Deadline        int64           `json:"deadline"`
	TaskDescription string          `json:"taskDescription" db:"task_description"`
	LessonNumber    int             `json:"lessonNumber" db:"lesson_number"`
	ExerciseItems   []ExerciseItems `json:"exerciseItems"`
}

type ExerciseItems struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	TaskDescription string `json:"taskDescription" db:"task_description"`
	ExerciseNumber  int    `json:"exerciseNumber" db:"exercise_number"`
	Path            string `json:"path"`
	Task            Task   `json:"task"`
}

type Task struct {
	Type           string          `json:"type"`
	CorrectAnswer  []CorrectAnswer `json:"correctAnswer"`
	ExpandBrackets ExpandBrackets  `json:"expandBrackets"`
}

type CorrectAnswer struct {
	Id     int    `json:"id"`
	Phrase string `json:"phrase"`
	Answer bool   `json:"answer"`
}

type ExpandBrackets struct {
	Text        string                      `json:"text"`
	AnswerItems []ExpandBracketsAnswerItems `json:"answerItems"`
}

type ExpandBracketsAnswerItems struct {
	Id     int    `json:"id"`
	Answer string `json:"answer"`
}

func newExercise(e *svccourse.ExerciseItems) ExerciseItems {
	exercise := ExerciseItems{
		Name:            e.Name,
		TaskDescription: e.TaskDescription,
		Path:            e.Path,
		ExerciseNumber:  int(e.ExerciseNumber),
		Type:            e.Type,
	}

	if e.Task == nil {
		return exercise
	}

	exercise.Task.Type = e.Task.Type
	switch exercise.Task.Type {
	case "correctAnswer":
		for _, answer := range e.Task.CorrectAnswer {
			exercise.Task.CorrectAnswer = append(exercise.Task.CorrectAnswer, CorrectAnswer{
				Id:     int(answer.Id),
				Answer: answer.Answer,
				Phrase: answer.Phrase,
			})
		}
	case "expandBrackets":
		exercise.Task.ExpandBrackets.Text = e.Task.ExpandBrackets.Text

		for _, answerItem := range e.Task.ExpandBrackets.AnswerItems {
			exercise.Task.ExpandBrackets.AnswerItems = append(exercise.Task.ExpandBrackets.AnswerItems,
				ExpandBracketsAnswerItems{
					Id:     int(answerItem.Id),
					Answer: answerItem.Answer,
				})
		}
	}

	return exercise
}

func newLesson(l *svccourse.LessonItems) LessonItems {
	lesson := LessonItems{
		Id:              l.Id,
		Name:            l.Name,
		TaskDescription: l.TaskDescription,
		Deadline:        l.Deadline,
		LessonNumber:    int(l.LessonNumber),
	}

	for _, exerciseItem := range l.ExerciseItems {
		lesson.ExerciseItems = append(lesson.ExerciseItems, newExercise(exerciseItem))
	}

	return lesson
}

func NewCourse(req *svccourse.CreateCourseRequest) *Course {
	return &Course{
		UserID:      req.UserId,
		Name:        req.Name,
		Language:    req.Language,
		Hidden:      req.Hidden,
		Description: req.Description,
		Price:       int(req.Price),
		CourseId:    req.CourseId,
	}
}

func NewAddLessons(req *svccourse.AddLessonRequest) *AddLessonToCreateCourse {
	l := &AddLessonToCreateCourse{
		UserID:            req.UserId,
		CourseID:          req.CourseId,
		DeleteLessonItems: req.DeleteLessonItems,
	}

	for _, item := range req.LessonItems {
		l.LessonItems = append(l.LessonItems, newLesson(item))
	}

	return l
}

func (c *Course) InsertCourse(db *sqlx.DB) (int, int, error) {
	stmtCourse, err := db.PrepareNamed(`
	INSERT INTO 
		course
		(user_id, name, language, description, hidden, price) 
	VALUES 
		(:user_id, :name, :language, :description, :hidden, :price)
	RETURNING id`)
	if err != nil {
		return 0, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	var courseID int
	if err := stmtCourse.Get(&courseID, c); err != nil {
		return 0, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	return courseID, http.StatusOK, nil
}

func (c *Course) SetCourse(db *sqlx.DB) (int64, int, error) {
	stmtCourse, err := db.PrepareNamed(`
	UPDATE
		course
	SET
	    name=:name,
	    language=:language,
	    description=:description,
	    hidden=:hidden,
	    price=:price,
	    editing_date=current_date
	WHERE
	    user_id=:user_id AND id=:id`)
	if err != nil {
		return 0, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	if _, err := stmtCourse.Exec(&c); err != nil {
		return 0, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	return c.CourseId, http.StatusOK, nil
}

func (a *AddLessonToCreateCourse) AddLessonToCreatedCourse(db *sqlx.DB) (int, error) {
	stmtLesson, err := db.Prepare(`
	INSERT INTO
		lesson
		(course_id, name, exercise, task_description, lesson_number)
	VALUES
		($1, $2, $3, $4, $5)`)
	if err != nil {
		return http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	for _, item := range a.LessonItems {
		exercise, err := json.Marshal(item.ExerciseItems)
		if err != nil {
			return http.StatusInternalServerError, errorlog.NewError(err.Error())
		}

		if _, err := stmtLesson.Exec(a.CourseID, item.Name, exercise, item.TaskDescription, item.LessonNumber); err != nil {
			return http.StatusInternalServerError, errorlog.NewError(err.Error())
		}
	}

	return http.StatusOK, nil
}

func (a *AddLessonToCreateCourse) SetLessons(db *sqlx.DB) (int, error) {
	for _, item := range a.DeleteLessonItems {
		if _, err := db.Exec(`
		DELETE FROM 
		           lesson 
		       WHERE 
		           id=$1`, item); err != nil {
			return http.StatusInternalServerError, errorlog.NewError(err.Error())
		}
	}

	for _, item := range a.LessonItems {
		exercise, err := json.Marshal(item.ExerciseItems)
		if err != nil {
			return http.StatusInternalServerError, errorlog.NewError(err.Error())
		}

		if item.Id != 0 {
			stmt, err := db.Prepare(`
			UPDATE
				lesson
			SET
			    name=$1,
			    exercise=$2,
			    task_description=$3,
			    lesson_number=$4,
			    date_editing=current_date
			WHERE
			    id=$5 AND course_id=$6`)
			if err != nil {
				return http.StatusInternalServerError, errorlog.NewError(err.Error())
			}

			if _, err := stmt.Exec(item.Name, exercise, item.TaskDescription,
				item.LessonNumber, item.Id, a.CourseID); err != nil {
				return http.StatusInternalServerError, errorlog.NewError(err.Error())
			}

			continue
		}

		stmtLesson, err := db.Prepare(`
		INSERT INTO
			lesson
			(course_id, name, exercise, task_description, lesson_number)
		VALUES
			($1, $2, $3, $4, $5)`)
		if err != nil {
			return http.StatusInternalServerError, errorlog.NewError(err.Error())
		}

		if _, err := stmtLesson.Exec(a.CourseID, item.Name, exercise, item.TaskDescription, item.LessonNumber); err != nil {
			return http.StatusInternalServerError, errorlog.NewError(err.Error())
		}
	}

	return http.StatusOK, nil
}
