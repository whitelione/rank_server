package process1

import "math"

const maxTs = 9999999999

var scoreExpandRatio = math.Pow10(10)

func ScoreWithTs(score, ts int64) (scoreWithTs int64) {
	return score*int64(scoreExpandRatio) + maxTs - ts
}

func ParseScore(scoreWithTs int64) (score int64) {
	return scoreWithTs / int64(scoreExpandRatio)
}
