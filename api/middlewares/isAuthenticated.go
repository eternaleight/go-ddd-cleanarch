package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")

		// エラーメッセージをまとめる関数
		unauthorized := func() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "権限がありません。"})
			c.Abort()
		}

		if authHeader == "" {
			unauthorized()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, bearerSchema)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			unauthorized()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if idFloat, ok := claims["id"].(float64); ok {
				userId := int(idFloat)
				c.Set("userID", userId)
				fmt.Println("UserID set in middleware:", userId) // この行を追加
				c.Next()
				return
			}
		}

		// 何らかの理由で認証が失敗した場合
		unauthorized()
	}
}
