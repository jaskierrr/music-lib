package service

import (
	"context"
	"main/api/restapi/operations"
	"reflect"
	"strings"
)

const reverseSolidus = rune(0x005C)
const runeN = rune('n')

func (s service) GetLyrics(ctx context.Context, params operations.GetSongsLyricsParams) (string, error) {
	lyricsStr, err := s.repo.GetLyrics(ctx, params)
	if err != nil || lyricsStr == "" {
		return "", err
	}

	lyricsData := []rune(lyricsStr)

	lyricsArr := []string{}
	lyricsNode := lyricsData[:4]
	newCouplet := []rune{reverseSolidus, runeN, reverseSolidus, runeN}

	for i := 4; i < len(lyricsData); i++ {

		if reflect.DeepEqual(lyricsData[i-4:i], newCouplet) {
			lyricsArr = append(lyricsArr, string(lyricsNode))
			lyricsNode = []rune{}
		}

		lyricsNode = append(lyricsNode, lyricsData[i])
	}

	lyricsArr = append(lyricsArr, string(lyricsNode))

	lenArr := int64(len(lyricsArr))

	if *params.Couplet >= lenArr {
		return "", err
	}
	if *params.Limit >= lenArr {
		*params.Limit = lenArr
	}

	lyrics := ""

	if *params.Couplet == 0 && *params.Limit == 0 {
		lyrics = strings.Join(lyricsArr, "")
	}
	if *params.Couplet != 0 && *params.Limit != 0 {
		lyrics = strings.Join(lyricsArr[*params.Couplet-1:*params.Couplet+*params.Limit-1], "")
	}
	if *params.Couplet != 0 && *params.Limit == 0 {
		lyrics = strings.Join(lyricsArr[*params.Couplet-1:], "")
	}
	if *params.Couplet == 0 && *params.Limit != 0 {
		lyrics = strings.Join(lyricsArr[:*params.Limit], "")
	}

	lyrics = strings.ReplaceAll(string(lyrics), "\\n", "\n")

	return lyrics, nil
}
