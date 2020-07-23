package user

import (
	"weshierNext/model"
)

// RegisterReqStruct register request body struct
type RegisterReqStruct struct {
	Username string `zh:"账号" json:"username" binding:"required" validate:"required,lt=5"`
	Password string `zh:"密码" json:"password" binding:"required" validate:"required,lt=5"`
}

// LoginReqStruct login request body struct
type LoginReqStruct struct {
	UserName string `zh:"账号" json:"userName" binding:"required"`
	PassWord string `zh:"密码" json:"passWord" binding:"required"`
}

// LoginResStruct login response body struct
type LoginResStruct struct {
	model.UserAuth
	model.UserModel
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}

// GithubAuthRedirectStruct github 登录，响应的 code 结构
type GithubAuthRedirectStruct struct {
	Code  string `json:"code"`
	State uint32 `json:"state"`
}

// GithubAccessTokenRedirectStruct github 登录根据 code 去获取 access_token 请求的响应 body 结构
type GithubAccessTokenRedirectStruct struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// GithubUserInfoStruct github 登录后拿到的 user 信息结构
type GithubUserInfoStruct struct {
	ID                uint64 `json:"id"`
	UserName          string `json:"login"`
	Avatar            string `json:"avatar_url"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog"`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Bio               string `json:"bio"`
	PublicRepos       uint32 `json:"public_repos"`
	PublicGists       uint32 `json:"public_gists"`
	Followers         uint32 `json:"followers"`
	Following         uint32 `json:"following"`
}
