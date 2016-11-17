package lib

import "errors"

var (
	ErrWrongDateFormat = errors.New(`Formato de fecha erróneo - use MM/DD/YYYY instead. \n`)
)
