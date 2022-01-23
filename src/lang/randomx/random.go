package randomx

import (
	"math/rand"
	"time"
)

// spy 2022/1/21

const NUMBER_AND_LETTER = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// String 生成n个字符
func String(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = NUMBER_AND_LETTER[rand.Intn(len(NUMBER_AND_LETTER))]
	}
	return string(b)
}

//Alphabet 生成n个字母
func Alphabet(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = LETTERS[rand.Intn(len(LETTERS))]
	}
	return string(b)
}

// Number 生成count个[start,end)结束的不重复的随机数
func Number(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
