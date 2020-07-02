package main

import (
	"fmt"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func main() {
	myOptions := levenshtein.Options{
		InsCost: 1,
		DelCost: 1,
		SubCost: 1, // DefaultOptionsでは2
		Matches: levenshtein.IdenticalRunes,
	}
	source := "mail"
	target := "googlemap"
	distance := levenshtein.DistanceForStrings([]rune(source), []rune(target), myOptions)
	fmt.Printf("Distance between %s and %s computed as %d\n", source, target, distance)
}

// sourcesの中から、targetに最も類似している単語を返す関数
// target, sources共に事前に小文字か空白除去等の前処理をしておく
func getSimilarWord(target string, sources []string) string {
	// レーベンシュタイン距離を計算する際の重みづけ
	// 削除の際の距離を小さくしている
	myOptions := levenshtein.Options{
		InsCost: 10,
		DelCost: 1,
		SubCost: 10,
		Matches: levenshtein.IdenticalRunes,
	}
	distances := make([]int, len(sources))
	for i := 0; i < len(sources); i++ {
		distances[i] = levenshtein.DistanceForStrings([]rune(sources[i]), []rune(target), myOptions)
	}

	// 距離が最小のもののインデックスを取得する
	minIdx := 0
	minDistance := 1000000 // 十分大きな数
	for i := 0; i < len(distances); i++ {
		if distances[i] <= minDistance {
			minDistance = distances[i]
			minIdx = i
		}
	}

	return sources[minIdx]
}
