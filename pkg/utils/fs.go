package utils

import (
	"crypto/md5"
	"fiber-layout/validator/form"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strconv"
)

// GetDirData 获取目录下文件/文件夹
func GetDirDataList(path string) ([]form.ListResponse, error) {
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

/*
效验文件后缀
extName：文件后缀
extMap：效验后缀
返回值：布尔值
*/

func allExtMap(extMap []string, extName string) bool {
	allowExtMap := make(map[string]bool)
	for _, val := range extMap {
		allowExtMap[val] = true
	}
	// 判断excel上传是否合法
	if _, ok := allowExtMap[extName]; !ok {

		return false
	}
	return true
}

/**
创建文件夹/文件名
extName：文件后缀
route：设置特定目录后缀
返回值：bool,路径
*/

func Mkdir(extName, route string) (error, string, string) {
	pwd, _ := os.Getwd()
	// 组成 文件路径
	dir := pwd + "/static/upload/" + GetFileDay() + route
	// 创建文件路径
	if err := os.MkdirAll(dir, 0666); err != nil {
		return err, "", ""
	}
	//生成文件名称   144325235235.png
	fileUnixName := strconv.FormatInt(GetUnixNano(), 10)
	saveDir := path.Join(dir, fileUnixName+"--"+extName)
	return nil, saveDir, fileUnixName + "--" + extName
}

/**
创建文件夹/文件名
extName：文件后缀
route：设置特定目录后缀
返回值：bool,路径
*/

func MkdirInfo(extName, route string) (error, string, string) {
	pwd, _ := os.Getwd()
	// 组成 文件路径
	dir := pwd + "/static" + route
	// 创建文件路径
	if err := os.MkdirAll(dir, 0666); err != nil {
		return err, "", ""
	}
	saveDir := path.Join(dir, extName)
	return nil, saveDir, extName
}

// 获取文件md5
func GetFileMd5(file *multipart.FileHeader) string {
	md5hash := md5.New()
	f, _ := file.Open()
	io.Copy(md5hash, f)
	has := md5hash.Sum(nil)
	return fmt.Sprintf("%x", has)
}
