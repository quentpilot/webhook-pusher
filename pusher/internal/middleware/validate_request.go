package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetValidated retrieves request payload struct T handled by BindAndValidate middleware.
// Returns false if the type not corresponding to the last stored validated struct.
func GetValidated[T any](c *gin.Context) (*T, bool) {
	val, exists := c.Get("validated")
	if !exists {
		return nil, false
	}

	obj, ok := val.(*T)

	return obj, ok
}

/*
getErrorMessage retrieve the relative error message to rule defined in struct field.

Return a default message if no message is defined for current rule validation error.
*/
func getErrorMessage(req any, fe validator.FieldError) string {
	t := reflect.TypeOf(req)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	field := strings.ToLower(fe.Field())
	rule := fe.Tag()

	mKey := fmt.Sprintf("message_%s", rule)
	mDef := fmt.Sprintf("field '%s' must be %s", field, rule)

	if f, ok := t.FieldByName(fe.Field()); ok {
		msg := f.Tag.Get(mKey)
		if msg != "" {
			return fmt.Sprintf("field '%s' %s", field, msg)
		}
	}
	return mDef
}

/*
BindAndValidate middleware allows to bind request payload as T struct.

Abort request with a structured error message list for invalid fields.

Stores validated request into struct T provided to be retrieve easily with GetValidated() function.
*/
func BindAndValidate[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.ShouldBindJSON(&req); err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				errorsList := make([]map[string]string, 0)
				for _, fe := range verr {
					field := strings.ToLower(fe.Field())

					// Format each field in error with the reason
					errorsList = append(errorsList, map[string]string{
						"field": field,
						"error": getErrorMessage(req, fe),
					})
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "invalid field(s)",
					"details": errorsList,
				})
				return
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid request body",
				"details": err.Error(),
			})
			return
		}

		// Store the parsed object in context
		c.Set("validated", &req)
		c.Next()
	}
}
