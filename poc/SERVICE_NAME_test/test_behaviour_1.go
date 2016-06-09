package main

import (
	"log"

	common "github.com/alexgarzao/ms-gen/poc/SERVICE_NAME_common"
)

type TestBehaviour1 struct {
}

func NewTestBehaviour1() *TestBehaviour1 {
	return &TestBehaviour1{}
}

func (t *TestBehaviour1) RunAllTests() {
	log.Println("Checking behaviour 1")

	t.test1() // Remember to use great names :-)
	t.test2()
	//	...
	//	t.testN()
}

func (t *TestBehaviour1) test1() {
	// Action 1
	//     For example, send a valid request to method 1, method 2, ...
	//     Sure, you can send invalid requests to test desired behaviours...
	// Action 2
	// ...
	// Action N

	request := common.ServiceRequest1{
		FieldName1: "xxx",
		FieldName2: "yyy",
	}

	SendTestValidRequest1("Request XXX with valid infos", "parameter_value", request)
}

func (T *TestBehaviour1) test2() {
	request := common.ServiceRequest1{
		FieldName1: "invalid",
		FieldName2: "yyy",
	}

	SendTestInvalidRequest1("Request XXX with invalid infos", "parameter_value", request, "Expected message error")
}
