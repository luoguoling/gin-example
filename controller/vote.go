package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/models"
)

func PostVoteHandler(c *gin.Context) {
	//接收参数
	//userid  postid  vote
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		fmt.Println(ok)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		zap.L().Debug("controller.vote c.ShouldBindJson(p) error", zap.Any("err", err))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	//userID, err := getCurrentUserId(c)
	//if err != nil {
	//	ResponseError(c, CodeNeedLogin)
	//}

	////处理逻辑
	//err := logic.PostVote(userID, p)
	//if err != nil {
	//	zap.L().Error("logic.PostVote failed ", zap.Error(err))
	//}
	//返回结果
	ResponseSuccess(c, nil)
}
