package util

import (
	"math"
	"runtime"
	"sync"
)

type EnumCallback func(i int)

func getGoroutineCount() int {
	return int(math.Max(1, float64(runtime.NumCPU()-1)))
}

func ConcurrentEnum(start, end int, callback EnumCallback) {
	var wg sync.WaitGroup
	usingCore := getGoroutineCount()
	wg.Add(usingCore)
	xloop := func(start, end int) {
		defer wg.Done()
		for i := start; i < end; i++ {
			callback(i)
		}
	}

	work := (end - start) / usingCore
	for c := start; c < usingCore; c++ {
		if c == usingCore-1 { // Last
			xloop(start+c*work, end)
		} else {
			go xloop(start+c*work, (c+1)*work)
		}
	}
	wg.Wait()
}
