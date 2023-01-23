package libraries

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationRules(messages []map[string]interface{}, err error) interface{} {
	var verr validator.ValidationErrors
	var out []interface{}
	if errors.As(err, &verr) {
		for _, fe := range verr {
			for _, msg := range messages {
				if field, ok := msg["field"].(string); ok {
					if strings.ToLower(field) == strings.ToLower(fe.Field()) {
						if rec, ok := msg["messages"].(map[string]interface{}); ok {
							for key, val := range rec {
								if key == fe.Tag() {
									out = append(out, map[string]interface{}{
										"field": strings.ToLower(fe.Field()),
										"error": val,
									})
								}
							}
						} else {
							panic("Custom validate : messages harus map[string]interface{}")
						}
					}
				} else {
					panic("Custom validate : " + field + " (value harus string)")
				}
			}
		}
	}
	return out
}
