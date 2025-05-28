package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareSubtitle_WithTypo(t *testing.T) {
	vtt := `WEBVTT

1
00:00:22.230 --> 00:00:24.606
これは最初の字幕です。

2 Some Text
00:00:30.739 --> 00:00:34.074
これは 2 番目です。

3
00:00:34.159 --> 00:00:35.743
これは 3 番目です。
`

	typo := `WEBVTT

1
00:00:22.230 --> 00:00:24.606
これは会社の字幕です。

2 Some Text
00:00:30.739 --> 00:00:34.074
これは 2 番目です。

3
00:00:34.159 --> 00:00:35.743
これはアンダンテです。
`

	distance, err := CompareSubtitle(strings.NewReader(vtt), strings.NewReader(typo))
	assert.NoError(t, err)
	assert.Greater(t, distance, 0.0, "タイプミスがある字幕との距離は0より大きくなるはず")
	assert.Less(t, distance, 0.1, "タイプミスがある字幕との距離は0.1未満になるはず")
}

func TestCompareSubtitle_SameContent(t *testing.T) {
	vtt := `WEBVTT

1
00:00:22.230 --> 00:00:24.606
これは最初の字幕です。

2 Some Text
00:00:30.739 --> 00:00:34.074
これは 2 番目です。

3
00:00:34.159 --> 00:00:35.743
これは 3 番目です。
`

	same := `WEBVTT

00:00:22.230 --> 00:00:24.606
これは最初の字幕です。

00:00:30.739 --> 00:00:35.743
これは 2 番目です。これは 3 番目です。
`

	distance, err := CompareSubtitle(strings.NewReader(vtt), strings.NewReader(same))
	assert.NoError(t, err)
	assert.Equal(t, 0.0, distance)
}

func TestCompareSubtitle_DifferentContent(t *testing.T) {
	vtt := `WEBVTT

1
00:00:22.230 --> 00:00:24.606
これは最初の字幕です。

2 Some Text
00:00:30.739 --> 00:00:34.074
これは 2 番目です。

3
00:00:34.159 --> 00:00:35.743
これは 3 番目です。
`

	different := `WEBVTT

1
00:00:22.230 --> 00:00:24.606
お電話ありがとうございます。

2 Some Text
00:00:30.739 --> 00:00:34.074
株式会社 ABC でございます。

3
00:00:34.159 --> 00:00:35.743
どのようなご用件でしょうか。
`

	distance, err := CompareSubtitle(strings.NewReader(vtt), strings.NewReader(different))
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, distance, 0.5, "完全に異なる内容の字幕との距離は0.5以上になるはず")
}
