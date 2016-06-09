package main

import (
	"github.com/verdverm/frisby"
)

const (
	MYSERVICE_URL = "http://localhost:8090/" // TODO: put in config.yaml
)

func main() {
	frisby.Global.PrintProgressName = true

	NewTestBehaviour1().RunAllTests()
	//	NewTestBehaviour2().RunAllTests()
	//	...
	//	NewTestBehaviourN().RunAllTests()

	frisby.Global.PrintReport()
}
