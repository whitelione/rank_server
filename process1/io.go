package process1

import (
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

var conn redis.Conn

const rankKey = "rank"

var RedisAddr = "127.0.0.1:6379"

func IoInit(addr string) (err error) {
	timeout := 2 * time.Second
	conn, err = redis.Dial("tcp", addr,
		redis.DialConnectTimeout(timeout),
		redis.DialWriteTimeout(timeout),
		redis.DialReadTimeout(timeout),
		redis.DialDatabase(0))
	return
}

// 设置分数
func SetScore(playerId string, score int64) (err error) {
	_, err = conn.Do("zadd", rankKey, score, playerId)
	return
}

// 获取排名
func GetIndex(playerId string) (index int64, err error) {
	index, err = redis.Int64(conn.Do("zrevrank", rankKey, playerId))
	return
}

// 获取前n名
func GetTopN(n int64) (rankInfo []*RankInfo, err error) {
	bts, err := redis.ByteSlices(conn.Do("zrevrange", rankKey, 0, n-1, "withscores"))
	if err != nil {
		return
	}
	for i := 0; i < len(bts)-1; i += 2 {
		playerId := string(bts[i])
		var score int
		if score, err = strconv.Atoi(string(bts[i+1])); err != nil {
			return
		}
		rankInfo = append(rankInfo, &RankInfo{
			PlayerId: playerId,
			Score:    ParseScore(int64(score)),
			Rank:     int64(i)/2 + 1,
		})
	}
	return
}

// 获取范围数据
func GetRnage(start, end int) (rankInfo []*RankInfo, err error) {
	bts, err := redis.ByteSlices(conn.Do("zrevrange", rankKey, start, end, "withscores"))
	if err != nil {
		return
	}
	index := start
	for i := 0; i < len(bts)-1; i += 2 {
		playerId := string(bts[i])
		var score int
		if score, err = strconv.Atoi(string(bts[i+1])); err != nil {
			return
		}
		rankInfo = append(rankInfo, &RankInfo{
			PlayerId: playerId,
			Score:    ParseScore(int64(score)),
			Rank:     int64(index) + 1,
		})
		index++
	}
	return
}
