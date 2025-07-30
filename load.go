package main

import(
	"os"
	"io"
	"fmt"
)

// readAudioFile 讀取音頻文件並轉為 float32 數組
func LoadAudioFile(audioPath string)([]float32, error) {
	file, err := os.Open(audioPath)
	if err != nil {
		return nil, fmt.Errorf("打開文件失敗: %w", err)
	}
	defer file.Close()
	// 這裡假設音頻文件是 WAV 格式，16kHz，16位，單聲道，在實際應用中，您可能需要使用音頻解析庫來處理不同格式
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("讀取文件失敗: %w", err)
	}
	headerSize := 44  // WAV 文件頭通常是44字節
	if len(data) < headerSize {
		return nil, fmt.Errorf("文件太小，不是有效的WAV文件")
	}
	audioBytes := data[headerSize:]
	samples := make([]float32, len(audioBytes)/2)
	// 假設音頻數據是16位PCM格式，將其轉換為 float32
	for i := 0; i < len(samples); i++ {   // 將16位PCM數據轉換為float32
		sample := int16(audioBytes[i*2]) | int16(audioBytes[i*2+1])<<8   // 小端序讀取16位整數
		samples[i] = float32(sample) / 32768.0   
	}
	return samples, nil
}