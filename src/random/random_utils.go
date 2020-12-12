package random

import (
	"log"

	"github.com/google/uuid"
)

func GetUuid() string {
	uuidObject, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	return uuidObject.String()
}
