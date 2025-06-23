package calls

import (
	"fmt"
	"log"
	"os"
)

type ReadZZArgs struct {
	Paths []string `json:"paths"`
}
type ReadZZResponse struct {
	Data map[string]string `json:"data"`
}

func HandleReadZZ(arguments ReadZZArgs) (ReadZZResponse, error) {
	zzzsPath := os.Getenv("ZZZS_PATH")
	response := ReadZZResponse{
		Data: map[string]string{},
	}
	log.Printf("readzz %v", arguments)
	for _, path := range arguments.Paths {

		filePath := zzzsPath + "/" + path
		content, err := os.ReadFile(filePath)
		if err != nil {
			return ReadZZResponse{}, fmt.Errorf("failed to read file %s: %v", filePath, err)
		}
		response.Data[path] = string(content)
	}
	return response, nil
}
