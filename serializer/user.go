package serializer

import "clicli/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}

// LoginResult 登录结果结构
type LoginResult struct {
	User
	Token string `json:"token"`
}

// LoginResultResponse 登录结果序列化
type LoginResultResponse struct {
	Response
	Data LoginResult `json:"data"`
}

// BuildLoginResult 序列化用户
func BuildLoginResult(result LoginResult) LoginResult {
	return LoginResult{
		User:  result.User,
		Token: result.Token,
	}
}

// BuildLoginResultResponse 序列化用户响应
func BuildLoginResultResponse(result LoginResult) LoginResultResponse {
	return LoginResultResponse{
		Data: BuildLoginResult(result),
	}
}
