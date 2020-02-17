package postgres_utils

import (
	"github.com/lib/pq"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
	"log"
	"strings"
)

const (
	NOT_FOUND_KEY = "converting NULL"
)

func HandlePQError(err error) *errors.RestErr {
	prepareErr, ok := err.(*pq.Error)
	if !ok {
		if strings.Contains(err.Error(), NOT_FOUND_KEY) {
			return errors.NewBadRequestError("Not found")
		}
		return errors.NewInternalServerError("Server error")
	}
	log.Println(prepareErr.Code.Name(), prepareErr.Message, prepareErr.Detail, prepareErr.Hint)
	switch prepareErr.Code.Name() {
	case "unique_violation":
		return errors.NewBadRequestError(prepareErr.Detail)
	}
	return errors.NewInternalServerError("Server error")
}
