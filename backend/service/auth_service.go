package service

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yrk9/cardboard_task_app/model"
	"github.com/yrk9/cardboard_task_app/repository"
	"os"
	"time"
	"errors"
	"strings"
	"unicode/utf8"
	"strconv"
)

const minPasswordLength = 8

var (
	ErrJWTToken = errors.New("ハッシュの生成エラー")
	ErrEmailNil = errors.New("メールアドレスが入力されていません")
	ErrEmailType = errors.New("メールアドレスの型が正しくありません")
	ErrPasswordLength = errors.New("パスワードの長さは" + strconv.Itoa(minPasswordLength) + "文字以上にしてください")
    ErrUserAlreadyExists  = errors.New("このメールアドレスは既に登録されています")
    ErrInvalidCredentials = errors.New("メールアドレスまたはパスワードが正しくありません")
)

type AuthService interface {
	Register(email, password string) (string, error)
	Login(email, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

//JWTを作る
func generateJWT(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        return "", ErrJWTToken
    }
	return signedToken, err
}


func (a* authService) Register(email, password string) (string, error) {
	newUser := &model.User{}

	// Emailのバリデーション
	if email == "" {
		return "", ErrEmailNil 
	} else if !strings.Contains(email, "@") {
		return "", ErrEmailType
	}
	// パスワードのバリデーション
	if utf8.RuneCountInString(password) < minPasswordLength {
		return "", ErrPasswordLength
	}
	//Emailの重複チェック
	existing, _ := a.userRepo.FindByEmail(email)
	if existing != nil {
		return "", ErrUserAlreadyExists
	}
	//パスワードのハッシュ化
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
    	return "", err
	}
	//新規ユーザの作成
	newUser.Email = email
	newUser.PasswordHash = string(hashed)
	a.userRepo.CreateUser(newUser)
	//JWTの生成
	return generateJWT(newUser.ID)
}

func (a* authService) Login(email, password string) (string, error) {
	//emailがあるかどうかの確認
	user, _ := a.userRepo.FindByEmail(email)
	if user == nil {
		return "", ErrInvalidCredentials
	}
	//パスワード照合
	err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", ErrInvalidCredentials
	}
	return generateJWT(user.ID)
}
