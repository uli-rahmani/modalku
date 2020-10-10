package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type DetailData struct {
	Len   int
	Value string
}

func main() {
	// Set stdIn.
	datas := bufio.NewReader(os.Stdin)
	fmt.Print("Enter len data: ")
	lenDatas, _ := datas.ReadString('\n')
	lenDatasInt := StringToInt(lenDatas)

	detailData := make(map[int]DetailData)
	for i := 0; i < lenDatasInt; i++ {
		data := bufio.NewReader(os.Stdin)
		fmt.Print("Enter len data ", i+1, " : ")
		lenData, _ := data.ReadString('\n')
		lenInt := StringToInt(lenData)

		values := bufio.NewReader(os.Stdin)
		fmt.Print("Enter value data ", i+1, " : ")
		valueData, _ := values.ReadString('\n')
		valueStr := SetString(valueData)

		detailData[i] = DetailData{
			Len:   lenInt,
			Value: valueStr,
		}
	}

	// Get Maximal displacement.
	result := GetMaxDisplacement(detailData, lenDatasInt)

	fmt.Println("\nResult: ")
	fmt.Println(strings.Join(result, "\n"))
}

func StringToInt(s string) int {
	data := strings.Split(s, "\n")

	i, err := strconv.Atoi(data[0])
	if err != nil {
		return 0
	}

	return i
}

func SetString(s string) string {
	data := strings.Split(s, "\n")
	return data[0]
}

func GetMaxDisplacement(data map[int]DetailData, lenData int) []string {
	result := make([]string, lenData)

	for key, value := range data {
		valueStr := strings.Split(value.Value, "")

		sumChar := GetSumChar(valueStr, "L", "R", "?")

		if sumChar[0] == 0 && sumChar[1] == 0 && sumChar[2] == 0 {
			result[key] = "0"
			continue
		}

		if sumChar[0] == sumChar[1] && sumChar[2] == 0 {
			result[key] = "0"
			continue
		}

		if sumChar[0] != sumChar[1] && sumChar[2] == 0 {
			result[key] = strconv.Itoa(int(math.Abs(float64(sumChar[0] - sumChar[1]))))
			continue
		}

		if sumChar[0] == 0 && sumChar[1] == 0 && sumChar[2] > 0 {
			result[key] = strconv.Itoa(sumChar[2])
			continue
		}

		// Last case.
		if sumChar[0] > sumChar[1] {
			sumChar[0] += sumChar[2]
		} else {
			sumChar[1] += sumChar[2]
		}

		result[key] = strconv.Itoa(int(math.Abs(float64(sumChar[0] - sumChar[1]))))
	}

	return result
}

func GetSumChar(data []string, key ...string) []int {
	sum := make([]int, len(key))

	for _, val := range data {
		for k, v := range key {
			if val == v {
				sum[k]++
			}
		}
	}

	return sum
}
