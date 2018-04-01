package initialization

import (
	"errors"
)

func InitByFlag(value string) error {

	switch value {
	case "db":
		return InitDb()
	}

	return errors.New("this value is not defined")
}
