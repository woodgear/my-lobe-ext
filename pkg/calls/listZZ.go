package calls

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type ListZZArgs struct {
	Ok bool `json:"ok"`
}
type ListZZResponse struct {
	Data string `json:"data"`
}

func HandleListZZ(arguments ListZZArgs) (ListZZResponse, error) {
	zzzsPath := os.Getenv("ZZZS_PATH")
	log.Printf("zzzs %v", zzzsPath)
	content, err := exec.Command("tree", "--du", "-h", zzzsPath).Output()
	if err != nil {
		return ListZZResponse{}, fmt.Errorf("failed to execute command: %s %v", content, err)
	}

	return ListZZResponse{
		Data: string(content),
	}, nil
}
