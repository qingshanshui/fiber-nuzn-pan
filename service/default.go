package service

import (
	"fiber-layout/pkg/utils"
	"fiber-layout/validator/form"
)

type Default struct {
}

func NewDefaultService() *Default {
	return &Default{}
}

func (t *Default) GetList(list form.ListRequest) ([]form.ListResponse, error) {
	data, err := utils.GetDirData(list.Path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
