package process3

// 题目三密集排名接口实现说明
type LeaderboardServiceI interface {
	// 更新玩家分数接口和题目一保持一致,只是取消时间戳与分数结合逻辑
	UpdateScore(playerId string, score int64) error

	// 获取玩家当前排名,获取玩家在zset中的索引,遍历该索引之前所有玩家信息,使用 GetRankByDenseRnak 函数获取玩家密集排名的名次
	GetPlayerRank(playerId string) (rank int64, err error)

	// 获取排行榜前N名后使用 DenseRnak 函数计算密集排名,根据返回值判断最大名次是否为N,如果不够再向后遍历玩家,直到取到密集排名的前N名玩家
	GetTopN(n int64) ([]RankInfo, error)

	// 查询自己名次前后共N名玩家的分数和名次
	// 在题目一此接口逻辑基础上,获取自己排名,并向前和向后分别按密集排名取部分用户,直到满足前后共N名玩家
	GetPlayerRankRange(playerId string, n int64) ([]RankInfo, error)
}

// 在题目一排名算法基础上计算密集排名逻辑,由于更新分数接口取消时间戳与分数结合逻辑,所以无需考虑时间戳的影响
// 返回值为密集排名信息和当前最大名次
func DenseRnak(rankInfo []*RankInfo) (denseRankInfo []*RankInfo, rank int64) {
	rankScore := -1
	for _, v := range rankInfo {
		if rankScore == -1 || rankScore != int(v.Score) {
			rank++
			rankScore = int(v.Score)
		}
		v.Rank = int64(rank)
	}
	return
}

// 根据密集排名算法获取玩家排名,-1代表无玩家信息,由于更新分数接口取消时间戳与分数结合逻辑,所以无需考虑时间戳的影响
func GetRankByDenseRnak(rankInfo []*RankInfo, playerId string) (rank int64) {
	rankScore := -1
	for _, v := range rankInfo {
		if rankScore == -1 || rankScore != int(v.Score) {
			rank++
			rankScore = int(v.Score)
		}
		if v.PlayerId == playerId {
			return rank
		}
	}
	return -1
}

type RankInfo struct {
	// 业务使用字段
	PlayerId string
	Score    int64
	Rank     int64
}
