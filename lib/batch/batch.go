package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, 0, n)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for k := int64(0); k < n; k += pool {

		wg.Add(int(pool))
		for i := int64(0); i < pool; i++ {
			go func(wg *sync.WaitGroup, mu *sync.Mutex, i int64) {
				defer wg.Done()
				user := getOne(k + i)
				mu.Lock()
				res = append(res, user)
				mu.Unlock()
			}(wg, mu, i)
		}
		wg.Wait()
	}

	return res
}
