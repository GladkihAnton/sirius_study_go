Реализовать свой аналог asyncio.gather

```go
func gather(funcs []func() any) []any {
    // выполните все переданные функции,
    // соберите результаты в срез
    // и верните его
}

func squared(n int) func() any {
    return func() any {
        time.Sleep(time.Duration(n) * 100 * time.Millisecond)
        return n * n
    }
}

func main() {
    funcs := []func() any{squared(2), squared(3), squared(4)}

    start := time.Now()
    nums := gather(funcs)
    elapsed := float64(time.Since(start)) / 1_000_000

    fmt.Println(nums) // тут должно вернуть 4, 9, 16. Т.е. порядок важен
    fmt.Printf("Took %.0f ms\n", elapsed)
}
```

T: + (только goleak или объяснить почему их точно не будет)
B: -