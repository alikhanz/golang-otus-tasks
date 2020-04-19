package middlewares

import (
	"bytes"
	"github.com/rs/zerolog/log"
	"github.com/urfave/negroni"
	"net/http"
	"text/template"
	"time"
)

var LoggerDefaultFormat = "{{.StartTime}} | {{.Status}} | {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}}"

func NewLogger() *Logger {
	logger := &Logger{dateFormat: negroni.LoggerDefaultDateFormat}
	logger.SetFormat(LoggerDefaultFormat)
	return logger
}

type Logger struct {
	dateFormat string
	template   *template.Template
}

func (l *Logger) SetFormat(format string) {
	l.template = template.Must(template.New("negroni_parser").Parse(format))
}

func (l *Logger) SetDateFormat(format string) {
	l.dateFormat = format
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	next(rw, r)

	res := rw.(negroni.ResponseWriter)
	loggerEntry := negroni.LoggerEntry{
		StartTime: start.Format(l.dateFormat),
		Status:    res.Status(),
		Duration:  time.Since(start),
		Hostname:  r.Host,
		Method:    r.Method,
		Path:      r.URL.Path,
		Request:   r,
	}

	buff := &bytes.Buffer{}
	_ = l.template.Execute(buff, loggerEntry)
	log.Debug().Msg(buff.String())
}
