Популярный способ защитить сервер от перегрузки запросами — ограничитель скорости (rate limiter):

Настраиваем, скажем, что сервер готов обрабатывать 10 запросов в секунду.
Теперь сколько бы запросов не прислал клиент, за секунду обработаются только 10 штук. Остальным придется ждать.
Сделайте функцию-обертку withRateLimit(limit, fn), которая следит, чтобы функция fn выполнялась не более limit раз в секунду:

```go
// withRateLimit следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
func withRateLimit(limit int, fn func()) (handle func() error, cancel func()) {
    // ...
}

func main() {
    work := func() {
        fmt.Print(".")
    }

    handle, cancel := withRateLimit(5, work)
    defer cancel()

    start := time.Now()
    const n = 10
    for i := 0; i < n; i++ {
        handle()
    }
    fmt.Println()
    fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
```

Если подряд сделаны несколько вызовов handle(), ограничитель должен выполнять fn не одномоментно, а равномерно:

- лимит 2/сек — интервал между запусками fn() 500 мс;
- имит 5/сек — интервал 200 мс;
- лимит 10/сек — интервал 100 мс;

и так далее.
Ограничитель должен учитывать только количество вызовов, но не время работы fn. Например, установлен лимит 10/сек, а handle() вызывается постоянно. Тогда:

- если fn занимает 1 мс, то каждую секунду должны запускаться 10 fn();
- если fn занимает 5 сек, то каждую секунду должны запускаться все те же 10 fn().

Обратите внимание на нюансы:

- и handle, и cancel не будут вызываться одновременно из разных горутин (то есть поддерживать одновременные вызовы не требуется).
- fn выполняется синхронно (в той же горутине, что вызвала handle).
- handle() может быть вызвана после того, как выполнилась cancel(). Тогда она должна вернуть ошибку ErrCanceled.
- cancel() должна освободить ресурсы, занятые ограничителем.
- cancel() может быть вызвана несколько раз. Тогда в первый раз она должна остановить ограничитель, а в последующие — ничего не делать.

Тесты:
- T: + (достаточно только goleak)
- B: -