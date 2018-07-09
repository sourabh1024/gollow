package main

import (
	"sync"
	"time"
)

type surge struct {
	rwMutex    sync.RWMutex
	surgeCache map[int]map[int64]float64
}

type geoHashSurge struct {
	geoHashSurge map[int64]float64
}

type mySurge struct {
	vehicleSurge map[int]*geoHashSurge
}

func (s *surge) read(i int) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	_ = s.surgeCache[i]
	return
}

func (s *surge) set(i int) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.surgeCache[i] = map[int64]float64{}
}

func main() {

	su := &surge{
		rwMutex:    sync.RWMutex{},
		surgeCache: map[int]map[int64]float64{},
	}

	n := 2000000
	start := time.Now()
	var wg sync.WaitGroup

	wg.Add(n)

	for i := 0; i < n/2; i++ {
		go func() {
			su.set(i)
			wg.Done()
		}()
	}

	for i := 0; i < n/2; i++ {
		go func() {
			su.read(i)
			wg.Done()
		}()
	}

	wg.Wait()

	println("Time :", time.Since(start))

	start = time.Now()

	var sm sync.Map

	wg.Add(n)

	x := map[int]float64{}
	for i := 0; i < n/2; i++ {
		go func() {
			sm.Store(i, x)
			wg.Done()
		}()
	}

	for i := 0; i < n/2; i++ {
		go func() {
			x, _ := sm.Load(i)
			if x != nil {

				y := x.(map[int]float64)
				if y == nil {

				}
			}
			wg.Done()
		}()
	}

	wg.Wait()

	println("Time :", time.Since(start))

}
