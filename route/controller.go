package route

import (
	"LoadBalance/presistence"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/liangdas/mqant/log"
)

func loadbalanceController(context *gin.Context) {
	sendData := context.PostForm("sendData")
	commandLayer := LSCommandLayer{}
	err := json.Unmarshal([]byte(sendData), &commandLayer)
	if err != nil {
		panic(err)
	}
	//data := commandLayer.Data.(map[string]interface{})
	switch commandLayer.Cmd {
	case "loadBalanceHttpIpGet":
		ipGetResp := IPGetRespData{
			ServerID:   10005,
			GameTicket: "1111-2222-3333-4444",
			IP:         "127.0.0.1",
		}
		resp := LSRespLayer{
			Code:    0,
			Message: "Success",
			Ret:     commandLayer.Cmd,
			Data:    ipGetResp,
		}
		responseMessage(context, resp)
		break
	case "loadBalanceHttpServerAdd":
		resp := LSRespLayer{
			Code:    0,
			Message: "Success",
			Ret:     commandLayer.Cmd,
			Data:    "",
		}
		responseMessage(context, resp)
		break
	case "loadBalanceHttpMemberLogin":
		respMemLogin := memberLogoutResData{
			ThirdPartyUserID: 1234,
			UserID:           999,
			Nickname:         "shinjuwu",
			Status:           0,
			Balance:          200000000,
		}

		resp := LSRespLayer{
			Code:    0,
			Message: "Success",
			Ret:     commandLayer.Cmd,
			Data:    respMemLogin,
		}
		responseMessage(context, resp)
		break
	case "loadBalanceHttpMemberLogout":
		resp := LSRespLayer{
			Code:    0,
			Message: "Success",
			Ret:     commandLayer.Cmd,
			Data:    "",
		}
		responseMessage(context, resp)
	default:
		resp := LSRespLayer{
			Code:    0,
			Message: "Success",
			Ret:     commandLayer.Cmd,
			Data:    "",
		}
		responseMessage(context, resp)
	}

}

func responseMessage(context *gin.Context, response interface{}) {
	s, ok := response.(string)
	if ok {
		rawdata := json.RawMessage(s)
		response = rawdata
	}
	byteData, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	context.String(http.StatusOK, string(byteData))
}

func getGosIP(context *gin.Context) {
	// token := context.Query("token")
	// gameCode := context.Query("gameCode")
	// key := token + ":" + gameCode
	// byteData, err := Query(key)
	// if err != nil {
	// 	panic(err)
	// }
	context.String(http.StatusOK, string("twgsg1.dreamtech8.com"))
}

func addNewServer(context *gin.Context) {
	conn := presistence.GetRedisConn()
	defer conn.Close()
	serverList, err := redis.Bytes(conn.Do("SET", "ServerList"))
	if err != nil {
		fmt.Println("Get server list error!")
	}
	fmt.Println(string(serverList))
}

func Query(key string) ([]byte, error) {
	conn := presistence.GetRedisConn()
	defer conn.Close()
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Error("redis Get failure, key : %s , Error: %v", key, err)
		return nil, err
	}
	return data, nil
}
