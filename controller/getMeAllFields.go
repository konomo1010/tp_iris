package controller

import (
	"github.com/kataras/iris/v12"
	"tp_iris/model"
)

func GetMeAllFields(ctx iris.Context) {
	emotion_info := model.ME_Field{}
	res := emotion_info.Queryall()
	ctx.JSON(res)
}
