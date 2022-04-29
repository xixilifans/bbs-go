package api

import (
	"bbs-go/controllers/render"
	"bbs-go/model"
	"bbs-go/pkg/common"
	"bbs-go/pkg/msg"
	"bbs-go/services"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web"
)

type MessageController struct {
	Ctx iris.Context
}

//私信列表
func (c *MessageController) GetMessages() *web.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	if user == nil {
		return web.JsonError(common.ErrorNotLogin)
	}

	var req struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}
	if err := c.Ctx.ReadJSON(&req); err != nil {
		return web.JsonError(common.ErrorRequestParams)
	}

	messages, paging := services.MessageService.GetMessageList(user.Id, req.Page, req.Limit)
	// if err != nil {
	// 	return web.JsonError(common.ErrorGetMessages)
	// }
	if messages == nil {
		return web.JsonError(common.ErrorGetMessages)
	}

	return web.JsonPageData(render.BuildMessages(messages), paging)
}

//私信
func (c *MessageController) PostSendMessage() *web.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	if user == nil {
		return web.JsonError(common.ErrorNotLogin)
	}
	var req struct {
		ToUserId int64  `json:"toUserId"`
		Content  string `json:"content"`
	}
	if err := c.Ctx.ReadJSON(&req); err != nil {
		return web.JsonError(common.ErrorRequestParams)
	}
	quoteContent := fmt.Sprintf("with %s private message", req.ToUserId)

	if err := services.MessageService.SendMsg(user.Id, req.ToUserId, msg.TypePrivateMessage,
		"private message", req.Content, quoteContent, nil); err != nil {
		return web.JsonError(common.ErrorSendMessage)
	}

	return web.JsonSuccess()
}

//私信删除
func (c *MessageController) PostDeleteMessage() *web.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	if user == nil {
		return web.JsonError(common.ErrorNotLogin)
	}
	var req struct {
		Id int64 `json:"id"`
	}
	if err := c.Ctx.ReadJSON(&req); err != nil {
		return web.JsonError(common.ErrorRequestParams)
	}
	if err := services.MessageService.DeleteMsg(user.Id, req.Id); err != nil {
		return web.JsonError(common.ErrorDeleteMessage)
	}

	return web.JsonSuccess()
}

//私信已读
func (c *MessageController) PostReadMessage() *web.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	if user == nil {
		return web.JsonError(common.ErrorNotLogin)
	}
	var req struct {
		Id int64 `json:"id"`
	}
	if err := c.Ctx.ReadJSON(&req); err != nil {
		return web.JsonError(common.ErrorRequestParams)
	}
	if err := services.MessageService.ReadMsg(user.Id, req.Id); err != nil {
		return web.JsonError(common.ErrorReadMessage)
	}

	return web.JsonSuccess()
}

//私信未读数
func (c *MessageController) GetUnreadMessageCount() *web.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	if user == nil {
		return web.JsonError(common.ErrorNotLogin)
	}
	var count int64 = 0
	var messages []model.Message
	if user != nil {
		count = services.MessageService.GetUnReadCount(user.Id)
		messages = services.MessageService.Find(sqls.NewCnd().Eq("user_id", user.Id).
			Eq("status", msg.StatusUnread).Desc("id"))
	}
	return web.NewEmptyRspBuilder().Put("count", count).Put("messages", render.BuildMessages(messages)).JsonResult()
}
