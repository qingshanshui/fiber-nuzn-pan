package utils

import (
	"fiber-layout/validator/form"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// GetDirData 获取目录下文件/文件夹
func GetDirData(path string) ([]form.ListResponse, error) {
	var l []form.ListResponse
	pwd, _ := os.Getwd()
	url := pwd + "/static" + path
	exists, err := PathExists(url)
	if err != nil {
		return nil, err
	}
	if exists {
		//获取文件或目录相关信息
		fileInfoList, err := ioutil.ReadDir(url)
		if err != nil {
			log.Fatal(err)
		}
		for i := range fileInfoList {
			l = append(l, form.ListResponse{
				Path:  path + "/" + fileInfoList[i].Name(),
				IsDir: fileInfoList[i].IsDir(),
				Time:  fileInfoList[i].ModTime().Format("2006-01-02 15:04:05"),
				Name:  fileInfoList[i].Name(),
				Size:  fileInfoList[i].Size(),
			})
		}
		return l, err
	}
	return nil, nil
}

// GetDirFile 获取目录下文件信息
func GetDirFile(path string) ([]form.GetResponse, error) {
	var l []form.GetResponse
	pwd, _ := os.Getwd()
	url := pwd + "/static" + path
	exists, err := PathExists(url)
	if err != nil {
		return l, err
	}
	fmt.Println(exists, "文件是否存在")
	if exists {
		//获取文件或目录相关信息
		fileInfoList, err := os.Stat(url)
		if err != nil {
			return l, err
		}
		l = append(l, form.GetResponse{
			Path:  path,
			IsDir: fileInfoList.IsDir(),
			Time:  fileInfoList.ModTime().Format("2006-01-02 15:04:05"),
			Name:  fileInfoList.Name(),
			Size:  fileInfoList.Size(),
		})
		return l, err
	}
	return nil, nil
}

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
