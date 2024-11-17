package benchmarking

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func createRandomArray2(size int) []string {
	res := make([]string, 0)
	for range size {
		randI := rand.Intn(size + 1)
		randStr := strconv.Itoa(randI)
		res = append(res, randStr)
	}
	return res
}

func concatanateWithPlus(strs []string) string {
	var res string
	for _, str := range strs {
		res += str
	}
	return res
}
func concatanateWithJoin(strs []string) string {
	res := strings.Join(strs, "")
	return res
}

func RunTask6() {
	arr := createRandomArray2(10)
	fmt.Println(concatanateWithJoin(arr))
	fmt.Println(concatanateWithPlus(arr))
}
