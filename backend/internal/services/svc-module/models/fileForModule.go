package models

import (
	svccourse "english/backend/api/proto/svc-module"
	"english/backend/shared/errorlog"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type FileForCourse struct {
	CourseId  int64 `db:"id"`
	UserId    int64 `db:"user_id"`
	FileItems map[string]string
}

type File struct {
	TypeFile   string
	PathToFile string
}

func (f *FileForCourse) SetFileMain(db *sqlx.DB) (int, error) {
	if _, err := db.Exec(`
	UPDATE 
	    course 
	SET 
	    main_video=(
	        select 
	            case
	            	when $1!='' then $2
	        	else main_video
	        end
	    ), main_img=(
	        select
	            case
	            	when $3!='' then $4
	        	else main_img
	        end
	    )
	WHERE 
	    id=$5
	  AND 
	    user_id=$6`, f.FileItems["mainVideo"], f.FileItems["mainVideo"], f.FileItems["mainImg"], f.FileItems["mainImg"], f.CourseId, f.UserId); err != nil {
		return http.StatusInternalServerError, errorlog.NewError(err.Error())
	}

	return http.StatusOK, nil
}

func NewFileForCourse(req *svccourse.FileForCourseRequest) *FileForCourse {
	data := &FileForCourse{
		CourseId: req.CourseId,
		UserId:   req.UserId,
	}

	data.FileItems = make(map[string]string)

	for _, item := range req.FileItems {
		data.FileItems[item.TypeFile] = item.Path
	}

	return data
}
