1) Design your implementation of the circular double-ended queue (deque).

Implement the MyCircularDeque class:

- MyCircularDeque(int k) Initializes the deque with a maximum size of k.
- boolean insertFront() Adds an item at the front of Deque. Returns true if the operation is successful, or false otherwise.
- boolean insertLast() Adds an item at the rear of Deque. Returns true if the operation is successful, or false otherwise.
- boolean deleteFront() Deletes an item from the front of Deque. Returns true if the operation is successful, or false otherwise.
- boolean deleteLast() Deletes an item from the rear of Deque. Returns true if the operation is successful, or false otherwise.
- int getFront() Returns the front item from the Deque. Returns -1 if the deque is empty.
- int getRear() Returns the last item from Deque. Returns -1 if the deque is empty.
- boolean isEmpty() Returns true if the deque is empty, or false otherwise.
- boolean isFull() Returns true if the deque is full, or false otherwise.


```go
type MyCircularDeque struct {

}


func Constructor(k int) MyCircularDeque {

}


func (this *MyCircularDeque) InsertFront(value int) bool {

}


func (this *MyCircularDeque) InsertLast(value int) bool {

}


func (this *MyCircularDeque) DeleteFront() bool {

}


func (this *MyCircularDeque) DeleteLast() bool {

}


func (this *MyCircularDeque) GetFront() int {

}


func (this *MyCircularDeque) GetRear() int {

}


func (this *MyCircularDeque) IsEmpty() bool {

}


func (this *MyCircularDeque) IsFull() bool {

}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

```

2) В качестве аргумента функция calculator должна получать два канала только для чтения, возвращает канал только для чтения.
   - Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу. Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел. 
   - Функция calculator должна быть неблокирующей, сразу возвращая управление 
   - Выходной канал должен быть закрыт после выполнения всех оговоренных условий, если вы этого не сделаете, то превысите предельное время работы.


3) Реализовать свой аналог asyncio.gather

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