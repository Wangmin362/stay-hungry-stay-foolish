package main

func main() {

}

func firstUploadLog(logs []int) (res int) {
	score := make([][2]int, len(logs)) // 记录每个时刻，上报日志以及不上报日志的得分
	var prefixLogSum int
	for i := 0; i < len(logs); i++ {
		if prefixLogSum > 100 {
			break
		}
		if i == 0 {
			score[i][0] = logs[i] // 上报日志的得分
			score[i][1] = 0
			prefixLogSum += logs[i] // 累加之前所有未上报的日志
			continue
		}

		score[i][1] = score[i-1][1] + prefixLogSum         // 当前时刻所有没有上报的日志，全部扣分
		score[i][0] = logs[i] + prefixLogSum - score[i][1] // 上前的日志条数，以及之前所有没有上报的日志条数
		prefixLogSum += logs[i]
	}

	maxVal := 0
	for i := 0; i < len(score); i++ {
		maxVal = max(maxVal, score[i][0])
	}
	return maxVal
}
