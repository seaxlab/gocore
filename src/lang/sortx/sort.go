package sortx

import (
	"sort"
)

// spy 2022/1/21

func StringAsc(data []string) {
	sort.Strings(data)
}

func StringDesc(data []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(data)))
}
