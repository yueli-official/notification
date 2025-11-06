package model

// 通用成功响应
type SuccessResponse struct {
	Message string `json:"message" example:"操作成功"`
}

// 通用错误响应
type ErrorResponse struct {
	Error string `json:"error" example:"请求参数错误: 缺少字段 subject"`
}

// 批量发送响应
type BatchResponse struct {
	Message string   `json:"message" example:"批量发送完成,部分失败"`
	Errors  []string `json:"errors,omitempty" example:"[\"邮箱格式错误\", \"手机号无效\"]"`
	Total   int      `json:"total" example:"20"`
	Failed  int      `json:"failed,omitempty" example:"2"`
}
