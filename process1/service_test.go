package process1

import (
	"math/rand"
	"sort"
	"testing"
)

// 服务接口测试用例
func TestService(t *testing.T) {
	IoInit(RedisAddr)
	// 随机生成玩家数据
	playerNum := 20
	rankInfo := GenPlayerData(20)

	// 更新玩家分数
	svr := &LeaderboardService{}
	for _, v := range rankInfo {
		if err := svr.UpdateScore(v.PlayerId, v.Score, v.Timestamp); err != nil {
			t.Fatalf("UpdateScore err: %v", err)
		}
	}

	// 计算正确排名数据
	sort.Slice(rankInfo, func(i, j int) bool {
		if rankInfo[i].Score != rankInfo[j].Score {
			return rankInfo[i].Score > rankInfo[j].Score
		}
		return rankInfo[i].Timestamp < rankInfo[j].Timestamp
	})
	t.Log("------正确排名数据:start------")
	for i, v := range rankInfo {
		v.Rank = int64(i) + 1
		t.Log(v.ToString())
	}
	t.Log("------正确排名数据:end------")

	// 获取前n名玩家排名信息
	n := playerNum / 2
	topNRankInfo, err := svr.GetTopN(int64(n))
	if err != nil {
		t.Fatalf("GetTopN err: %v", err)
	}
	for i, v := range topNRankInfo {
		if v.PlayerId != rankInfo[i].PlayerId || v.Rank != rankInfo[i].Rank {
			t.Fatalf("获取前n名玩家排名信息错误, 当前排名: %d, 正确排名: %d, 玩家id: %s\n", v.Rank, rankInfo[i].Rank, v.PlayerId)
		}
	}
	t.Log("获取前n名玩家排名信息 ok")

	// 获取随机玩家当前排名
	index := rand.Intn(playerNum)
	playerId := rankInfo[index].PlayerId
	t.Logf("获取排名玩家索引: %d, 玩家id: %s\n", index, playerId)
	rank, err := svr.GetPlayerRank(playerId)
	if err != nil {
		t.Fatalf("GetPlayerRank err: %v", err)
	}
	correctRank := rankInfo[index].Rank
	if rank != correctRank {
		t.Fatalf("获取玩家当前排名错误, 当前排名: %d, 正确排名: %d\n", rank, correctRank)
	}
	t.Log("获取玩家当前排名 ok")

	// 查询自己名次前后共 N 名玩家的分数和名次
	n = playerNum / 3
	rangeRankInfo, err := svr.GetPlayerRankRange(playerId, int64(n))
	if err != nil {
		t.Fatalf("GetPlayerRankRange err: %v", err)
	}
	t.Logf("------玩家id: %s 前后共 %d 名玩家的排名信息:start------", playerId, n)
	for _, v := range rangeRankInfo {
		t.Log(v.ToString())
	}
	t.Logf("------玩家id: %s 前后共 %d 名玩家的排名信息:end------", playerId, n)
	t.Log("查询自己名次前后共 N 名玩家的分数和名次 ok")
}
