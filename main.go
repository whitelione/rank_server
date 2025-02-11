package main

import (
	"log"
	"math/rand"
	"rank_server/process1"
)

func main() {
	// 初始化io
	if err := process1.IoInit(process1.RedisAddr); err != nil {
		log.Fatalf("IoInit err: %v", err)
	}

	// 生成排名数据
	playerNum := 20
	rankInfo := process1.GenPlayerData(20)
	svr := &process1.LeaderboardService{}
	for _, v := range rankInfo {
		if err := svr.UpdateScore(v.PlayerId, v.Score, v.Timestamp); err != nil {
			log.Fatalf("UpdateScore err: %v", err)
		}
	}

	// 获取玩家当前排名
	index := rand.Intn(playerNum)
	playerId := rankInfo[index].PlayerId
	rank, err := svr.GetPlayerRank(playerId)
	if err != nil {
		log.Fatalf("GetPlayerRank err: %v", err)
	}
	log.Printf("获取排名玩家id: %s, 玩家排名: %d\n", playerId, rank)

	// 获取前n名玩家排名信息
	n := playerNum / 4
	topNRankInfo, err := svr.GetTopN(int64(n))
	if err != nil {
		log.Fatalf("GetTopN err: %v", err)

	}
	log.Printf("获取前 %d 名玩家排名信息: start\n", n)
	for _, v := range topNRankInfo {
		log.Println(v.ToString())
	}
	log.Printf("获取前 %d 名玩家排名信息: end\n", n)

	// 查询指定用户名次前后共 N 名玩家的分数和名次
	rangeRankInfo, err := svr.GetPlayerRankRange(playerId, int64(n))
	if err != nil {
		log.Fatalf("GetPlayerRankRange err: %v", err)
	}
	log.Printf("玩家id: %s 前后共 %d 名玩家的排名信息:start", playerId, n)
	for _, v := range rangeRankInfo {
		log.Println(v.ToString())
	}
	log.Printf("玩家id: %s 前后共 %d 名玩家的排名信息:end", playerId, n)
}
