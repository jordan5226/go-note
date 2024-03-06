package main

import (
	"fmt"
	"jyutil"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var a int = 10
var mxLock sync.Mutex

func foo() {
	for i := 1; i <= 10; i++ {
		fmt.Println("foo: ", i)
		time.Sleep(10 * time.Millisecond)
	}

	wg.Done()
}

func bar() {
	for i := 1; i <= 10; i++ {
		fmt.Println("bar: ", i)
		time.Sleep(10 * time.Millisecond)
	}

	wg.Add(-1)
}

func assignValue(value int) {
	defer wg.Done()

	mxLock.Lock()
	{
		result := a
		time.Sleep(300 * time.Millisecond)
		result -= value
		a = result
		fmt.Println("a: ", a)

		mxLock.Unlock()
	}
}

// ===============================================
// Routine1
// ===============================================
type Routine1 struct {
	end chan bool
	str chan string
}

// Constructor
func NewRoutine1() Routine1 {
	c := Routine1{}
	c.end = make(chan bool)
	c.str = make(chan string)
	return c
}

// Thread
func (c *Routine1) Thread() {
	var strTmp string = "test"

	for i := 0; i < 5; i++ {
		strTmp += strconv.Itoa(i)
		c.str <- strTmp             // Send string to channel:str
		time.Sleep(time.Second * 1) // Sleep 1 Second
	}

	c.end <- true // Send true to channel:end
}

// waitThreadEnd
func (c *Routine1) waitThreadEnd() {
	<-c.end // Wait channel:end until received value, then send end
	close(c.str)
}

func testGoRoutines() {
	jyutil.Output("======= Goroutine =======")

	//
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

	//
	wg.Add(2)
	go assignValue(5)
	go assignValue(10)
	wg.Wait()

	//
	r := NewRoutine1()

	go r.Thread()
	go r.waitThreadEnd()

	for v := range r.str { // continue waitting until r.str closed
		jyutil.Output("Channel: ", v)
	}
	close(r.end)
}
