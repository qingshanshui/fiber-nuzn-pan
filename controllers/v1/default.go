package v1

import (
	"errors"
	"fiber-layout/controllers"
	"fiber-layout/initalize"
	"fiber-layout/pkg/utils"
	"fiber-layout/service"
	"fiber-layout/validator"
	"fiber-layout/validator/form"
	"os"

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
		return c.JSON(t.Fail(err))
	}
	initalize.Log.Info(api)

	return c.JSON(t.Ok(api)) // => ✋ register
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
		return c.JSON(t.Fail(err))
	}
	initalize.Log.Info(api)

	return c.JSON(t.Ok(api)) // => ✋ register
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
	return c.JSON(t.Ok(api)) // => ✋ register
}

func (t *DefaultController) Upload(c *fiber.Ctx) error {
	// 接收文件file
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	//// 获取文件后缀
	//extName := path.Ext(file.Filename)
	// 拼接文件路径
	err, pathDir := utils.Mkdir(file.Filename, "")
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	// 保存文件
	if err := c.SaveFile(file, pathDir); err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(pathDir))
}
