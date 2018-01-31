package dto

import (
	"github.com/miroslavLalev/cloudsync/src/objects"
)

type SpaceDto struct {
	Id        string         `json:"id"`
	Owner     string         `json:"owner"`
	Resources []*ResourceDto `json:"resources"`
}

func NewSpace(s *objects.Space) *SpaceDto {
	return &SpaceDto{Id: s.Id(), Owner: s.Owner(), Resources: transformResources(s.Resources())}
}

func transformResources(resources []*objects.Resource) []*ResourceDto {
	result := []*ResourceDto{}
	for _, r := range resources {
		result = append(result, NewResource(r.Id(), r.Name(), string(r.Type())))
	}

	return result
}
