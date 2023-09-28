package models

import (
	"encoding/json"
	svcmodule "english/backend/api/proto/svc-module"
	"english/backend/shared/errorlog"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type Module struct {
	ModuleId    int64         `json:"moduleId" db:"id"`
	UserID      int64         `json:"userId" db:"user_id"`
	Name        string        `json:"name"`
	Language    string        `json:"language"`
	Hidden      bool          `json:"hidden"`
	Description string        `json:"description" db:"description"`
	Price       int           `json:"price"`
	SubStudents pq.Int64Array `json:"subStudents" db:"sub_students"`
}

type AddLessonToCreateModule struct {
	UserID            int64         `json:"userId"`
	ModuleID          int64         `json:"moduleId"`
	LessonItems       []LessonItems `json:"lessonItems"`
	DeleteLessonItems []int64       `json:"deleteLessonItems"`
}

type LessonItems struct {
	Id              int64           `json:"id" db:"id"`
	ModuleID        int             `json:"-" db:"module_id"`
	Name            string          `json:"name"`
	DeadlineDate    time.Time       `json:"deadlineDate" db:"deadline_date"`
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

func newExercise(e *svcmodule.ExerciseItems) ExerciseItems {
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

func newLesson(l *svcmodule.LessonItems) (LessonItems, error) {
	lesson := LessonItems{
		Id:              l.Id,
		Name:            l.Name,
		TaskDescription: l.TaskDescription,
		LessonNumber:    int(l.LessonNumber),
	}

	var err error
	lesson.DeadlineDate, err = time.Parse(time.RFC3339, l.DeadlineDate)
	if err != nil {
		return LessonItems{}, errorlog.NewError(err.Error())
	}

	for _, exerciseItem := range l.ExerciseItems {
		lesson.ExerciseItems = append(lesson.ExerciseItems, newExercise(exerciseItem))
	}

	return lesson, nil
}

func NewModule(req *svcmodule.CreateModuleRequest) *Module {
	return &Module{
		UserID:      req.UserId,
		Name:        req.Name,
		Language:    req.Language,
		Hidden:      req.Hidden,
		Description: req.Description,
		Price:       int(req.Price),
		ModuleId:    req.ModuleId,
		SubStudents: req.SubStudents,
	}
}

func NewAddLessons(req *svcmodule.AddLessonRequest) (*AddLessonToCreateModule, error) {
	l := &AddLessonToCreateModule{
		UserID:            req.UserId,
		ModuleID:          req.ModuleId,
		DeleteLessonItems: req.DeleteLessonItems,
	}

	for _, item := range req.LessonItems {
		lesson, err := newLesson(item)
		if err != nil {
			return nil, err
		}

		l.LessonItems = append(l.LessonItems, lesson)
	}

	return l, nil
}

func (c *Module) InsertModule(db *sqlx.DB) (int, int, error) {
	stmtModule, err := db.PrepareNamed(`
	INSERT INTO 
		module
		(user_id, name, language, description, hidden, price, sub_students) 
	VALUES 
		(:user_id, :name, :language, :description, :hidden, :price, :sub_students)
	RETURNING id`)
	if err != nil {
		return 0, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	var moduleID int
	if err := stmtModule.Get(&moduleID, c); err != nil {
		return 0, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	return moduleID, http.StatusOK, nil
}

func (c *Module) SetModule(db *sqlx.DB) (int64, int, error) {
	stmtModule, err := db.PrepareNamed(`
	UPDATE
		module
	SET
	    name=:name,
	    language=:language,
	    description=:description,
	    hidden=:hidden,
	    price=:price,
	    editing_date=current_date,
	    sub_students=:sub_students
	WHERE
	    user_id=:user_id AND id=:id`)
	if err != nil {
		return 0, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	if _, err := stmtModule.Exec(&c); err != nil {
		return 0, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	return c.ModuleId, http.StatusOK, nil
}

func (a *AddLessonToCreateModule) AddLessonToCreatedModule(db *sqlx.DB) (int, error) {
	stmtLesson, err := db.Prepare(`
	INSERT INTO
		lesson
		(module_id, name, exercise, task_description, lesson_number, deadline_date)
	VALUES
		($1, $2, $3, $4, $5, $6)`)
	if err != nil {
		return http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	for _, item := range a.LessonItems {
		exercise, err := json.Marshal(item.ExerciseItems)
		if err != nil {
			return http.StatusInternalServerError, errorlog.NewError(err.Error())
		}

		if _, err := stmtLesson.Exec(a.ModuleID, item.Name, exercise, item.TaskDescription, item.LessonNumber, item.DeadlineDate); err != nil {
			return http.StatusInternalServerError, errorlog.NewError(err.Error())
		}
	}

	return http.StatusOK, nil
}

func (a *AddLessonToCreateModule) SetLessons(db *sqlx.DB) (int, error) {
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
			    date_editing=current_date,
			    deadline_date=$5
			WHERE
			    id=$6 AND module_id=$7`)
			if err != nil {
				return http.StatusInternalServerError, errorlog.NewError(err.Error())
			}

			if _, err := stmt.Exec(item.Name, exercise, item.TaskDescription,
				item.LessonNumber, item.DeadlineDate, item.Id, a.ModuleID); err != nil {
				return http.StatusInternalServerError, errorlog.NewError(err.Error())
			}

			continue
		}

		stmtLesson, err := db.Prepare(`
		INSERT INTO
			lesson
			(module_id, name, exercise, task_description, lesson_number, deadline_date)
		VALUES
			($1, $2, $3, $4, $5, $6)`)
		if err != nil {
			return http.StatusInternalServerError, errorlog.NewError(err.Error())
		}

		if _, err := stmtLesson.Exec(a.ModuleID, item.Name, exercise, item.TaskDescription, item.LessonNumber, item.DeadlineDate); err != nil {
			return http.StatusInternalServerError, errorlog.NewError(err.Error())
		}
	}

	return http.StatusOK, nil
}
