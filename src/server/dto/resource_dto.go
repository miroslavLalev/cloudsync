package dto

import (
	"github.com/miroslavLalev/cloudsync/src/objects"
)

type ResourceDto struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Resources []*ResourceDto `json:"resources"`
}

func NewResource(id, name, rtype string) *ResourceDto {
	return &ResourceDto{Id: id, Name: name, Type: rtype}
}

func NewResourceDetailed(r *objects.Resource) *ResourceDto {
	return &ResourceDto{
		Id:        r.Id(),
		Name:      r.Name(),
		Type:      string(r.Type()),
		Resources: transformResources(r.Resources()),
	}
}
