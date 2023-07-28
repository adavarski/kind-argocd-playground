package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/prometheus/prometheus/promql/parser"
	"go.uber.org/zap"
)

var lineBreak = regexp.MustCompile(`\r?\n`)

type server struct {
	path      string
	interval  time.Duration
	startTime time.Time
	logger    *zap.Logger
}

func newServer(path string, interval time.Duration) *server {
	return &server{
		path:      path,
		interval:  interval,
		startTime: time.Now(),
		logger:    zap.L().Named("dummy-metrics"),
	}
}

func (s server) metrics(w http.ResponseWriter, r *http.Request) {

	content, err := os.ReadFile(s.path)
	if err != nil {
		http.Error(w, "failed to read file", http.StatusInternalServerError)
		s.logger.Error("failed to read file", zap.String("path", s.path), zap.Error(err))
		return
	}

	elapsedSeconds := time.Since(s.startTime).Seconds()
	flame := int(elapsedSeconds / s.interval.Seconds())

	lines := lineBreak.Split(string(content), -1)
	var commentLines []string
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			commentLines = append(commentLines, line)
			continue
		}
		labels, values, err := parser.ParseSeriesDesc(line)
		if err != nil {
			http.Error(w, "failed to parse line", http.StatusInternalServerError)
			s.logger.Error("failed to parse line", zap.String("line", line), zap.Error(err))
			return
		}
		if len(values) == 0 {
			continue
		}
		var name string
		var ls []string
		for _, l := range labels {
			if l.Name == "__name__" {
				name = l.Value
			} else {
				ls = append(ls, fmt.Sprintf(`%s="%s"`, l.Name, l.Value))
			}
		}
		index := flame % len(values)
		if !values[index].Omitted {
			comment := strings.Join(commentLines, "\n")
			if len(comment) != 0 {
				comment += "\n"
			}
			_, err = io.WriteString(w, fmt.Sprintf("%s%s{%s} %f\n", comment, name, strings.Join(ls, ","), values[index].Value))
			if err != nil {
				http.Error(w, "failed to write response", http.StatusInternalServerError)
				s.logger.Error("failed to write response", zap.Error(err))
				return
			}
		}
		commentLines = []string{}
	}
}
