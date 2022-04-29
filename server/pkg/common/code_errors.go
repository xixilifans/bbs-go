package common

import (
	"github.com/mlogclub/simple/web"
)

var (
	ErrorNotLogin              = web.NewError(1, "请先登录")
	CaptchaError               = web.NewError(1000, "验证码错误")
	ForbiddenError             = web.NewError(1001, "已被禁言")
	UserDisabled               = web.NewError(1002, "账号已禁用")
	InObservationPeriod        = web.NewError(1003, "账号尚在观察期")
	EmailNotVerified           = web.NewError(1004, "请先验证邮箱")
	ErrorRequestParams         = web.NewError(1005, "请求参数错误")
	ErrorGetMessages           = web.NewError(1006, "获取私信列表失败")
	ErrorSendMessage           = web.NewError(1007, "发送私信失败")
	ErrorDeleteMessage         = web.NewError(1008, "删除私信失败")
	ErrorReadMessage           = web.NewError(1009, "标记私信为已读失败")
	ErrorGetUnreadMessageCount = web.NewError(1010, "获取未读私信数量失败")
)
