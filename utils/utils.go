package utils

import (
	"hash/fnv"
)

func HashID(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
