package training

import (
	"fmt"
	"time"
)

func training() {
	fmt.Println("==== Starting Go Training! ====")
	// Conditionals()
	// Loops()
	// ArrayTypeStructures()
	// ObjectTypeStructures()
	// Pointers()
	// ErrorHandling()
	// GoRoutings()
	// GoChannels()
	// GoBufferChannels()
	// SelectStatements()
	// Interfaces()
	Embedding()
}

// Conditionals is a function that demos use of go conditionals
func Conditionals() {
	i1, i2 := 5, 10
	i3, i4 := 12, 6

	fmt.Println("==== If/Else Conditions ====")

	if i1 > i2 {
		fmt.Println(fmt.Sprintf("false: %v is greater than %v", i1, i2))
	} else if i2 > i1 {
		fmt.Println(fmt.Sprintf("false: %v is greater than %v", i2, i1))
	}

	fmt.Println(fmt.Sprintf("true: %v and %v are equal", i2, i1))

	fmt.Println("==== Switch Statement ====")

	switch {
	case i3 > i4:
		fmt.Println(fmt.Sprintf("false: %v is greater than %v", i3, i4))

	case i4 > i3:
		fmt.Println(fmt.Sprintf("false: %v is greater than %v", i4, i3))
	}

	fmt.Println(fmt.Sprintf("true: %v and %v are equal", i4, i3))
}

// Loops will demo looping in go
func Loops() {
	fmt.Println("==== For loops ====")

	target := 10
	for i := 0; i < target+1; i++ {
		fmt.Println(i)
	}
}

// ArrayTypeStructures will demo go array/slice data structures
func ArrayTypeStructures() {
	fmt.Println("==== Arrays ====")

	// var array [5]int = [5]int{1, 2, 3, 4, 5}
	array2 := [5]int{1, 2, 3, 4, 5} // array with type inference (note: size is fixed)
	fmt.Println("Fixed array:", array2)

	fmt.Println("==== Slices ====")

	// var slice []int = []int{1, 2, 3}
	slice2 := []int{1, 2, 3} // slice with type inference (note: size in dynamic)
	fmt.Println("Initial slice:", slice2)
	slice2 = append(slice2, 4, 5, 6)
	fmt.Println("Appended slice:", slice2)

	slice3 := array2[2:4] // alternate declaration of slice composed of taking part of array
	fmt.Println("Slice from array:", slice3)
}

// ObjectTypeStructures will demo maps/struct data structures
func ObjectTypeStructures() {
	fmt.Println("==== Maps ====")

	// var myMap map[string]int = make(map[string]int)
	myMap2 := make(map[string]int) // map with type inference
	myMap2["first"] = 1
	myMap2["second"] = 2
	fmt.Println(myMap2["first"]) // should print out first index of myMap2
	_, ok := myMap2["third"]     // search map key returns two values (value at index and bool [true/false]).
	fmt.Println("Value found at myMap2[third]:", ok)
	delete(myMap2, "first") // delete key from map

	fmt.Println("==== Structs ====")

	type person struct {
		Name    string
		Age     int
		Address string
	}

	jason := person{
		Name:    "Jason D.",
		Age:     30,
		Address: "Holland",
	}
	fmt.Println(jason)
}

// Pointers will demo using pointers in go
func Pointers() {
	var pI *int // memory address of a value of type int
	i := 3
	pI = &i                               // pI points to location of variable i
	println("Value of pI (address):", pI) // print memory address of pI
	println("Value of i:", i)             // print value of i
	increment(pI)
	println("Value of i (after increment):", i) // print value of i after increment
}

func increment(pI *int) {
	*pI++ // dereferencing
}

// ErrorHandling will demo using defers/panics/etc.
func ErrorHandling() {
	testDefers()
	testPanics()
}

func testDefers() {
	fmt.Println("==== Defers ====")

	// defer will store the code following the keyword for execution after the entire func (ErrorHandling has been completed)
	// defers are executed in a LIFO structure (note: should print 'Hello Super World')
	defer fmt.Println("World")
	defer fmt.Println("Super")
	fmt.Println("Hello")
}

func testPanics() {
	fmt.Println("==== Panics ====")

	// panics are used to throw exections (ex: throw error() in nodejs)
	// panic("A panic occurred")
	// fmt.Println("This will not be printed due to panic")

	// You need an anonymous function defered to handler recovery from panics (sorta like a catch())
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("We recovered from a panic!!")
		}
	}()
	panic("A panic occurred")
}

