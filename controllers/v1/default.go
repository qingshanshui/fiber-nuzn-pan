package v1

import (
	"fiber-layout/controllers"
	"fiber-layout/initalize"
	"fiber-layout/service"
	"fiber-layout/validator"
	"fiber-layout/validator/form"
	"github.com/gofiber/fiber/v2"
)

type DefaultController struct {
	controllers.Base
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

func (t *DefaultController) List(c *fiber.Ctx) error {
	// 初始化参数结构体
	ListRequestForm := form.ListRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &ListRequestForm); err != nil {
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
