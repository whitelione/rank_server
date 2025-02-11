package process1

import (
	"math/rand"
	"strconv"
	"time"
)

// 生成 num 个玩家数据
func GenPlayerData(num int) (rankInfo []*RankInfo) {
	for i := 0; i < num; i++ {
		rankInfo = append(rankInfo, &RankInfo{
			PlayerId:  strconv.Itoa(num*10 + i),
			Score:     rand.Int63n(int64(num)),
			Timestamp: time.Now().Unix() + int64(i),
		})
	}
	return
}
