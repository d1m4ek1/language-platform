package uploads

import (
	"english/backend/shared/errorlog"
	"os"
)

type InitUserDir struct {
	UserID    int64
	PathToDir string
}

func (i *InitUserDir) existsDir() (bool, error) {
	if _, err := os.Stat(i.PathToDir); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, errorlog.NewError(err.Error())
	}

	return true, nil
}

func (i *InitUserDir) Init() error {
	if isExist, err := i.existsDir(); err != nil {
		return err
	} else if isExist {
		return nil
	}

	if err := os.MkdirAll(i.PathToDir, os.ModePerm); err != nil {
		return errorlog.NewError(err.Error())
	}

	return nil
}
