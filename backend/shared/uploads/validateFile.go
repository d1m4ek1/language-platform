package uploads

import (
	"english/backend/shared/errorlog"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

var (
	ErrorFileIsNotSupported = errors.New("Загружаемый файл не поддерживается")
)

func ValidateFile(file *multipart.FileHeader) (bool, error) {
	bs := make([]byte, 516)
	f, err := file.Open()

	_, err = f.Read(bs)
	if err != nil && err != io.EOF {
		return false, errorlog.NewError(err.Error())
	}

	detectedContentType := http.DetectContentType(bs)
	var isValidFile bool

	if strings.Contains(detectedContentType, "image/") ||
		strings.Contains(detectedContentType, "video/") ||
		strings.Contains(detectedContentType, "application/octet-stream") {
		isValidFile = true
	}

	return isValidFile, nil
}
