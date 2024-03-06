package worker

import (
	"fmt"
)

type workerFunc func(lpCmdData IMyCmd) bool

func workerCmd1(lpCmdData IMyCmd) bool {
	switch t := lpCmdData.(type) {
	case *WorkerCmd1Data:
		break
	default:
		fmt.Println("pkg worker.workerCmd1 - Unknown type: ", t)
		return false
	}

	pCmdData := lpCmdData.(*WorkerCmd1Data)

	fmt.Println("pkg worker.workerCmd1 - Cmd Name: ", pCmdData.cmdName)
	fmt.Println("pkg worker.workerCmd1 - Size of Cmd: ", pCmdData.GetCmdSize())

	return true
}

func workerCmd2(lpCmdData IMyCmd) bool {
	switch t := lpCmdData.(type) {
	case *WorkerCmd2Data:
		break
	default:
		fmt.Println("pkg worker.workerCmd2 - Unknown type: ", t)
		return false
	}

	pCmdData := lpCmdData.(*WorkerCmd2Data)

	fmt.Println("pkg worker.workerCmd2 - Cmd Name: ", pCmdData.cmdName)
	fmt.Println("pkg worker.workerCmd2 - Size of Cmd: ", pCmdData.GetCmdSize())

	return true
}

func workerCmd3(lpCmdData IMyCmd) bool {
	switch t := lpCmdData.(type) {
	case *WorkerCmd3Data:
		break
	default:
		fmt.Println("pkg worker.workerCmd3 - Unknown type: ", t)
		return false
	}

	pCmdData := lpCmdData.(*WorkerCmd3Data)

	fmt.Println("pkg worker.workerCmd3 - Cmd Name: ", pCmdData.cmdName)
	fmt.Println("pkg worker.workerCmd3 - Size of Cmd: ", pCmdData.GetCmdSize())

	return true
}

var mapFunc = map[int]workerFunc{
	WORKER_CMD_CMD1: workerCmd1,
	WORKER_CMD_CMD2: workerCmd2,
	WORKER_CMD_CMD3: workerCmd3,
}

func lookupCmd(lpCmdData IMyCmd) workerFunc {
	var cmdID int

	// Do type assert
	switch t := lpCmdData.(type) {
	case *WorkerCmd1Data:
		cmdID = lpCmdData.(*WorkerCmd1Data).GetCmdID()
	case *WorkerCmd2Data:
		cmdID = lpCmdData.(*WorkerCmd2Data).GetCmdID()
	case *WorkerCmd3Data:
		cmdID = lpCmdData.(*WorkerCmd3Data).GetCmdID()
	default:
		fmt.Println("pkg worker.lookupCmd - Unknown type: ", t)
		return nil
	}

	if cmdID >= WORKER_CMD_COUNT {
		return nil
	}

	return mapFunc[cmdID]
}
