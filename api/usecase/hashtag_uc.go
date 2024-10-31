package usecase

import (
	"regexp"
)

func ExtractHashtags(content string) []string {
	// ハッシュタグの正規表現パターン
	re := regexp.MustCompile(`#\p{L}[\p{L}\p{N}_]*`)
	// contentからハッシュタグをすべて抽出
	hashtags := re.FindAllString(content, -1)
	return hashtags
}