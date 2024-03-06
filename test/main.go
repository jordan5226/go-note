package main

import (
	"jyutil"
	"log"
	"net/http"
	"worker"
)

func useCList() {
	jyutil.Output("======= CList =======")
	//listTmp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Slice store type: [3]int
	listTmp := [][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11}}
	var listALL jyutil.CList[[3]int]
	listALL.Set(listTmp)

	jyutil.Output("listALL: ", listALL.GetData())
	listALL.Del(2)
	jyutil.Output("listALL: ", listALL.GetData())
	listALL.DelValue([3]int{3, 4, 5})
	jyutil.Output("listALL: ", listALL.GetData())
}

func Math(x int, y int, callback func(int, int) int) int {
	return callback(x, y)
}

func useCallback() {
	ans1 := Math(2, 3, func(x int, y int) int {
		return (x + y)
	})
	jyutil.Output("Math Callback 1: ", ans1)

	ans2 := Math(2, 3, func(x int, y int) int {
		return (x * y)
	})
	jyutil.Output("Math Callback 2: ", ans2)
}

func useDefer() {
	defer jyutil.Output("ENDDO: useDefer")
	jyutil.Output("Into useDefer")
}

func errorHandle() {
	resp, err := http.Get("https://haha.net/")

	if err != nil {
		log.Fatalln(err)
		// 上面這一行等價於以下兩行
		//jyutil.Output(err)
		//os.Exit(1)
	}

	jyutil.Output(resp)
}

func main() {
	// Declare
	useDeclare()

	// Function
	useFunction()

	// if
	useIf()

	// map, for
	useMapFor()

	// Slice
	useSlice()

	// switch
	useSwitch()

	// CList
	useCList()

	// test function return function and API encapsulation
	data1 := worker.NewWorkerCmd1Data()
	worker.Worker_Execute(&data1)

	data2 := worker.NewWorkerCmd2Data()
	worker.Worker_Execute(&data2)

	data3 := worker.NewWorkerCmd3Data()
	worker.Worker_Execute(&data3)

	// callback
	useCallback()

	// defer
	useDefer()

	// invoke func itself
	func() {
		jyutil.Output("Invoke func itself.")
	}()

	// Goroutines
	testGoRoutines()

	// error handle
	errorHandle()

	// context
	testContext()
}
