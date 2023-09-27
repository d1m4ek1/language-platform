package models

import (
	"encoding/json"
	svc_client "english/backend/api/proto/svc-client"
	"english/backend/shared/errorlog"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type DataClient struct {
	UserID            int64             `db:"id"`
	FirstName         string            `json:"firstName" db:"first_name"`
	LastName          string            `json:"lastName" db:"last_name"`
	RegDate           string            `json:"regDate" db:"reg_date"`
	LanguageItems     []LanguageItems   `json:"languageItems" db:"language_items"`
	AboutMe           string            `json:"aboutMe" db:"about_me"`
	IsTeacher         bool              `json:"isTeacher" db:"teacher"`
	InterfaceLang     string            `json:"interfaceLang" db:"interface_lang"`
	RegistrationToken RegistrationToken `json:"registrationToken"`
	Avatar            string
}

type RegistrationToken struct {
	RegToken     string `json:"regToken" db:"token"`
	CreatedDate  string `json:"createdDate" db:"created_date"`
	DeadlineDate string `json:"deadlineDate" db:"deadline_date"`
}

type LanguageItems struct {
	Lang string `json:"lang"`
	Lvl  string `json:"lvl"`
}

func (d *DataClient) GetDataClient(db *sqlx.DB) (*svc_client.GetDataClientResponse, error) {
	res := &svc_client.GetDataClientResponse{
		ResponseStatus: &svc_client.ResponseStatus{
			Status: http.StatusOK,
		},
	}

	var jsonLanguageItems []byte
	if err := db.QueryRow(`
	SELECT 
	    first_name, last_name, reg_date, about_me, avatar, teacher, interface_lang, language_items
	FROM
		client_info
	WHERE
	    id=$1`, d.UserID).Scan(&d.FirstName, &d.LastName, &d.RegDate,
		&d.AboutMe, &d.Avatar, &d.IsTeacher, &d.InterfaceLang, &jsonLanguageItems); err != nil {
		return res, errorlog.NewError(err.Error())
	}

	var hasRegTokenColumn bool
	if err := db.Get(&hasRegTokenColumn, `
	SELECT exists(
		select 
		    * 
		from 
		    registration_token 
		where 
		    client_id=$1)`, d.UserID); err != nil {
		return res, errorlog.NewError(err.Error())
	}

	if hasRegTokenColumn {
		if err := db.QueryRow(`
		SELECT 
			token, created_date, deadline_date
		FROM
			registration_token
		WHERE
	    client_id=$1`, d.UserID).Scan(&d.RegistrationToken.RegToken, &d.RegistrationToken.CreatedDate,
			&d.RegistrationToken.DeadlineDate); err != nil {
			return res, errorlog.NewError(err.Error())
		}
	}

	if err := json.Unmarshal(jsonLanguageItems, &d.LanguageItems); err != nil {
		return res, errorlog.NewError(err.Error())
	}

	dataClient := &svc_client.DataClient{}
	if err := copier.Copy(&dataClient, d); err != nil {
		return res, errorlog.NewError(err.Error())
	}
	res.DataClient = dataClient

	return res, nil
}

func (d *DataClient) UpdateDataClient(db *sqlx.DB) (*svc_client.SaveDataClientResponse, error) {
	res := &svc_client.SaveDataClientResponse{
		ResponseStatus: &svc_client.ResponseStatus{
			Status: http.StatusOK,
		},
	}

	stmt, err := db.Prepare(`
	UPDATE
		client_info
	SET
	    first_name=(
	        select 
	            case
	            	when $1 !='' then $2
	        	else first_name
	        end
	    ), 
	    last_name=(
	        select
	            case
	            	when $3 !='' then $4
	        	else last_name
	        end
	    ),
	    language_items=(
	        select
	            case
	            	when $5::json is not null then $6::json
	        	else language_items
	        end
	    ),
	    about_me=(
	        select
	            case
	            	when $7 !='' then $8
	        	else about_me
	        end
	    ),
	    interface_lang=(
	        select
	            case
	            	when $9 !='' then $10
	        	else interface_lang
	        end
	    ),
		avatar=(
	        select
	            case
	            	when $11 !='' then $12
	        	else avatar
	        end
	    )
	WHERE 
	    id=$13`)
	if err != nil {
		return res, errorlog.NewError(err.Error())
	}

	languageItems, err := json.Marshal(d.LanguageItems)
	if err != nil {
		return res, errorlog.NewError(err.Error())
	}

	if _, err := stmt.Exec(d.FirstName, d.FirstName, d.LastName, d.LastName, languageItems, languageItems,
		d.AboutMe, d.AboutMe, d.InterfaceLang, d.InterfaceLang, d.Avatar, d.Avatar, d.UserID); err != nil {
		return res, errorlog.NewError(err.Error())
	}

	if err := d.UpdateRegistrationToken(db); err != nil {
		return res, err
	}

	return res, nil
}

func (d *DataClient) UpdateRegistrationToken(db *sqlx.DB) error {
	var hasRegTokenColumn bool
	if err := db.Get(&hasRegTokenColumn, `
	SELECT exists(
		select 
		    * 
		from 
		    registration_token 
		where 
		    client_id=3)`); err != nil {
		return errorlog.NewError(err.Error())
	}

	if hasRegTokenColumn {
		stmt, err := db.Prepare(`
	UPDATE
		registration_token
	SET
	    token=(
	        select 
	            case
	            	when $1 !='' then $2
	        	else token
	        end
	    ), 
	    created_date=(
	        select
	            case
	            	when $3 !='' then $4
	        	else created_date
	        end
	    ),
	    deadline_date=(
	        select 
	            case
	            	when $5 != '' then $6
	        	else deadline_date
	        end
	    )
	WHERE
	    client_id=$7`)
		if err != nil {
			return errorlog.NewError(err.Error())
		}

		if _, err := stmt.Exec(d.RegistrationToken.RegToken, d.RegistrationToken.RegToken,
			d.RegistrationToken.CreatedDate, d.RegistrationToken.CreatedDate, d.RegistrationToken.DeadlineDate,
			d.RegistrationToken.DeadlineDate, d.UserID); err != nil {
			return errorlog.NewError(err.Error())
		}
	} else {
		stmt, err := db.Prepare(`
		INSERT INTO 
			registration_token (token, created_date, deadline_date, client_id) 
			VALUES ($1, current_date, current_date + interval '1 month', $2)
		`)
		if err != nil {
			return errorlog.NewError(err.Error())
		}

		if _, err := stmt.Exec(d.RegistrationToken.RegToken, d.UserID); err != nil {
			return errorlog.NewError(err.Error())
		}
	}

	return nil
}

func NewDataClient(req *svc_client.DataClient) *DataClient {
	return &DataClient{
		UserID: req.UserId,
	}
}
