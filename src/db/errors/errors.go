package errors

import (
	"fmt"
)

type SpaceAlreadyExists struct{}

func (e SpaceAlreadyExists) Error() string {
	return "Space already exists"
}

type NoSpaceExists struct {
	id string
}

func (e NoSpaceExists) Error() string {
	return fmt.Sprintf("No space with id '%v' exists", e.id)
}
