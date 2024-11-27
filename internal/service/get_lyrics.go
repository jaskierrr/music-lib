package service

import (
	"context"
	"log"
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

	log.Println(string([]rune{reverseSolidus, runeN, reverseSolidus, runeN}))


	lyricsData := []rune(lyricsStr)

	lyricsArr := []string{}
	lyricsNode := lyricsData[:4]
	newCouplet := []rune{reverseSolidus, runeN, reverseSolidus, runeN}

	for i := 4; i < len(lyricsData); i++ {

		log.Println("AAAAAAAAAAAAAAAAAAAAAAAAA")
		log.Println(string(lyricsNode))
		log.Println(string(lyricsData[i-4:i]))
		log.Println(string(newCouplet))

		if reflect.DeepEqual(lyricsData[i-4:i], newCouplet) {
			log.Println("NEW LINE")
			log.Println(string(lyricsNode))
			lyricsArr = append(lyricsArr, string(lyricsNode))
			lyricsNode = []rune{}
		}

		lyricsNode = append(lyricsNode, lyricsData[i])
	}

	log.Println(lyricsArr)
	lyricsArr = append(lyricsArr, string(lyricsNode))

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
