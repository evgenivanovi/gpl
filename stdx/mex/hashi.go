package mex

import (
	"strings"

	"github.com/evgenivanovi/gpl/std/str"
	me "github.com/hashicorp/go-multierror"
)

func AppendFormat(sep string) me.ErrorFormatFunc {
	return func(es []error) string {
		res := strings.Builder{}
		for index, err := range es {
			if index > 0 {
				res.WriteString(sep)
			}
			res.WriteString(err.Error())
		}
		return res.String()
	}
}

func AppendWithoutAfterSuffixFormat(sep string, suffix string) me.ErrorFormatFunc {
	return func(es []error) string {
		res := strings.Builder{}
		for index, err := range es {
			if index > 0 {
				res.WriteString(sep)
			}
			res.WriteString(str.TruncateAfter(err.Error(), suffix))
		}
		return res.String()
	}
}
