# webvtt-jsd

WebVTTファイル間の類似度を計算するツールです。Jaro-Winkler距離を使用して、2つの字幕ファイルの内容の類似度を数値で示します。

## インストール

```bash
go install github.com/kumackey/webvtt-jsd@latest
```

## 使い方

```bash
webvtt-jsd <file1> <file2>
```

### 例

```bash
# 同じ内容のファイルを比較
webvtt-jsd testdata/sample.vtt testdata/sample_same_content.vtt
# 出力: 0

# タイプミスがあるファイルを比較
webvtt-jsd testdata/sample.vtt testdata/sample_typo.vtt
# 出力: 0.09898989898989896

# 完全に異なる内容のファイルを比較
webvtt-jsd testdata/sample.vtt testdata/sample_errors.vtt
# 出力: 0.6372053872053871
```
