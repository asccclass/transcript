module github.com/asccclass/transcript

go 1.24.1

require (
	github.com/ggerganov/whisper.cpp/bindings/go v0.0.0-20250728100232-d0a9d8c7f8f7
	github.com/joho/godotenv v1.5.1
)

// 如果需要本地路徑
replace github.com/ggerganov/whisper.cpp => ../whisper.cpp

// replace github.com/ggerganov/whisper.cpp/bindings/go => /path/to/whisper.cpp/bindings/go
