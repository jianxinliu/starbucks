package utils

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"starbucks/starbucks/global"
	"strconv"
)

type UserAuthRequest struct {
	Token      string
	AuthSecret string
}

type UserAuthResponse struct {
	UserId string `json:"userId"`
}

const UserTokenExpire = "token_expire:%s"

func UserAuth(in *UserAuthRequest) (*UserAuthResponse, error) {
	//in.Token
	token, err := jwt.Parse(in.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.UserAuth.Secret), nil
	}, jwt.WithJSONNumber())
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token claims failed")
	}
	userId := claims["userId"]
	if userId == "" {
		return nil, errors.New("token error")
	}
	iat, err := claims["iat"].(json.Number).Int64()
	if err != nil {
		return nil, errors.New("token iat error")
	}

	// redis中存储了强制过期时间，需要跟token生成时间做对比
	expireStr, _ := global.Redis.Get(fmt.Sprintf(UserTokenExpire, userId))
	if expireStr != "" {
		expire, err := strconv.ParseInt(expireStr, 10, 64)
		if err == nil && expire != 0 && expire > iat {
			return nil, jwt.ErrTokenExpired
		}
	}

	return &UserAuthResponse{
		UserId: userId.(string),
	}, nil
}
