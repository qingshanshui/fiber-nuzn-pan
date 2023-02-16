package middleware

import (
	"fiber-layout/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.JSON(fiber.Map{
			"code": 0,
			"data": "权限不足，禁止操作",
			"msg":  "操作失败",
		})
	} else {
		_, err := utils.ParseToken(token, viper.GetString("Jwt.Secret"))
		if err != nil {
			return ctx.JSON(fiber.Map{
				"code": 0,
				"data": "权限不足，禁止操作",
				"msg":  "操作失败",
			})
		} else {
			return ctx.Next()
		}
	}
}
