package main

import (
	"bytes"
	"fmt"

	"github.com/s-vvardenfell/Adipiscing/redis_cache"
)

var p = fmt.Println

func main() {
	rs := redis_cache.New("localhost", 0, 100)
	rs.Set("123", bytes.NewReader([]byte("11111")))
	p(string(rs.Get("123")))
}
