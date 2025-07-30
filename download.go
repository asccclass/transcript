package main

import(
	"os"
	"io"
	"fmt"
	"strings"
	"net/http"
   "path/filepath"
)

func getModelFromHuggingFace(path, url string) (error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("正在下載模型 %s...", url)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("下載失敗，狀態碼: %d", resp.StatusCode)
		}
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}
		fmt.Printf("模型 %s 下載完成", path)
	}
	return nil
}


// 下載模型檔案
func downloadModel(modelName string) (string, error) {
	dir, err := getCurrentDirectory()
	if err != nil {	
		return "", err
	}
	modelPath := filepath.Join(dir, "models") // 目前目錄下的 models 資料夾，不存在則建立該目錄
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		if err := os.Mkdir(modelPath, 0755); err != nil {
			return "", err
		}
	}
	// 模型下載 URL 映射
	modelURLs := map[string]string{
		"ggml-tiny.bin":     "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-tiny.bin",
		"ggml-base.bin":     "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-base.bin",
		"ggml-small.bin":    "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-small.bin",
		"ggml-medium.bin":   "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-medium.bin",
		"ggml-large-v1.bin": "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-large-v1.bin",
		"ggml-large-v2.bin": "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-large-v2.bin",
		"ggml-large-v3.bin": "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-large-v3.bin",
	}
	// 檢查模型名稱是否在映射中
	modelName = strings.ToLower(modelName)
	if modelName == "" {  // 預設採用最小的
		modelName = "tiny"
	}
	modelName = "ggml-" + modelName + ".bin"
	url, exists := modelURLs[modelName]
	if !exists {
		return "", fmt.Errorf("不支援的模型: %s", modelName)
	}
	if err := getModelFromHuggingFace(modelPath + "/" + modelName, url); err != nil {
		return "", err
	}
	return modelPath + "/" + modelName, nil
}