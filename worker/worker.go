package worker

import (
	"unsafe"
)

const (
	WORKER_CMD_CMD1 = iota
	WORKER_CMD_CMD2
	WORKER_CMD_CMD3
	WORKER_CMD_COUNT
)

// To implement a Class like C++
// 1. Define a interface type, and define a method declaration as constructor in the interface.
// ==============================================
// MyCmd Interface
// ==============================================
type IMyCmd interface {
	InitMyCmd(cmdID int, cmdSize int)
}

// 2. Define a type as base class and at least implement the method of interface as constructor.
// ==============================================
// MyCmd
// ==============================================
const MYCMD_SIZE int = (int)(unsafe.Sizeof(MyCmd{}))

type MyCmd struct {
	cmdID   int
	cmdSize int
}

func (c *MyCmd) InitMyCmd(cmdID int, cmdSize int) {
	c.cmdID = cmdID
	c.cmdSize = cmdSize
}

func (c *MyCmd) GetCmdID() int {
	return c.cmdID
}

func (c *MyCmd) GetCmdSize() int {
	return c.cmdSize
}

// 3. Define a derived class type also implement the constructor method of interface.
// And declare a base class attribute in derived class.
// ==============================================
// WorkerCmd1Data
// ==============================================
const WORKERCMD1DATA_SIZE int = (int)(unsafe.Sizeof(WorkerCmd1Data{}))

type WorkerCmd1Data struct {
	MyCmd
	cmdName string
}

func (c *WorkerCmd1Data) InitMyCmd(cmdID int, cmdSize int) {
	c.MyCmd.InitMyCmd(cmdID, cmdSize)
}

func NewWorkerCmd1Data() WorkerCmd1Data {
	c := WorkerCmd1Data{}
	c.InitMyCmd(WORKER_CMD_CMD1, WORKERCMD1DATA_SIZE)
	c.cmdName = "WorkerCmd1Data"
	return c
}

// ==============================================
// WorkerCmd2Data
// ==============================================
const WORKERCMD2DATA_SIZE int = (int)(unsafe.Sizeof(WorkerCmd2Data{}))

type WorkerCmd2Data struct {
	MyCmd
	cmdName string
}

func (c *WorkerCmd2Data) InitMyCmd(cmdID int, cmdSize int) {
	c.MyCmd.InitMyCmd(cmdID, cmdSize)
}

func NewWorkerCmd2Data() WorkerCmd2Data {
	c := WorkerCmd2Data{}
	c.InitMyCmd(WORKER_CMD_CMD2, WORKERCMD2DATA_SIZE)
	c.cmdName = "WorkerCmd2Data"
	return c
}

// ==============================================
// WorkerCmd3Data
// ==============================================
const WORKERCMD3DATA_SIZE int = (int)(unsafe.Sizeof(WorkerCmd3Data{}))

type WorkerCmd3Data struct {
	MyCmd
	cmdName string
}

func (c *WorkerCmd3Data) InitMyCmd(cmdID int, cmdSize int) {
	c.MyCmd.InitMyCmd(cmdID, cmdSize)
}

func NewWorkerCmd3Data() WorkerCmd3Data {
	c := WorkerCmd3Data{}
	c.InitMyCmd(WORKER_CMD_CMD3, WORKERCMD3DATA_SIZE)
	c.cmdName = "WorkerCmd3Data"
	return c
}

// ==============================================
// Worker_Execute
// ==============================================
func Worker_Execute(lpCmdData IMyCmd) bool {
	function := lookupCmd(lpCmdData)

	if function == nil {
		return false
	}

	return function(lpCmdData)
}
