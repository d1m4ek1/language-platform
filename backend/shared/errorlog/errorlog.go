package errorlog

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Error interface {
	getLogDate()

	wrapError() string
}

type ErrorData struct {
	Err      string
	Filename string
	Date     string
	Line     int
}

func (e *ErrorData) getLogDate() {
	t := time.Now()
	e.Date = fmt.Sprintf(`%d %s %d, %s in %d:%d:%d`, t.Day(), t.Month().String(), t.Year(),
		t.Weekday().String(), t.Hour(), t.Minute(), t.Second())
}

func (e *ErrorData) wrapError() string {
	var breakLine []string
	formatString := `%vError: %s%vDate: %s%vLine: %d%vFilename: %s%v%s`

	for i := 0; i < len("Filename: "+e.Filename); i++ {
		breakLine = append(breakLine, "-")
	}

	return fmt.Sprintf(formatString, "\n", e.Err, "\n", e.Date, "\n", e.Line, "\n", e.Filename, "\n",
		strings.Join(breakLine, ""))
}

func NewError(str string) error {
	_, filename, line, _ := runtime.Caller(1)

	var wrapError Error = &ErrorData{
		Err:      str,
		Filename: filename,
		Line:     line,
	}

	wrapError.getLogDate()

	return fmt.Errorf(wrapError.wrapError())
}
