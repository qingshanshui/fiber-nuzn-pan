package service

import (
	"errors"
	"fiber-layout/pkg/utils"
	"fiber-layout/validator/form"

	"github.com/spf13/viper"
)

type Default struct {
}

func NewDefaultService() *Default {
	return &Default{}
}

func (t *Default) GetList(list form.ListRequest) ([]form.ListResponse, error) {
	data, err := utils.GetDirDataList(list.Path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t *Default) GetFile(list form.GetRequest) ([]form.GetResponse, error) {
	data, err := utils.GetDirFile(list.Path)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (t *Default) Login(login form.LoginRequest) (string, error) {
	if login.Username == "13122256420" && login.Password == "123456" {
		token, err := utils.CreateToken(login.Username, viper.GetString("Jwt.Secret"))
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", errors.New("密码错误")
}
