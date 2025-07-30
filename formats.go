package main

import (
	"fmt"
	"strings"
	whisper "github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

// 轉換為純文字格式
func formatTXT(model whisper.Model) (string) {
	var result strings.Builder
	/*
	for i := 0; i < model.NumSegments(); i++ {
		segment, err := model.GetSegment(i)
		if err != nil {
			continue
		}
		result.WriteString(segment.Text + " ")
	}
		*/
	result.WriteString("this is a test")
	return strings.TrimSpace(result.String())
}

// 轉換為 VTT 格式
func formatVTT(segments []Segment) (string) {
	var result strings.Builder
	result.WriteString("WEBVTT\n\n")
	for _, segment := range segments {
		result.WriteString(fmt.Sprintf("%s --> %s\n", 
			formatTimestamp(segment.Start), 
			formatTimestamp(segment.End)))
		result.WriteString(fmt.Sprintf("%s\n\n", segment.Text))
	}
	return result.String()
}

// 轉換為 TSV 格式
func formatTSV(segments []Segment) (string) {
	var result strings.Builder
	result.WriteString("start\tend\ttext\n")
	for _, segment := range segments {
		result.WriteString(fmt.Sprintf("%.3f\t%.3f\t%s\n", 
			segment.Start, segment.End, segment.Text))
	}
	return result.String()
}