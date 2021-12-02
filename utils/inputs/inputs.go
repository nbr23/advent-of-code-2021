package inputs

import (
	"os"
	"strconv"
	"strings"
)

func GetInput(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

func StrListToIntList(input []string) []int64 {
	intlist := make([]int64, 0, len(input))
	for _, s := range input {
		i, _ := strconv.ParseInt(s, 10, 64)
		intlist = append(intlist, i)
	}
	return intlist
}

func GetInputInt64(filename string) []int64 {
	return StrListToIntList(strings.Split(GetInput(filename), "\n"))
}
