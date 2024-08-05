package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
)

func ParseBody(r *http.Request, x any) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}

func UpdateFields(target any, updates any) {
	targetVal := reflect.ValueOf(target).Elem()
	updatesVal := reflect.ValueOf(updates).Elem()

	for i := 0; i < updatesVal.NumField(); i++ {
		field := updatesVal.Field(i)
		if !field.IsZero() {
			targetVal.Field(i).Set(field)
		}
	}
}
