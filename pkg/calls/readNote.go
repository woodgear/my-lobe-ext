package calls

import (
	"fmt"
	"log"
	"os"
	"time"
)

type ReadNoteArgs struct {
	NoteName     string `json:"noteName"`
	TodayJournal bool   `json:"todayJournal"`
	Journal      bool   `json:"journal"`
}
type ReadNoteResponse struct {
	Data string `json:"data"`
}

func HandleReadNote(arguments ReadNoteArgs) (ReadNoteResponse, error) {
	// 构建文件路径
	var filePath string
	if arguments.TodayJournal {
		filePath = os.Getenv("LOGSEQ_PATH") + "/journals/" + time.Now().Format("2006_01_02") + ".md"
	} else if arguments.Journal {
		filePath = os.Getenv("LOGSEQ_PATH") + "/journals/" + arguments.NoteName + ".md"
	} else {
		filePath = os.Getenv("LOGSEQ_PATH") + "/pages/" + arguments.NoteName + ".md"
	}
	log.Printf("Reading file: %s", filePath)

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		return ReadNoteResponse{}, fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	return ReadNoteResponse{
		Data: string(content),
	}, nil
}
