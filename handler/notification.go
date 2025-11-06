package handler

import (
	"net/http"
	"notification/model"
	"notification/service"

	"github.com/gin-gonic/gin"
)

// ----------------------------
// 处理器定义
// ----------------------------

type NotificationHandler struct {
	service *service.NotificationService
}

func NewNotificationHandler(service *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

// SendEmail godoc
// @Summary 发送邮件
// @Description 发送单封邮件，可指定邮件提供商或使用默认启用的
// @Tags Notification
// @Accept json
// @Produce json
// @Param request body model.EmailRequest true "邮件发送请求参数"
// @Success 200 {object} model.SuccessResponse "邮件发送成功"
// @Failure 400 {object} model.ErrorResponse "请求参数错误"
// @Failure 500 {object} model.ErrorResponse "发送邮件失败"
// @Router /notification/email [post]
func (h *NotificationHandler) SendEmail(c *gin.Context) {
	var req model.EmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "请求参数错误: " + err.Error(),
		})
		return
	}

	if err := h.service.SendEmail(req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "发送邮件失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "邮件发送成功",
	})
}

// SendSMS godoc
// @Summary 发送短信
// @Description 发送单条短信，可指定短信服务商
// @Tags Notification
// @Accept json
// @Produce json
// @Param request body model.SMSRequest true "短信发送请求参数"
// @Success 200 {object} model.SuccessResponse "短信发送成功"
// @Failure 400 {object} model.ErrorResponse "请求参数错误"
// @Failure 500 {object} model.ErrorResponse "发送短信失败"
// @Router /notification/sms [post]
func (h *NotificationHandler) SendSMS(c *gin.Context) {
	var req model.SMSRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "请求参数错误: " + err.Error(),
		})
		return
	}

	if err := h.service.SendSMS(req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "发送短信失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "短信发送成功",
	})
}

// SendBatch godoc
// @Summary 批量发送通知
// @Description 批量发送邮件和短信，返回发送结果统计
// @Tags Notification
// @Accept json
// @Produce json
// @Param request body model.BatchRequest true "批量发送请求参数"
// @Success 200 {object} model.BatchResponse "批量全部成功"
// @Success 206 {object} model.BatchResponse "部分失败"
// @Failure 400 {object} model.ErrorResponse "请求参数错误"
// @Router /notification/batch [post]
func (h *NotificationHandler) SendBatch(c *gin.Context) {
	var req model.BatchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "请求参数错误: " + err.Error(),
		})
		return
	}

	errors := h.service.SendBatch(req)

	if len(errors) > 0 {
		errorMsgs := make([]string, len(errors))
		for i, err := range errors {
			errorMsgs[i] = err.Error()
		}

		c.JSON(http.StatusPartialContent, model.BatchResponse{
			Message: "批量发送完成,部分失败",
			Errors:  errorMsgs,
			Total:   len(req.Emails) + len(req.SMS),
			Failed:  len(errors),
		})
		return
	}

	c.JSON(http.StatusOK, model.BatchResponse{
		Message: "批量发送全部成功",
		Total:   len(req.Emails) + len(req.SMS),
	})
}
