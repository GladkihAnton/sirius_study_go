Сделать так, чтобы вывод был first, second и third. Использовать mutex

```go
func MainRace() {
    wg := sync.WaitGroup{}
    wg.Add(3)
    
    go func() {
        fmt.Println("first")
        wg.Done()
    }()
    go func() {
        fmt.Println("second")
        wg.Done()
    }()
    go func() {
        fmt.Println("third")
        wg.Done()
    }()

    wg.Wait()
}
```
