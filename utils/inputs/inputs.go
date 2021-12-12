package inputs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ParseDecInt(str string) (res int64) {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func getToken() string {
	token, err := os.ReadFile("./.token")
	if err != nil {
		panic(err)
	}
	return string(token)
}

func GetInput(day int) string {
	token := getToken()

	client := http.Client{}

	request, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day), nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Cookie", fmt.Sprintf("session=%s", token))

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return strings.Trim(string(data), "\n")
}

func StrListToIntList(input []string) []int64 {
	intlist := make([]int64, 0, len(input))
	for _, s := range input {
		i, _ := strconv.ParseInt(s, 10, 64)
		intlist = append(intlist, i)
	}
	return intlist
}

func InputToIntList(input string) []int64 {
	return StrListToIntList(strings.Split(input, "\n"))
}
