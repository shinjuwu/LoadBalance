package route

import (
	"github.com/gin-gonic/gin"
)

type LSCommandLayer struct {
	Sys      string      `json:"sys"`
	Cmd      string      `json:"cmd"`
	Sn       int         `json:"sn"`
	IsEncode bool        `json:"isEncode"`
	Data     interface{} `json:"data"`
}

type LSRespLayer struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Ret     string      `json:"ret"`
	Data    interface{} `json:"data"`
}

type LSIPGet struct {
	PlatformID    int    `json:"platformId"`
	MemberCode    int    `json:"memberCode"`
	AgentID       int    `json:"agentId"`
	GameID        int    `json:"gameId"`
	GameCode      string `json:"gameCode"`
	Token         string `json:"token"`
	Account       string `json:"account"`
	Password      string `json:"password"`
	ClientType    string `json:"clientType"`
	ClientVersion string `json:"clientVersion"`
	GameTicket    string `json:"gameTicket"`
	ServerIdx     int    `json:"serverIdx"`
}

type IPGetRespData struct {
	ServerID   int    `json:"serverId"`
	IP         string `json:"ip"`
	ServerPort int    `json:"serverPort"`
	GameTicket string `json:"gameTicket"`
	ResUrl     string
}

type LSmemberLogin struct {
	ServerID            int    `json:"serverId"`
	PlatformID          int    `json:"platformId"`
	ThirdPartyUserID    int64  `json:"thirdPartyUserId"`
	ThirdPartyUserIDStr string `json:"thirdPartyUserIdStr"`
	TotalUser           int    `json:"totalUser"`
	Token               string `json:"token"`
	GameTicket          string `json:"gameTicket"`
	Account             string `json:"account"`
	Password            string `json:"password"`
	GameID              int    `json:"gameId"`
	GameCode            string `json:"gameCode"`
	ClientType          string `json:"clientType"`
	ClientVersion       string `json:"clientVersion"`
	IsLogin             bool   `json:"isLogin"`
}

type MemberLoginResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Ret     string      `json:"ret"`
	Data    interface{} `json:"data"`
}

type memberLoginResData struct {
	PlatformID          int    `json:"platformId"`
	MemberCode          int    `json:"memberCode"`
	AgentID             int    `json:"agentId"`
	UserID              int    `json:"user_id"`
	ThirdPartyUserID    int64  `json:"thirdPartyUserId"`
	ThirdPartyUserIDStr string `json:"thirdPartyUserIdStr"`
	Account             string `json:"account"`
	Nickname            string `json:"nickname"`
	GameTicket          string `json:"gameTicket"`
	Token               string `json:"token"`
	Balance             int64  `json:"balance"`
	BalanceCI           int64  `json:"balance_ci"`
	BalanceWin          int64  `json:"balance_win"`
	VIPRank             int    `json:"vip_rank"`
	LoginTime           string `json:"loginTime"`
	CreateTime          string `json:"createTime"`
	UpdateTime          string `json:"updateTime"`
	Status              int    `json:"status"`
	OrderState          int    `json:"orderState"`
}

type LSMemberLogout struct {
	ServerID            int    `json:"serverId"`
	PlatformID          int    `json:"platformId"`
	UserID              int    `json:"user_id"`
	ThirdPartyUserID    int    `json:"thirdPartyUserId"`
	ThirdPartyUserIDStr string `json:"thirdPartyUserIdStr"`
	TotalUser           int    `json:"totalUser"`
	Token               string `json:"token"`
	GameTicket          string `json:"gameTicket"`
	Account             string `json:"account"`
	Password            string `json:"password"`
	GameID              int    `json:"gameId"`
	GameCode            string `json:"gameCode"`
	ClientType          string `json:"clientType"`
	ClientVersion       string `json:"clientVersion"`
	IsLogin             bool   `json:"isLogin"`
}

type MemberLogoutResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Ret     string              `json:"ret"`
	Data    memberLogoutResData `json:"data"`
}

type memberLogoutResData struct {
	PlatformID          int    `json:"platformId"`
	MemberCode          int    `json:"memberCode"`
	AgentID             int    `json:"agentId"`
	UserID              int    `json:"user_id"`
	ThirdPartyUserID    int64  `json:"thirdPartyUserId"`
	ThirdPartyUserIDStr string `json:"thirdPartyUserIdStr"`
	Account             string `json:"account"`
	Nickname            string `json:"nickname"`
	GameTicket          string `json:"gameTicket"`
	Token               string `json:"token"`
	Balance             int64  `json:"balance"`
	BalanceCI           int64  `json:"balance_ci"`
	BalanceWin          int64  `json:"balance_win"`
	VIPRank             int    `json:"vip_rank"`
	LoginTime           string `json:"loginTime"`
	CreateTime          string `json:"createTime"`
	UpdateTime          string `json:"updateTime"`
	Status              int    `json:"status"`
	OrderState          string `json:"orderState"`
}

func RegisterRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/LoadBalance", loadbalanceController)
	return router
}
