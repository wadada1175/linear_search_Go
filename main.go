package main

import (
	"fmt"
	"time"
)

func main() {
	var SIZE float64 = 60
	data := []int{
		101, 103, 107, 109, 113, 127, 131, 137, 139, 149,
		151, 157, 163, 167, 173, 179, 181, 191, 193, 197,
		199, 211, 223, 227, 229, 233, 239, 241, 251, 257,
		263, 269, 271, 277, 281, 283, 293, 307, 311, 313,
		317, 331, 337, 347, 349, 353, 359, 367, 373, 379,
		383, 389, 397, 401, 409, 419, 421, 431, 433, 439,
	}
	var sum_time float64 = 0
	for i := 0; i < len(data); i++ {
		target := data[i]
		start := time.Now() //時間計測開始
		for j := 0; j < len(data); j++ {
			if data[j] == target {
				end := time.Now() //時間計測終了
				fmt.Printf("値 %d は配列の %d 番目に見つかりました。\n", target, j)
				t := end.Sub(start).Seconds()
				fmt.Printf("処理時間は%.9fs\n", t)
				sum_time += t
			}
		}
	}
	ave_time := sum_time / SIZE
	fmt.Printf("平均処理時間は%.9fs", ave_time)
}
