package homework07

import (
	"fmt"
	"sync"
	"time"
)

func SayHello() {
	fmt.Println("Hello world")
}

func ExcuteTest() {
	TestWaitGroup()
}

func TestGoroutine() {
	go SayHello()
	time.Sleep(1 * time.Second)
}

func TestChannel() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	value := <-ch
	fmt.Println(value)
}

// 互斥锁
var count int
var mutex sync.Mutex

func increment() {
	mutex.Lock()
	defer mutex.Unlock()

	count++
}

func TestMutex() {
	for i := 0; i < 10; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(count)
}

// 读写锁
var count2 int
var rwMutex sync.RWMutex

func read() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	fmt.Println("Count:", count)
}

func write() {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	count++
}

func TestRWMutex() {
	for i := 0; i < 10; i++ {
		go read()
	}

	for i := 0; i < 10; i++ {
		go write()
	}

	time.Sleep(1 * time.Second)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker", id, "started")
	time.Sleep(1 * time.Second)
	fmt.Println("Worker", id, "finished")
}

func TestWaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All workers finished")
}
