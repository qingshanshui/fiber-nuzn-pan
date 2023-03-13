package v1

import (
	"errors"
	"fiber-layout/controllers"
	"fiber-layout/initalize"
	"fiber-layout/models"
	"fiber-layout/pkg/utils"
	"fiber-layout/service"
	"fiber-layout/validator"
	"fiber-layout/validator/form"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DefaultController struct {
	controllers.Base
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

func (t *DefaultController) GetList(c *fiber.Ctx) error {
	// 初始化参数结构体
	ListRequestForm := form.ListRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &ListRequestForm); err != nil {
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().GetList(ListRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(api))
}

func (t *DefaultController) GetFile(c *fiber.Ctx) error {
	// 初始化参数结构体
	GetRequestForm := form.GetRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &GetRequestForm); err != nil {
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().GetFile(GetRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(api))
}

func (t *DefaultController) Download(c *fiber.Ctx) error {
	// 初始化参数结构体
	DownloadRequestForm := form.DownloadRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &DownloadRequestForm); err != nil {
		return err
	}
	pwd, _ := os.Getwd()
	url := pwd + "/static" + DownloadRequestForm.Path
	exists, err := utils.PathExists(url)
	if err != nil {
		return err
	}
	if exists {
		return c.SendFile(url)
	}
	return errors.New("文件错误")
}
func (t *DefaultController) Login(c *fiber.Ctx) error {
	// 初始化参数结构体
	DownloadRequestForm := form.LoginRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &DownloadRequestForm); err != nil {
		return err
	}
	api, err := service.NewDefaultService().Login(DownloadRequestForm)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(api))
}

func (t *DefaultController) Upload(c *fiber.Ctx) error {
	// 接收文件file
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	// 获取md5 值
	md5 := utils.GetFileMd5(file)
	// 查 md5 是否存在库
	fi := models.NewFileInfo()
	FileInfos, err := fi.Md5Verify(md5)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	if len(FileInfos) != 0 {
		return c.JSON(t.Ok(FileInfos))
	}

	if c.Query("type") == "1" && c.Query("type") == "2" {
		return c.JSON(t.Fail(errors.New("参数错误")))
	}
	var pathDir = ""  // 文件路径
	var FileName = "" //文件名

	//  api上传
	if c.Query("type") == "1" {
		// 拼接文件路径
		err, pathDir, FileName = utils.Mkdir(file.Filename, "")
		if err != nil {
			return c.JSON(t.Fail(err))
		}
	}
	// 上传到当前目录
	if c.Query("type") == "2" {
		// 拼接文件路径
		err, pathDir, FileName = utils.MkdirInfo(file.Filename, c.Query("url"))
		if err != nil {
			return c.JSON(t.Fail(err))
		}
	}

	// 保存文件
	if err := c.SaveFile(file, pathDir); err != nil {
		return c.JSON(t.Fail(err))
	}
	fi.CreatedAt = time.Now()
	fi.Md5 = md5
	fi.Name = FileName
	fi.Path = pathDir
	fi.Size = int(file.Size)
	if err := fi.Create(); err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(pathDir))
}
