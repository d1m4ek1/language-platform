package handler

import (
	"english/backend/shared/errorlog"
	"fmt"
	"html/template"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func initTemplates() (*template.Template, error) {
	var files []string

	var paths = []string{
		"./web/teacher/pages/*html",
		"./web/student/pages/*html",
	}

	for _, path := range paths {
		file, err := filepath.Glob(path)
		if err != nil {
			return nil, errorlog.NewError(err.Error())
		}

		files = append(files, file...)
	}

	tmplItems, err := template.ParseFiles(files...)
	if err != nil {
		return nil, errorlog.NewError(err.Error())
	}

	return tmplItems, nil
}

func UpdateTemplates(ginEngine *gin.Engine, productMode bool) error {
	const tickTimeMS = 2000
	ticker := time.NewTicker(tickTimeMS * time.Millisecond)

	templates, err := initTemplates()
	if err != nil {
		ticker.Stop()
		return err
	}

	ginEngine.SetHTMLTemplate(templates)

	if !productMode {
		fmt.Println("Template data has been updated")

		for range ticker.C {
			templates, err := initTemplates()
			if err != nil {
				ticker.Stop()
				return err
			}

			ginEngine.SetHTMLTemplate(templates)
		}
	}

	return nil
}
