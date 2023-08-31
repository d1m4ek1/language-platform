package main

import (
	"english/backend/config"
	"english/backend/internal/services/main/handler"
	"english/backend/internal/services/main/services/auth"
	"english/backend/internal/services/main/services/client"
	"english/backend/internal/services/main/services/home"
	"english/backend/internal/services/main/services/student"
	"english/backend/internal/services/main/services/teacher"
	"english/backend/shared/errorlog"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"strings"
)

func main() {
	cfg, err := config.NewMainConfig()
	if err != nil {
		log.Fatalln("Failed at config: ", err)
		return
	}

	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	engine.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})

	engine.Static("/teacher/assets", "./web/teacher/assets")
	engine.Static("/student/assets", "./web/student/assets")
	engine.Static("/teachers", "./sources/teachers")
	engine.Static("/default", "./sources/default")

	var goError = make(chan error, 1)
	goError <- nil

	go func() {
		err := handler.UpdateTemplates(engine, false)
		if err != nil {
			goError <- err
		}
	}()

	if <-goError != nil {
		err := <-goError
		fmt.Println(errorlog.NewError(err.Error()))
		return
	}

	authService := auth.RegisterRoutes(engine, &cfg)

	home.RegisterRoutes(engine, authService)
	teacher.RegisterRoutes(engine, authService, &cfg)
	student.RegisterRoutes(engine, authService)
	client.RegisterRoutes(engine, authService, &cfg)

	log.Println("Server is listening on ", cfg.Port)

	if err := engine.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		return
	}
}
