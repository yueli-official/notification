package middleware

import (
	"net/http"
	"notification/model"

	"github.com/gin-gonic/gin"
)

const apiKeyHeader = "X-API-Key"
const apiKeyCookie = "api_key"

// APIKeyAuth 验证请求头或 Cookie 中的 API Key
func APIKeyAuth(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader(apiKeyHeader)
		if key == "" {
			key, _ = c.Cookie(apiKeyCookie)
		}

		if key == "" || key != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Error: "未授权：API Key 无效或缺失",
			})
			return
		}

		c.Next()
	}
}
