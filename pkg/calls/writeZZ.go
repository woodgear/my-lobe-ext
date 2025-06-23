package calls

import (
	"fmt"
	"os"
	"path/filepath"
)

type WriteZZArgs struct {
	Path string `json:"path"`
	Data string `json:"data"`
}
type WriteZZResponse struct {
	Ok bool `json:"ok"`
}

func HandleWriteZZ(arguments WriteZZArgs) (WriteZZResponse, error) {
	zzzsPath := os.Getenv("ZZZS_PATH")
	filePath := zzzsPath + "/" + arguments.Path
	os.MkdirAll(filepath.Dir(filePath), 0755)
	err := os.WriteFile(filePath, []byte(arguments.Data), 0644)
	if err != nil {
		return WriteZZResponse{}, fmt.Errorf("failed to write file %s: %v", filePath, err)
	}

	return WriteZZResponse{
		Ok: true,
	}, nil
}
