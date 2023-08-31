package uploads

import (
	"english/backend/shared/errorlog"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type UploadData struct {
	IsTeacher    bool
	File         *multipart.FileHeader
	UserID       int64
	CourseID     int
	LessonID     int
	ExerciseID   int
	FormFileName string
}

func setFile(path string, file *multipart.FileHeader) error {
	fileWriter, err := os.Create(path)
	if err != nil {
		return errorlog.NewError(err.Error())
	}

	fileReader, err := file.Open()
	if err != nil {
		return errorlog.NewError(err.Error())
	}
	defer fileReader.Close()

	if _, err := io.Copy(fileWriter, fileReader); err != nil {
		return errorlog.NewError(err.Error())
	}

	return nil
}

func UploadMainFileCourse(data UploadData) (string, int, error) {
	if !data.IsTeacher {
		return "", http.StatusOK, errors.New("пользователь не является преподавателем")
	}

	pathToMain := fmt.Sprintf("./sources/teachers/teacher_%d/course_%d/main", data.UserID, data.CourseID)

	initDir := InitUserDir{
		UserID:    data.UserID,
		PathToDir: pathToMain,
	}

	if err := initDir.Init(); err != nil {
		return "", http.StatusInternalServerError, err
	}

	pathToMainFile := fmt.Sprintf("%s/%s_%s", pathToMain, data.FormFileName, data.File.Filename)
	if err := setFile(pathToMainFile, data.File); err != nil {
		return "", http.StatusInternalServerError, err
	}

	publicPath := fmt.Sprintf("/teachers/teacher_%d/course_%d/main/%s_%s", data.UserID, data.CourseID, data.FormFileName, data.File.Filename)
	return publicPath, http.StatusOK, nil
}

func UploadLessonFileCourse(data UploadData) (string, int, error) {
	if !data.IsTeacher {
		return "", http.StatusOK, errors.New("пользователь не является преподавателем")
	}

	pathToLesson := fmt.Sprintf("./sources/teachers/teacher_%d/course_%d/lessons/%s", data.UserID, data.CourseID, data.FormFileName)
	fileName := fmt.Sprintf("lesson_%d_exercise_%d_%s_%s", data.LessonID, data.ExerciseID, data.FormFileName, data.File.Filename)

	initDir := InitUserDir{
		UserID:    data.UserID,
		PathToDir: pathToLesson,
	}

	if err := initDir.Init(); err != nil {
		return "", http.StatusInternalServerError, err
	}

	pathToFile := fmt.Sprintf("%s/%s", pathToLesson, fileName)
	if err := setFile(pathToFile, data.File); err != nil {
		return "", http.StatusInternalServerError, err
	}

	publicPath := fmt.Sprintf("/teachers/teacher_%d/course_%d/lessons/%s/%s",
		data.UserID, data.CourseID, data.FormFileName, fileName)
	return publicPath, http.StatusOK, nil
}

func UploadAvatarFile(data UploadData) (string, int, error) {
	if !data.IsTeacher {
		return "", http.StatusOK, errors.New("пользователь не является преподавателем")
	}

	pathToAvatar := fmt.Sprintf("./sources/teachers/teacher_%d/%s", data.UserID, data.FormFileName)

	initDir := InitUserDir{
		UserID:    data.UserID,
		PathToDir: pathToAvatar,
	}

	if err := initDir.Init(); err != nil {
		return "", http.StatusInternalServerError, err
	}

	pathToFile := fmt.Sprintf("%s/%s_%s", pathToAvatar, data.FormFileName, data.File.Filename)
	if err := setFile(pathToFile, data.File); err != nil {
		return "", http.StatusInternalServerError, err
	}

	publicPath := fmt.Sprintf("/teachers/teacher_%d/%s/%s_%s",
		data.UserID, data.FormFileName, data.FormFileName, data.File.Filename)
	return publicPath, http.StatusOK, nil
}
