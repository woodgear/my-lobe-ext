package calls

import (
	"time"
)

type CurrentTimeArgs struct {
	Ok bool `json:"ok"`
}
type CurrentTimeResponse struct {
	Now       string `json:"now"`
	DayOfWeek int    `json:"dayOfWeek"`
}

// 处理currentTime API
func HandleCurrentTime(arguments CurrentTimeArgs) (CurrentTimeResponse, error) {
	// 返回当前时间
	currentTime := time.Now()
	return CurrentTimeResponse{
		Now:       currentTime.Format("2006-01-02 15:04:05"),
		DayOfWeek: int(currentTime.Weekday()),
	}, nil
}
