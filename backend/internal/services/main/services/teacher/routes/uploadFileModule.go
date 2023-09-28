package routes

import (
	"context"
	"encoding/json"
	svcmodule "english/backend/api/proto/svc-module"
	"english/backend/shared/errorlog"
	"english/backend/shared/uploads"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func uploadMainImage(ctx *gin.Context, c svcmodule.ModuleServiceClient) (int, error) {
	moduleId, err := strconv.Atoi(ctx.Query("module_id"))
	if err != nil {
		return http.StatusBadRequest, errorlog.NewError(err.Error())
	}
	userId, _ := ctx.Get("userId")
	isTeacher, _ := ctx.Get("isTeacher")

	form, _ := ctx.MultipartForm()
	files := form.File["mainFile"]
	fileForModuleRequest := &svcmodule.FileForModuleRequest{
		UserId:   userId.(int64),
		ModuleId: int64(moduleId),
	}

	for _, file := range files {
		headerContentType := file.Header.Get("Content-Type")
		contentType := ""

		if strings.Contains(headerContentType, "image/") {
			contentType = "mainImg"
		}

		if strings.Contains(headerContentType, "video/") {
			contentType = "mainVideo"
		}

		if isValidFile, err := uploads.ValidateFile(file); err != nil {
			return http.StatusBadRequest, err
		} else if !isValidFile {
			return http.StatusOK, uploads.ErrorFileIsNotSupported
		}

		resultPath, status, err := uploads.UploadMainFileModule(uploads.UploadData{
			File:         file,
			UserID:       userId.(int64),
			IsTeacher:    isTeacher.(bool),
			ModuleID:     moduleId,
			FormFileName: contentType,
		})
		if err != nil {
			return status, err
		}

		fileForModuleRequest.FileItems = append(fileForModuleRequest.FileItems, &svcmodule.FileForModuleRequest_File{
			TypeFile: contentType,
			Path:     resultPath,
		})
	}

	response, err := c.SetFileForModule(context.Background(), fileForModuleRequest)
	if err != nil {
		return int(response.Status), err
	}

	return http.StatusOK, nil
}

func containHeaderContentType(headerContentType string) string {
	if strings.Contains(headerContentType, "image/") {
		return "image"
	}

	if strings.Contains(headerContentType, "video/") {
		return "video"
	}

	if strings.Contains(headerContentType, "audio/") {
		return "audio"
	}

	return ""
}

func uploadLessonFiles(ctx *gin.Context) (map[int]string, int, error) {
	moduleId, err := strconv.Atoi(ctx.Query("module_id"))
	if err != nil {
		return nil, http.StatusBadRequest, errorlog.NewError(err.Error())
	}

	userId, _ := ctx.Get("userId")
	isTeacher, _ := ctx.Get("isTeacher")

	lessonId, err := strconv.Atoi(ctx.Query("lesson_id"))
	if err != nil {
		return nil, http.StatusBadRequest, errorlog.NewError(err.Error())
	}

	exerciseFiles := make(map[int]string)
	if err := json.Unmarshal([]byte(ctx.Query("exercise_ids")), &exerciseFiles); err != nil {
		return nil, http.StatusBadRequest, errorlog.NewError(err.Error())
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, http.StatusBadRequest, errorlog.NewError(err.Error())
	}

	for exerciseId := range exerciseFiles {
		files := form.File[fmt.Sprintf("exercise_%d", exerciseId)]
		for _, file := range files {
			contentType := containHeaderContentType(file.Header.Get("Content-Type"))

			if isValidFile, err := uploads.ValidateFile(file); err != nil {
				return nil, http.StatusBadRequest, err
			} else if !isValidFile {
				return nil, http.StatusOK, uploads.ErrorFileIsNotSupported
			}

			resultPath, status, err := uploads.UploadLessonFileModule(uploads.UploadData{
				File:         file,
				UserID:       userId.(int64),
				IsTeacher:    isTeacher.(bool),
				ModuleID:     moduleId,
				FormFileName: contentType,
				LessonID:     lessonId,
				ExerciseID:   exerciseId,
			})
			if err != nil {
				return nil, status, err
			}

			exerciseFiles[exerciseId] = resultPath
		}
	}

	return exerciseFiles, http.StatusOK, nil
}

func uploadAvatarFile(ctx *gin.Context) (string, int, error) {
	userId, _ := ctx.Get("userId")
	isTeacher, _ := ctx.Get("isTeacher")

	file, err := ctx.FormFile("avatar")
	if err != nil {
		return "", http.StatusBadRequest, errorlog.NewError(err.Error())
	}

	if !strings.Contains(file.Header.Get("Content-Type"), "image/") {
		fmt.Println(true)
		return "", http.StatusOK, uploads.ErrorFileIsNotSupported
	}

	if isValidFile, err := uploads.ValidateFile(file); err != nil {
		return "", http.StatusBadRequest, err
	} else if !isValidFile {
		return "", http.StatusOK, uploads.ErrorFileIsNotSupported
	}

	resultPath, status, err := uploads.UploadAvatarFile(uploads.UploadData{
		File:         file,
		UserID:       userId.(int64),
		IsTeacher:    isTeacher.(bool),
		FormFileName: "avatar",
	})
	if err != nil {
		return "", status, err
	}

	return resultPath, http.StatusOK, nil
}

func uploadMainFilesByUploadType(ctx *gin.Context, c svcmodule.ModuleServiceClient) {
	status, err := uploadMainImage(ctx, c)
	if errors.Is(err, uploads.ErrorFileIsNotSupported) {
		ctx.JSON(status, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	} else if err != nil {
		fmt.Println(err)
		ctx.JSON(status, gin.H{
			"success": false,
			"error":   "Не удалось загрузить файл",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func uploadLessonFilesByUploadType(ctx *gin.Context) {
	exerciseFiles, status, err := uploadLessonFiles(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(status, gin.H{
			"success": false,
			"error":   "Не удалось загрузить файл",
		})
		return
	}

	jsonStr, err := json.Marshal(exerciseFiles)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(status, gin.H{
			"success": false,
			"error":   "Не удалось загрузить файл",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":       true,
		"exerciseFiles": string(jsonStr),
	})
}

func uploadAvatarByUploadType(ctx *gin.Context) {
	resultPath, status, err := uploadAvatarFile(ctx)
	if errors.Is(err, uploads.ErrorFileIsNotSupported) {
		ctx.JSON(status, gin.H{
			"success": false,
			"error":   fmt.Sprintf(`%s`, uploads.ErrorFileIsNotSupported),
		})
		return
	} else if err != nil {
		fmt.Println(err)
		ctx.JSON(status, gin.H{
			"success": false,
			"error":   "Не удалось загрузить файл",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"avatar":  resultPath,
	})
}

func UploadFileModule(ctx *gin.Context, c svcmodule.ModuleServiceClient) {
	uploadType := ctx.Query("upload_type")

	if uploadType == "main_files" {
		uploadMainFilesByUploadType(ctx, c)
	}

	if uploadType == "lesson_files" {
		uploadLessonFilesByUploadType(ctx)
	}

	if uploadType == "avatar" {
		uploadAvatarByUploadType(ctx)
	}
}
