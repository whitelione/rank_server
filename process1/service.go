package process1

import "fmt"

type LeaderboardServiceI interface {
	// 更新玩家分数
	// zadd
	UpdateScore(playerId string, score, timestamp int64) error

	// 获取玩家当前排名
	// zrevrank
	GetPlayerRank(playerId string) (rank int64, err error)

	// 获取排行榜前N名
	// zrevrange(0, N-1)
	GetTopN(n int64) ([]*RankInfo, error)

	// 查询自己名次前后共N名玩家的分数和名次
	// 由于这个接口需求文档中没有给出明确的如何在前后取N名玩家,所以说明下我的实现逻辑:
	// 1.N为非1的奇数时前后各取N/2个玩家,加上自己共N名玩家,若前面或后面玩家不足N/2个,则有多少取多少,可能总数不足N
	// 2.N为偶数时前面取N/2个玩家,后面取N/2-1名玩家,加上自己共N名玩家,若前面或后面玩家不足N/2个,则有多少取多少,可能总数不足N
	// 3.N为1时只取自己
	GetPlayerRankRange(playerId string, n int64) ([]*RankInfo, error)
}

type LeaderboardService struct{}

func (s *LeaderboardService) UpdateScore(playerId string, score int64, timestamp int64) (err error) {
	newScore := ScoreWithTs(score, timestamp)
	err = SetScore(playerId, newScore)
	return
}

func (s *LeaderboardService) GetPlayerRank(playerId string) (rank int64, err error) {
	index, err := GetIndex(playerId)
	rank = index + 1
	return
}

func (s *LeaderboardService) GetTopN(n int64) (rankInfo []*RankInfo, err error) {
	rankInfo, err = GetTopN(n)
	return
}

func (s *LeaderboardService) GetPlayerRankRange(playerId string, n int64) (rankInfo []*RankInfo, err error) {
	index, err := GetIndex(playerId)
	if err != nil {
		return
	}
	start := index - n/2
	if start < 0 {
		start = 0
	}
	end := index + n/2
	if n%2 == 0 {
		end -= 1
	}
	if end < start {
		end = start
	}
	rankInfo, err = GetRnage(int(start), int(end))
	return
}

type RankInfo struct {
	// 业务使用字段
	PlayerId string
	Score    int64
	Rank     int64

	// 测试使用字段
	Timestamp int64 // 分数达成时间戳
}

func (r *RankInfo) ToString() string {
	return fmt.Sprintf("playerId: %s, rank: %d, score: %d, timestamp: %d", r.PlayerId, r.Rank, r.Score, r.Timestamp)
}