// GoRoutings will demo executing go code concurrently
func GoRoutings() {
	fmt.Println("==== Routines ====")

	// only the second Println will run due to program exiting before the routine has a chance to run
	// you need to utilize channels to effectively use goroutines
	go fmt.Println("Hello from a goroutine")
	fmt.Println("Hello from outside goroutine")
}

// GoChannels will demo executing go-routines via channels
func GoChannels() {
	fmt.Println("==== Channels ====")

	quitSignal := make(chan bool) // creates new channel of type bool

	go sayHelloFromGoRoutine(quitSignal)
	fmt.Println("Hello from outside new goroutine")
	v := <-quitSignal // will cause function to wait incoming value from goroutine
	fmt.Println("Value form goroutine:", v)

	ic := make(chan int)
	go periodicSend(ic)
	for i := range ic {
		fmt.Println(i)
	}
	// ic <- 3 // sending value to channel that has been closed (close()) will cause a panic
	_, ok := <-ic // ok will be true if channel is still open and false otherwise
	fmt.Println("Channel is open:", ok)
}

func sayHelloFromGoRoutine(qs chan bool) {
	// func takes a channel of type bool as argument
	fmt.Println("Hello from a new goroutine")
	qs <- true // sends a value into the channel
}

func periodicSend(ic chan int) {
	i := 0
	for i <= 10 {
		ic <- i
		i++
		time.Sleep(1 * time.Second)
	}
	close(ic) // will invalidate channel passed in
}

// GoBufferChannels will demo the use of buffer channels (channels that keeps buffer of stored values)
func GoBufferChannels() {
	fmt.Println("==== Buffer Channels ====")

	buffch := make(chan int, 5) // similar to channel declaration with exception of a size parameter (note: it's 5 in this case)
	buffch <- 3
	buffch <- 2
	fmt.Println(<-buffch) // buffers are a FIFO structure
	fmt.Println(<-buffch)
	fmt.Println(<-buffch) // channel buffers lock when they're either full or empty
}

// SelectStatements will demo waiting for data from multiple channels
func SelectStatements() {
	fmt.Println("==== Select Statements ====")

	ic := make(chan int)
	select {
	case v1 := <-waitAndSend(3, 1):
		fmt.Println(v1)
	case v2 := <-waitAndSend(5, 2):
		fmt.Println(v2)
	case ic <- 23:
		fmt.Println("ic received a value")
	default:
		fmt.Println("all channels are slow")
	}
}

func waitAndSend(v, i int) chan int {
	// this type of function signature is a channel generator
	// this function will wait i secs before sending value on return channel
	time.Sleep(time.Duration(i) * time.Second)
	retCh := make(chan int)
	go func() {
		time.Sleep(time.Duration(i) * time.Second)
		retCh <- v
	}()
	return retCh
}

// Interfaces will demo usage of interfaces in go
type testInterface interface {
	SayHello()
	Say(s string)
	Increment()
	GetInternalValue() int
}

type testStruct struct {
	i int
}

// Interfaces will demo usage of interfaces in Go
func Interfaces() {
	fmt.Println("==== Interfaces ====")

	// interfaces are a type in Go
	var calc testInterface
	calc = &testStruct{} // you can also do calc = new(testStruct) or newCalc() (which is a constructor method)
	calc.Say("I'm from Interfaces")
	calc.SayHello()
	calc.Increment()
	fmt.Println("General Value:", calc.GetInternalValue())
}

func newCalc() testInterface {
	// this is a constructor, it must return an interface type
	return &testStruct{}
}

func (tst *testStruct) SayHello() {
	fmt.Println("Hello")
}

func (tst *testStruct) Say(s string) {
	fmt.Println(s)
}

func (tst *testStruct) Increment() {
	tst.i++
}

func (tst *testStruct) GetInternalValue() int {
	return tst.i
}

// Embedding will demo encapsulating objects within other objects
func Embedding() {
	fmt.Println("==== Embedding ====")

	type testEmbed struct { // we want this struct to have all features available to testStruct (note: this is know as the outer type)
		*testStruct // the embed (note: this is known as the inner type)
	}

	te := testEmbed{testStruct: &testStruct{i: 50}}
	te.Say("I'm from Embedding")
	te.SayHello()
}
