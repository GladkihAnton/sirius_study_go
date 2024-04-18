Реализуйте тип Counter с помощью sync.RWMutex.

Интерфейс:
```go
// увеличивает значение по ключу на 1
Increment(str string)

// возвращает значение по ключу
Value(str string) int

// проходит по всем записям, и для каждой вызывает функцию fn,
// передавая в нее ключ и значение
Range(fn func(key string, val int))
```

Пример использования:
```go
func main() {
    counter := NewCounter()

    var wg sync.WaitGroup
    wg.Add(3)

    increment := func(key string, val int) {
        defer wg.Done()
        for ; val > 0; val-- {
            counter.Increment(key)
        }
    }

    go increment("one", 100)
    go increment("two", 200)
    go increment("three", 300)

    wg.Wait()

    fmt.Println("two:", counter.Value("two"))

    fmt.Print("{ ")
    counter.Range(func(key string, val int) {
        fmt.Printf("%s:%d ", key, val)
    })
    fmt.Println("}")
}
```