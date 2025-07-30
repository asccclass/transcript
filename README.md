

# ffmpeg 轉換
```
ffmpeg -i ex01.wav -ar 16000 -ac 1 -c:a pcm_s16le ex00.wav
```

# Prompt
```
使用go-whisper 建立GO STT 程式，功能如下：
1、能夠自動下載需要的模型檔案
2、有多語言辨識功能
3、要能分段處理，並輸出 srt、tsv、txt、vtt、json格式的文字檔
4、包裝成 HTTP API 服務
```