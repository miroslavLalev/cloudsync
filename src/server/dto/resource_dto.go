package dto

type ResourceDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewResourceDto(id, name, rtype string) *ResourceDto {
	return &ResourceDto{Id: id, Name: name, Type: rtype}
}
