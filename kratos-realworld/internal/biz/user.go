package biz

import (
	"context"
	"errors"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string // 数据库不能明文存储密码
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

// --------------------------------------------------------------------------------------------------
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur   UserRepo
	pr   ProfileRepo
	jwtc *conf.JWT
	log  *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, jwtc *conf.JWT, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, jwtc: jwtc, log: log.NewHelper(logger)}
}

// --------------------------------------------------------------------------------------------------
func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	// 1. 参数校验 TODO
	if len(username) < 3 || len(username) > 20 {
		return nil, errors.New("username length must be between 3 and 20")
	}
	if len(email) < 3 || len(email) > 20 {
		return nil, errors.New("email length must be between 3 and 20")
	}
	if len(password) < 4 || len(password) > 20 {
		return nil, errors.New("password length must be between 4 and 20")
	}

	// TODO 用户是否存在 邮箱是否重复

	// 2. 创建用户 (加密密码)
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	// 3. 返回用户信息
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    uc.generateToken(username),
		Bio:      "Nice to meet you",
		Image:    "default-avatar.png",
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	// 1. 查询用户
	// u, err := uc.ur.GetUserByEmail(ctx, email)
	// if err != nil {
	// 	return nil, err
	// }

	// 2. 验证密码
	// verifyPassword(u.PasswordHash, password)
	// if !verifyPassword(u.PasswordHash, password) {
	// 	return nil, errors.New("login failed")
	// }

	// 3. 返回用户信息
	return &UserLogin{
		// Email:    u.Email,
		// Username: u.Username,
		// Token:    uc.generateToken(u.Username),
		// Bio:      u.Bio,
		// Image:    u.Image,
		Username: "test",
	}, nil
}

// --------------------------------------------------------------------------------------------------
func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func verifyPassword(hashedPwd, inputPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd)); err != nil {
		return false
	}
	return true
}

func (uc *UserUsecase) generateToken(username string) string {
	return auth.GenerateToken(uc.jwtc.Token, username)
}
