package main

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

func CheckRequiredFields(request interface{}) (ok bool, result string) {
	valid := validation.Validation{}
	b, err := valid.Valid(request)
	if err != nil {
		return false, fmt.Sprintf("Error %v\n", err)
	}
	if !b {
		// validation does not pass
		result := ""
		for _, err := range valid.Errors {
			result += err.Key[0:len(err.Key)-len(".Required")] + ","
		}

		result = result[0 : len(result)-1]

		return false, result
	}

	return true, ""
}
