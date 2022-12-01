package structs

import (
	"fmt"
	"math/rand"
	"time"
)

type HashMap[T any] struct {
	values [256]*hashNode[T]
	len    int
}

type hashNode[T any] struct {
	key   string
	value T
}

var table [256]int = func() [256]int {
	var data [256]int
	rand.Seed(time.Now().Unix())
	for index := range data {
		data[index] = index
	}
	for index := range data {
		rv := rand.Intn(256)
		data[index], data[rv] = data[rv], data[index]
	}
	return data
}()

func hash8(message string) int {
	var hash int = len(message) % 256
	for _, value := range message {
		hash = table[(hash+int(value))%256]
	}
	return hash
}

func (self *HashMap[T]) Set(k string, v T) {
	if self.values[hash8(k)] == nil {
		self.values[hash8(k)] = &hashNode[T]{
			key:   k,
			value: v,
		}
		self.len++
	}
}

func (self *HashMap[T]) Get(k string) T {
	if self.values[hash8(k)] == nil {
		var zero T
		return zero
	}
	return self.values[hash8(k)].value
}

func (self *HashMap[T]) Dump() {
	for i := range self.values {
		if self.values[i] != nil {
			fmt.Println(i, self.values[i])
		}
	}
}
