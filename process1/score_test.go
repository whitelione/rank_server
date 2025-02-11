package process1

import (
	"math/rand"
	"testing"
	"time"
)

func TestScoreWithTs(t *testing.T) {
	num := 100
	for i := 0; i < num; i++ {
		score := rand.Intn(num)
		ts := time.Now().Unix()
		scoreWithTs := ScoreWithTs(int64(score), ts)
		parseScore := ParseScore(scoreWithTs)
		if score != int(parseScore) {
			t.Fatalf("分数和时间戳编解码错误, 编码前分数: %d, 解码后分数: %d", score, parseScore)
		}
	}
	t.Log("分数和时间戳编解码测试 ok")
}
