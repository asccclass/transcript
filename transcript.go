package main

import(
	"os"
	"fmt"
	// "time"
	// "strconv"
	// "strings"
	// "context"
	"github.com/joho/godotenv"
	whisper "github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

func run(modelName, audioPath string)(error) {
   modelPath, err := downloadModel(modelName)
   if err != nil {
       fmt.Println("Error downloading model:", err)
       return err
   }
   model, err := whisper.New(modelPath)
   if err != nil {
       fmt.Println("Error loading model:", err)
       return err
   }
	defer model.Close()
	// 載入音訊檔
	audioData, err := LoadAudioFile(audioPath)
	if err != nil {
		fmt.Println("Error loading audio file:", err)
		return err
	}
	
	context, err := model.NewContext()
	if err != nil {
		fmt.Println("Error creating context:", err)
		return err
	}


	// 設置語言
	lang := os.Getenv("Language")
	if lang == "" {
		lang = "auto"  // 默認為自動檢測語言
	}
	if err := context.SetLanguage(lang); err != nil {
		fmt.Println("Error setting language:", err)
		return err
	}
	context.ResetTimings()

	/*
	// 配置 Whisper 參數，專門針對中文優化
	params := whisper.Params{
		Strategy:      whisper.SAMPLING_GREEDY, // 貪婪採樣策略
		Language:      lang,                    // 設定語言為中文
		Translate:     false,                   // 不翻譯，保持原語言
		Offset:        0 * time.Millisecond,
		Duration:      0 * time.Millisecond, // 處理整個音頻
		MaxTokens:     0,                     // 不限制token數量
		Threads:       4,                     // 使用4個線程
		SpeedsUp:      false,
		AudioCtx:      0,
		TokenThreshold: 0.01, // token閾值
		Entropy:       2.4,   // 熵閾值
		LogProb:       -1.0,  // 對數概率閾值
		NoContext:     false, // 使用上下文
		SingleSegment: false, // 不限制為單一段落
		PrintSpecial:  false, // 不打印特殊token
		PrintProgress: true,  // 顯示進度
		PrintRealtime: false, // 不實時打印
		PrintTimestamps: true, // 打印時間戳
	}
*/		
  // 建立處理上下文

/*
	context.SetSplitOnWord(true)  // 啟用分段
	MaxSegment := 0                // 默認為不限制段落長度
	maxseg := os.Getenv("MaxSegment")
	if maxseg != "" && maxseg != "0" {
		var err error
	   MaxSegment, err = strconv.Atoi(maxseg)
		if err != nil {
			fmt.Println("Invalid MaxSegment value, using default 0")
			return err
		}
		context.SetDuration(time.Duration(MaxSegment) * time.Second)
	}
*/		
	// 執行轉錄
	err = context.Process(audioData, nil, func(segment whisper.Segment) {
		
	}, nil)
	if err != nil {
		fmt.Println("Error processing audio:", err)
		return err
	}
	// fullText := formatTXT(model)  // 獲取完整文字
	// fmt.Errorf("轉錄完成，共 %d 段，總時長 xx 秒\n%s", model.NumSegments(), duration, fullText)
   return nil
}

func main() {
	fmt.Println("Starting transcription...")
	if err := godotenv.Load("envfile"); err != nil {
      fmt.Println(err.Error())
      return
   }
	audioPath := "D:\\transcript\\ex01.wav" // 替換為實際的音訊檔案路徑
	if audioPath == "" {
		fmt.Println("No audio file specified")
		return
	}
	// 設置轉錄參數
	modelName := os.Getenv("ModelName")
	if err := run(modelName, audioPath); err != nil {
		fmt.Println("Error during transcription:", err)
		return
	}
	fmt.Println("Transcription completed successfully")
	os.Exit(0)
}