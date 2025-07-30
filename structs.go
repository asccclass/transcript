package main

import(
	whisper "github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

type STTService struct {
	model whisper.Model
}

type Segment struct {
	ID    int     `json:"id"`
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	Text  string  `json:"text"`
}