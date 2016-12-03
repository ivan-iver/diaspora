// Package time encapsulate common logic with related time operations
package time

import (
	"fmt"
	"strconv"
	"time"
)

const (
	stdFormat   = `2006/01/02 15:04 05`
	serieFormat = `20060102150405`
)

// Serie is used to create new formated value from a time data
type Serie struct {
	Timestamp time.Time
	Formated  string
	Value     int64
}

// NewSerie returns a Serie struct initializated and error if exists
func NewSerie() (serie *Serie, err error) {
	var init = time.Now()
	serie = &Serie{
		Timestamp: init,
		Formated:  fmt.Sprintf("%v", init.Format(serieFormat)),
	}
	serie.Value, err = strconv.ParseInt(serie.Formated, 10, 64)
	return
}

// Next generate a new time serie with 3 incremental values
func (s *Serie) Next() {
	s.Value = s.Value + 3
	s.Formated = fmt.Sprintf("%v", s.Value)
}
