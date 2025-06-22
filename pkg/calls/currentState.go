package calls

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type CurrentStateArgs struct {
	Ok bool `json:"ok"`
}
type CurrentStateResponse struct {
	Data string `json:"data"`
}

func HandleCurrentState(arguments CurrentStateArgs) (CurrentStateResponse, error) {
	// 获取当前日期并格式化为YYYY_MM_DD
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()

	dateStr := strconv.Itoa(year) + "_" +
		fmt.Sprintf("%02d", month) + "_" +
		fmt.Sprintf("%02d", day)

	// 构建文件路径
	filePath := os.Getenv("LOGSEQ_PATH") + "/journals/" + dateStr + ".md"
	log.Printf("Reading file: %s", filePath)

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		return CurrentStateResponse{}, fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	return CurrentStateResponse{
		Data: string(content),
	}, nil
}
