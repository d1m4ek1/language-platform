package models

import (
	"database/sql"
	"english/backend/shared/errorlog"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"net/http"
)

type Client struct {
	Id              int64         `json:"id"`
	Login           string        `json:"login"`
	Email           string        `json:"email"`
	FirstName       string        `json:"firstName" db:"first_name"`
	LastName        string        `json:"lastName" db:"last_name"`
	RegToken        string        `json:"reg_token"`
	Password        string        `json:"password"`
	IsTeacher       bool          `json:"isTeacher" db:"teacher"`
	SetSubCourseIDS pq.Int64Array `json:"setSubCourseIDS" db:"set_course_ids"`
}

func (u *Client) GetClient(db *sqlx.DB) (int, error) {
	if err := db.Get(u, `
	SELECT
	    client.id,
	    login,
	    password,
	    ci.teacher
	FROM
	    client,
	    client_info ci
	WHERE
	    login=$1 AND ci.id=client.id`, u.Login); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (u *Client) CheckHaveClientByLoginAndEmail(db *sqlx.DB) (string, int, error) {
	var isHaveClient bool
	var messageError string

	if err := db.Get(&isHaveClient, `
	SELECT EXISTS(SELECT login FROM client WHERE login=$1)`, u.Login); err != nil {
		return messageError, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}
	if isHaveClient {
		messageError = fmt.Sprintf(`Логин %s уже занят, попробуйте другой`, u.Login)
		return messageError, http.StatusOK, nil
	}

	if err := db.Get(&isHaveClient, `
	SELECT EXISTS(SELECT email FROM client WHERE email=$1)`, u.Email); err != nil {
		return messageError, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}
	if isHaveClient {
		messageError = "Электронная почта уже занята"
		return messageError, http.StatusOK, nil
	}

	return messageError, http.StatusOK, nil
}

func (u *Client) CheckRegToken(db *sqlx.DB) (bool, int, error) {
	if err := db.Get(&u.SetSubCourseIDS, `
	SELECT set_course_ids FROM registration_token WHERE token=$1`, u.RegToken); errors.Is(err, sql.ErrNoRows) {
		return false, http.StatusOK, nil
	} else if err != nil {
		return false, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	return true, http.StatusOK, nil
}

func (u *Client) InsertClient(db *sqlx.DB) (int, error) {
	stmt, err := db.PrepareNamed(`
	INSERT INTO 
		client
		(login, password, email)
	VALUES
	    (:login, :password, :email) 
	RETURNING id`)
	if err != nil {
		return http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	if err := stmt.Get(&u.Id, u); err != nil {
		return http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	if err := u.insertClientInfo(db); err != nil {
		return http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	return http.StatusOK, nil
}

func (u *Client) insertClientInfo(db *sqlx.DB) error {
	stmt, err := db.PrepareNamed(`
	INSERT INTO 
		client_info
		(id, first_name, last_name, sub_course_ids)
	VALUES 
	    (:id, :first_name, :last_name, :set_course_ids)`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(u); err != nil {
		return err
	}

	return nil
}
