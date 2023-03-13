package models

import (
	"fiber-layout/initalize"

	"gorm.io/gorm"
)

type FileInfo struct {
	gorm.Model
	Name string //文件名称
	Path string //文件路径
	Size int    //文件大小
	md5  string //文件标识（每个文件都有自己独特的md5）
}

func NewFileInfo() *FileInfo {
	return &FileInfo{}
}

func (t *FileInfo) GetList() ([]FileInfo, error) {
	var sys []FileInfo
	if err := initalize.DB.Raw("select * from Course LIMIT 10").Find(&sys).Error; err != nil {
		return nil, err
	}
	return sys, nil
}

func (t *FileInfo) Category(id string) (*FileInfo, error) {
	if err := initalize.DB.Raw("select * from Course WHERE CId = ? LIMIT 10", id).Find(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}
