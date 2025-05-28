package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/asticode/go-astisub"
	"github.com/mattn/go-jsd"
)

func extractSubtitleContent(sub *astisub.Subtitles) string {
	var builder strings.Builder
	for _, item := range sub.Items {
		builder.WriteString(item.String())
	}

	return builder.String()
}

func CompareSubtitle(reader1, reader2 io.Reader) (float64, error) {
	sub1, err := astisub.ReadFromWebVTT(reader1)
	if err != nil {
		return 0, fmt.Errorf("ファイル1の読み込みエラー: %v", err)
	}

	sub2, err := astisub.ReadFromWebVTT(reader2)
	if err != nil {
		return 0, fmt.Errorf("ファイル2の読み込みエラー: %v", err)
	}

	content1 := extractSubtitleContent(sub1)
	content2 := extractSubtitleContent(sub2)

	return jsd.StringDistance(content1, content2), nil
}

func openFile(filePath string) (io.ReadCloser, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("ファイル %s が見つかりません", filePath)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ファイル %s を開けません: %v", filePath, err)
	}
	return file, nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("使用方法: ./webvtt-jsd <file1> <file2>")
	}
	filePath1 := args[0]
	filePath2 := args[1]

	file1, err := openFile(filePath1)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	file2, err := openFile(filePath2)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	distance, err := CompareSubtitle(file1, file2)
	if err != nil {
		log.Fatalf("エラー: %v", err)
	}

	fmt.Println(distance)
}
