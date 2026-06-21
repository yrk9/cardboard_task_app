package middleware

import (
    "context"
    "errors"
    "net/http"
    "os"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

type contextKey string
const userIDKey contextKey = "user_id"

func GetUserID(r *http.Request) (int64, bool) {
    userID, ok := r.Context().Value(userIDKey).(int64)
    return userID, ok
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			// "Bearer " が無かった（除去されなかった）= 不正
			http.Error(w, "認証が必要です", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不正な署名方式")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			http.Error(w, "認証に失敗しました", http.StatusUnauthorized)
			return   // ← 忘れずに
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "認証に失敗しました", http.StatusUnauthorized)
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)   // JWTの数値は float64
		if !ok {
			http.Error(w, "認証に失敗しました", http.StatusUnauthorized)
			return
		}
		userID := int64(userIDFloat)

		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}