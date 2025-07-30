package main

import(
	"os"
	"fmt"
	"strings"
	"path/filepath"
)

// 取得目前目錄
func getCurrentDirectory() (string, error) {
   dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// 轉為絕對路徑，並補上結尾 /
	dir = filepath.Clean(dir)
	if !strings.HasSuffix(dir, string(os.PathSeparator)) {
		dir += string(os.PathSeparator)
	}
	return dir, nil
}

// 時間戳格式化
func formatTimestamp(seconds float64) string {
	hours := int(seconds) / 3600
	minutes := (int(seconds) % 3600) / 60
	secs := int(seconds) % 60
	millis := int((seconds - float64(int(seconds))) * 1000)
	return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, minutes, secs, millis)
}