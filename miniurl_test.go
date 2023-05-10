package miniurl_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuliyamun/miniurl"
)

func TestHashLength(t *testing.T) {
	const (
		input          = "https://github.com/yuliyamun/miniurl"
		expectedLength = 32
	)
	output := miniurl.Hash(input)
	assert.Len(t, output, expectedLength)
}

func TestHashIsDeterministic(t *testing.T) {
	const input = "https://github.com/yuliyamun/miniurl"
	output1 := miniurl.Hash(input)
	output2 := miniurl.Hash(input)
	assert.Equal(t, output1, output2)
}

func ExampleHash() {
	const input = "https://github.com/yuliyamun/miniurl"
	output := miniurl.Hash(input)
	fmt.Println(output)
}

func BenchmarkHash(b *testing.B) {
	const input = "https://github.com/yuliyamun/miniurl"
	for n := 0; n < b.N; n++ {
		miniurl.Hash(input)
	}

}
