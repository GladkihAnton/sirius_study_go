Реализуйте конвейер, который делает следующее:

генерит слова из 5 букв;
отсеивает слова, в которых буквы повторяются;
переворачивает оставшиеся слова;
печатает пары «исходное-перевернутое» для первых n из них.
Пример. Если исходно сгенерились слова water, happy и lemon, то в итоге должно напечататься:
```
water -> retaw
lemon -> nomel
```
(!) Все этапы должны поддерживать отмену.

```go
func main() {
    cancel := make(chan struct{})
    defer close(cancel)

    c1 := generate(cancel)
    c2 := takeUnique(cancel, c1)
    c3_1 := reverse(cancel, c2)
    c3_2 := reverse(cancel, c2)
    c4 := merge(cancel, c3_1, c3_2)
    print(cancel, c4, 10)
}
```

```go
func generate(cancel <-chan) <-chan {
	out := make(chan)
	go func() {
		// ...
	}()
	return out
}

func takeUnique(cancel <-chan, in <-chan) <-chan  {
	out := make(chan)
	go func() {
		// ...
	}()
	return out
}

func reverse(cancel <-chan, in <-chan) <-chan {
	out := make(chan)
	go func() {
		// ...
	}()
	return out
}

func merge(cancel <-chan, c1, c2 <-chan) <-chan {
	out := make(chan)
	go func() {
		// ...
	}()
	return out
}

func print(cancel <-chan, in <-chan, n int) {
}


func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}
```

T: + (достаточно только goleak)
B: -