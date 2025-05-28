package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	sameContentVtt := `WEBVTT

00:00:22.230 --> 00:00:24.606
これは最初の字幕です。

00:00:30.739 --> 00:00:35.743
これは 2 番目です。これは 3 番目です。
`

	distance, err := CompareSubtitle(strings.NewReader(vtt), strings.NewReader(sameContentVtt))
	assert.NoError(t, err)
	assert.Equal(t, 0.0, distance)
}
