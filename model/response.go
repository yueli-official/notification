package model

// EmailRequest 邮件请求参数
type EmailRequest struct {
	Provider string   `json:"provider" example:"account"` // 指定邮件提供商
	To       []string `json:"to" binding:"required" example:"user@example.com"`
	Subject  string   `json:"subject" binding:"required" example:"欢迎使用系统"`
	Body     string   `json:"body" binding:"required" example:"你好，欢迎使用我们的平台"`
	IsHTML   bool     `json:"is_html" example:"true"`
}

// SMSRequest 短信请求参数
type SMSRequest struct {
	Provider    string            `json:"provider" example:"tencent"`
	PhoneNumber string            `json:"phone_number" binding:"required" example:"13800138000"`
	TemplateID  string            `json:"template_id" binding:"required" example:"SMS_123456"`
	Params      map[string]string `json:"params,omitempty" example:"{\"code\":\"1234\"}"`
}

// BatchRequest 批量发送请求参数
type BatchRequest struct {
	Emails []EmailRequest `json:"emails"`
	SMS    []SMSRequest   `json:"sms"`
}
