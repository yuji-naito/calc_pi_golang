package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
#include <time.h>
*/
import "C"

func calcPi(testCount int) float64 {
	x := 0.0
	y := 0.0
	hitCount := 0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < testCount; i++ {
		// x座標、y座標を0.0～1.0の間でランダムに取得
		x = rand.Float64()
		y = rand.Float64()

		// 原点からの距離(x^2 + y^2)が1未満であれば扇の中に入ったものとしてhitCountをインクリメント
		if x*x+y*y < 1 {
			hitCount++
		}
	}

	// 4 * 円の扇の中に入った回数 / 全回数 = 円周率
	pi := 4 * float64(hitCount) / float64(testCount)

	return pi
}

func main() {
	testCount := 1000
	repeatCount := 6
	pi := 0.0

	for i := 0; i < repeatCount; i++ {
		testCount *= 10

		startClock := C.clock() // コンテキストスイッチ等の影響を考えてCPU時間を取得

		pi = calcPi(testCount)

		endClock := C.clock() // コンテキストスイッチ等の影響を考えてCPU時間を取得

		// 結果出力
		fmt.Printf("count: %10d\tpi: %.10f\ttime: %.10f\n", testCount, pi, float64(endClock-startClock)/float64(C.CLOCKS_PER_SEC))
	}

}
