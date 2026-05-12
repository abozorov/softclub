package main

import (
	"fmt"
	"unicode/utf8"
)

type Logger interface {
	Log(message string) error
}

type FileLogger struct {
	Logs []string
}
type DBLogger struct {
	Logs []string
}

func (f *FileLogger) Log(message string) error {
	f.Logs = append(f.Logs, message)
	return nil
}

func (d *DBLogger) Log(message string) error {
	if utf8.RuneCountInString(message) < 5 {
		return fmt.Errorf("message \"%s\" so short!", message)
	}
	d.Logs = append(d.Logs, message)
	return nil
}

func main() {

	fileLog := &FileLogger{
		Logs: []string{},
	}
	dbLog := &DBLogger{
		Logs: []string{},
	}

	data := map[string]Logger{
		"go":        fileLog,
		"some logs": fileLog,
		"lang":      dbLog,
		"worl!d!":   dbLog,
	}
	errs := make([]error, 0, len(data))

	for k, v := range data {
		err := v.Log(k)
		if err != nil {
			errs = append(errs, err)
		}
	}

	fmt.Printf("file logs: %v\nDB logs: %v\nDB errs: %v\n", *fileLog, *dbLog, errs)
}
