package utils

import (
	"go-article/responsegraph"
)

func IsEmptyString(variable ...string) bool {
	for _, isi := range variable {
		if isi == "" {
			return true
		}
	}
	return false
}

func GetResNoData(status string, message string) responsegraph.ResponseGenericIn {
	resp := responsegraph.ResponseGenericIn{
		Status:  status,
		Message: message,
	}
	return resp
}
