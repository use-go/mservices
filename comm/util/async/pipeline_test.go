package async

import (
	"context"
	"testing"
	"time"
)

func TestPipeLine_Wait(t *testing.T) {
	cwt, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	pipeline := NewPipeLine(cwt, 10)
	pipeline.Consumer(func(v int) {
		time.Sleep(1 * time.Second)
	})
	init := int(0)
	pipeline.ProductUntil(func(it *int) ([]int, bool) {
		*it = *it + 1
		return []int{*it}, *it == 100
	}, &init)
	pipeline.Wait()
}
