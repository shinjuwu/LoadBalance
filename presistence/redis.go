package presistence

import (
	"github.com/garyburd/redigo/redis"
	"github.com/liangdas/mqant/utils"
)

var uri string = "redis://:@127.0.0.1:6379"

func GetRedisConn() redis.Conn {
	return utils.GetRedisFactory().GetPool(uri).Get()
}
