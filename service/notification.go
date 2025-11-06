package service

import (
	"errors"
	"fmt"
	"notification/config"
	"notification/model"

	"gopkg.in/gomail.v2"
)

type NotificationService struct {
	config *config.Config
}

func NewNotificationService(cfg *config.Config) *NotificationService {
	return &NotificationService{
		config: cfg,
	}
}

// SendEmail 发送邮件
func (s *NotificationService) SendEmail(req model.EmailRequest) error {
	// 选择邮件提供商
	var provider *config.EmailProvider

	if req.Provider != "" {
		// 使用指定的提供商
		for i := range s.config.Email {
			if s.config.Email[i].Name == req.Provider && s.config.Email[i].Enabled {
				provider = &s.config.Email[i]
				break
			}
		}
		if provider == nil {
			return fmt.Errorf("未找到启用的邮件提供商: %s", req.Provider)
		}
	} else {
		// 使用第一个启用的提供商
		for i := range s.config.Email {
			if s.config.Email[i].Enabled {
				provider = &s.config.Email[i]
				break
			}
		}
		if provider == nil {
			return errors.New("没有可用的邮件提供商")
		}
	}

	// 创建邮件
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(provider.Username, provider.From))
	m.SetHeader("To", req.To...)
	m.SetHeader("Subject", req.Subject)

	if req.IsHTML {
		m.SetBody("text/html", req.Body)
	} else {
		m.SetBody("text/plain", req.Body)
	}

	// 发送邮件
	d := gomail.NewDialer(provider.Host, provider.Port, provider.Username, provider.Password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("发送邮件失败 (提供商: %s): %w", provider.Name, err)
	}

	return nil
}

// 发送短信
func (s *NotificationService) SendSMS(req model.SMSRequest) error {
	// 选择短信提供商
	var provider *config.SMSProvider

	if req.Provider != "" {
		for i := range s.config.SMS {
			if s.config.SMS[i].Name == req.Provider && s.config.SMS[i].Enabled {
				provider = &s.config.SMS[i]
				break
			}
		}
		if provider == nil {
			return fmt.Errorf("未找到启用的短信提供商: %s", req.Provider)
		}
	} else {
		for i := range s.config.SMS {
			if s.config.SMS[i].Enabled {
				provider = &s.config.SMS[i]
				break
			}
		}
		if provider == nil {
			return errors.New("没有可用的短信提供商")
		}
	}

	// 根据不同的提供商调用不同的API
	switch provider.Provider {
	case "aliyun":
		return s.sendAliyunSMS(provider, req)
	case "tencent":
		return s.sendTencentSMS(provider, req)
	default:
		return fmt.Errorf("不支持的短信提供商类型: %s", provider.Provider)
	}
}

// 批量发送
func (s *NotificationService) SendBatch(req model.BatchRequest) []error {
	var errors []error

	// 发送邮件
	for _, email := range req.Emails {
		if err := s.SendEmail(email); err != nil {
			errors = append(errors, fmt.Errorf("邮件发送失败: %w", err))
		}
	}

	// 发送短信
	for _, sms := range req.SMS {
		if err := s.SendSMS(sms); err != nil {
			errors = append(errors, fmt.Errorf("短信发送失败: %w", err))
		}
	}

	return errors
}

// 阿里云短信
func (s *NotificationService) sendAliyunSMS(provider *config.SMSProvider, req model.SMSRequest) error {
	fmt.Printf("使用阿里云发送短信到: %s, 模板: %s\n", req.PhoneNumber, req.TemplateID)

	// TODO: 阿里云短信发送逻辑
	// import "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	return nil
}

// 腾讯云短信发送
func (s *NotificationService) sendTencentSMS(provider *config.SMSProvider, req model.SMSRequest) error {
	fmt.Printf("使用腾讯云发送短信到: %s, 模板: %s\n", req.PhoneNumber, req.TemplateID)

	// TODO: 腾讯云短信发送逻辑
	// import "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

	return nil
}
