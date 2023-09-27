package models

import (
	"english/backend/shared/errorlog"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Client struct {
	Id         int64  `json:"id"`
	Login      string `json:"login"`
	Email      string `json:"email"`
	FirstName  string `json:"firstName" db:"first_name"`
	LastName   string `json:"lastName" db:"last_name"`
	RegToken   string `json:"reg_token"`
	Password   string `json:"password"`
	IsTeacher  bool   `json:"isTeacher" db:"teacher"`
	SubTeacher int    `json:"subTeacher" db:"sub_teacher"`
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
	if err := db.Get(&u.SubTeacher, `
	SELECT 
	case
		when exists(select token FROM registration_token WHERE token=$1) = true 
	    then (select client_id from registration_token WHERE token=$2)
		else 0
	end`, u.RegToken, u.RegToken); err != nil {
		return false, http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	return !(u.SubTeacher == 0), http.StatusOK, nil
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
		(id, first_name, last_name, sub_teacher)
	VALUES 
	    (:id, :first_name, :last_name, :sub_teacher)`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(u); err != nil {
		return err
	}

	if _, err := db.NamedExec(`
	UPDATE 
		client_info
	SET
	    sub_students=array_append(sub_students, :id)
	WHERE 
	    id=:sub_teacher`, u); err != nil {
		return err
	}

	return nil
}
