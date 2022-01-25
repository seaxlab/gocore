package arrayx

import "math/rand"

// spy 2020/5/1

func ToStrArray(data interface{}) []string {
	return data.([]string)
}

func Shuffle(data []int) {
	for i := range data {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}
