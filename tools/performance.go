package tools

import (
	"fmt"
	"runtime"
	"time"
)

const PERF_TIMES = 1000

type Performance struct {
	Time        time.Duration
	TotalMemory uint64
	Mallocs     uint64
	Params      PerformanceParam
}

type PerformanceParam map[string]uint64

func MeasurePerf(name string, solve func(int, *Performance)) {
	perf := new(Performance)
	perf.Params = make(map[string]uint64)

	for n := 0; n < PERF_TIMES; n++ {
		m1, start := preparePerf()

		solve(100000, perf)

		perf.addPerf(m1, start)
	}

	perf.print(name)
}

func (p *Performance) AddToParam(name string, value uint64) {
	if _, ok := p.Params[name]; ok {
		p.Params[name] += value

		return
	}

	p.Params[name] = value
}

func preparePerf() (*runtime.MemStats, time.Time) {
	m := new(runtime.MemStats)
	runtime.GC()
	runtime.ReadMemStats(m)
	start := time.Now()

	return m, start
}

func (p *Performance) addPerf(m1 *runtime.MemStats, start time.Time) {
	m2 := new(runtime.MemStats)

	runtime.GC()
	runtime.ReadMemStats(m2)

	p.Time += time.Since(start)
	p.TotalMemory += m2.TotalAlloc - m1.TotalAlloc
	p.Mallocs += m2.Mallocs - m1.Mallocs
}

func (p *Performance) print(name string) {
	fmt.Printf("========= %s =========\n", name)
	fmt.Println("time:", p.Time/PERF_TIMES)
	fmt.Println("total memory allocated:", p.TotalMemory/PERF_TIMES)
	fmt.Println("mallocs (cumulative count of heap objects allocated):", p.Mallocs/PERF_TIMES)
	for key, val := range p.Params {
		fmt.Println(fmt.Sprintf("%s:", key), val/PERF_TIMES)
	}
}

// func SinglePerf(name string, params ...*PerformanceParam) func() {
// 	var m1, m2 runtime.MemStats
// 	runtime.GC()
// 	runtime.ReadMemStats(&m1)

// 	start := time.Now()
// 	return func() {
// 		fmt.Printf("========= %s =========\n", name)
// 		fmt.Println("time:", time.Since(start))

// 		runtime.ReadMemStats(&m2)
// 		fmt.Println("total memory allocated:", m2.TotalAlloc-m1.TotalAlloc)
// 		fmt.Println("mallocs (cumulative count of heap objects allocated):", m2.Mallocs-m1.Mallocs)
// 		fmt.Println("stack in use", m1.StackInuse)

// 		for _, param := range params {
// 			fmt.Printf("%s: %v\n", param.Name, param.Value)
// 		}
// 	}
// }
