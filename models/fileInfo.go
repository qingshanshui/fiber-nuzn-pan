package models

import (
	"fiber-layout/initalize"

	"gorm.io/gorm"
)

type FileInfo struct {
	gorm.Model
	Name string `gorm:"comment:文件名称" json:"name"`                //文件名称
	Path string `gorm:"comment:文件路径" json:"path"`                //文件路径
	Size int    `gorm:"comment:文件大小" json:"size"`                //文件大小
	Md5  string `gorm:"comment:文件标识（每个文件都有自己独特的md5）" json:"md5"` //文件标识（每个文件都有自己独特的cmd5）
}

func NewFileInfo() *FileInfo {
	return &FileInfo{}
}

// md5 验证
func (t *FileInfo) Md5Verify(cmd5 string) ([]FileInfo, error) {
	var sys []FileInfo
	if err := initalize.DB.Raw("SELECT * FROM file_infos WHERE md5 = ? LIMIT 10", cmd5).Find(&sys).Error; err != nil {
		return nil, err
	}
	return sys, nil
}

// 存储文件数据
func (t *FileInfo) Create() error {
	if err := initalize.DB.Exec("INSERT INTO file_infos (file_infos.created_at,file_infos.`name`,file_infos.path,file_infos.size,file_infos.md5) VALUES(?,?,?,?,?)", t.CreatedAt, t.Name, t.Path, t.Size, t.Md5).Error; err != nil {
		return err
	}
	return nil
}
