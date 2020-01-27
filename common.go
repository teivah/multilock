package multilock

import (
	"fmt"
	"hash/fnv"
)

func addr(i interface{}) string {
	return fmt.Sprintf("%p", i)
}

func distribution(s string, length int) int {
	h := fnv.New32()
	_, _ = h.Write([]byte(s))
	return int(h.Sum32()) % length
}
