package bm22

import (
	"fmt"
	"strconv"
	"strings"
)

func compare(version1 string, version2 string) int {
	// write code here
	var v1, v2 []int
	for _, s := range strings.Split(version1, ".") {
		n, _ := strconv.Atoi(s)
		v1 = append(v1, n)
	}
	for _, s := range strings.Split(version2, ".") {
		n, _ := strconv.Atoi(s)
		v2 = append(v2, n)
	}
	var i, j int
	for i = len(v1) - 1; i >= 0 && v1[i] == 0; i-- {
	}
	if i >= 0 {
		v1 = v1[0 : i+1]
	}
	for j = len(v2) - 1; j >= 0 && v2[j] == 0; j-- {
	}
	if j >= 0 {
		v2 = v2[0 : j+1]
	}
	return compareVersionList(v1, v2)
}

func compareVersionList(v1 []int, v2 []int) int {
	fmt.Printf("%v %v\n", v1, v2)
	bool2Int := func(a, b int) int {
		if a == b {
			return 0
		}
		if a < b {
			return -1
		}
		return 1
	}
	for i := 0; i < len(v1) && i < len(v2); i++ {
		if v1[i] == v2[i] {
			continue
		}
		return bool2Int(v1[i], v2[i])
	}

	return bool2Int(len(v1), len(v2))
}
