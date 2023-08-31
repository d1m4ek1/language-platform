package uploads

import (
	svccourse "english/backend/api/proto/svc-course"
	"english/backend/shared/errorlog"
	"os"
)

func DeleteFile(paths map[string]string) error {
	for _, path := range paths {
		if err := os.Remove("./sources" + path); err != nil {
			return errorlog.NewError(err.Error())
		}
	}

	return nil
}

func DeleteAllFile(paths []string) error {
	for _, path := range paths {
		if err := os.Remove("./sources" + path); err != nil {
			return errorlog.NewError(err.Error())
		}
	}

	return nil
}

func DeleteLessonFileItems(lessonItems []*svccourse.LessonItems) error {
	for _, lesson := range lessonItems {
		for _, path := range lesson.DeleteFiles {
			if err := os.Remove("./sources" + path); err != nil {
				return errorlog.NewError(err.Error())
			}
		}

		lesson.DeleteFiles = nil
	}

	return nil
}
