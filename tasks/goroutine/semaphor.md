```go 
// начало решения

// Semaphore представляет семафор синхронизации.
type Semaphore

// NewSemaphore создает новый семафор указанной вместимости.
func NewSemaphore(n int) *Semaphore {
    // ...
}

// Acquire занимает место в семафоре, если есть свободное.
// В противном случае блокирует вызывающую горутину.
func (s *Semaphore) Acquire() {
    // ...
}

// Release освобождает место в семафоре и разблокирует
// одну из заблокированных горутин (если такие были).
func (s *Semaphore) Release() {
    // ...
}

// конец решения

func main() {
    const maxConc = 4
    sema := NewSemaphore(maxConc)
    start := time.Now()

	const nCalls = 12
	var wg sync.WaitGroup
	wg.Add(nCalls)

	for i := 0; i < nCalls; i++ {
		sema.Acquire()
		go func() {
			defer wg.Done()
			defer sema.Release()
			time.Sleep(10 * time.Millisecond)
			fmt.Print(".")
		}()
	}

	wg.Wait()

	fmt.Printf("\n%d calls took %d ms\n", nCalls, time.Since(start).Milliseconds())
	/*
		............
		12 calls took 30 ms
	*/
}
```
