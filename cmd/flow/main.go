package main

import (
	"encoding/json"
	"fmt"
	flow "github.com/s8sg/goflow/flow/v1"
	goflow "github.com/s8sg/goflow/v1"
	"math/rand"
	"strconv"
)

func Input(data []byte, option map[string][]string) ([]byte, error) {
	var input map[string]int
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	outputInt := input["input"]
	return []byte(strconv.Itoa(outputInt)), nil
}

func AddOne(data []byte, option map[string][]string) ([]byte, error) {
	num, _ := strconv.Atoi(string(data))
	outputInt := num + rand.Intn(10) + 1
	fmt.Println("AddOne=", outputInt)
	return []byte(strconv.Itoa(outputInt)), nil
}

func AddTwo(data []byte, option map[string][]string) ([]byte, error) {
	num, _ := strconv.Atoi(string(data))
	outputInt := num + rand.Intn(101) + 100
	fmt.Println("AddTwo=", outputInt)
	return []byte(strconv.Itoa(outputInt)), nil
}

func Output(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("Output=", string(data))
	return []byte("ok"), nil
}

func Myflow(workflow *flow.Workflow, context *flow.Context) error {
	dag := workflow.Dag()
	// 构建节点
	dag.Node("input", Input)
	dag.Node("add-one", AddOne)
	dag.Node("add-two", AddTwo)
	dag.Node("output", Output)

	// 构建依赖关系
	dag.Edge("input", "add-one")
	dag.Edge("add-one", "add-two")
	dag.Edge("add-two", "output")
	return nil
}

func main() {
	fs := &goflow.FlowService{
		Port:              8080,
		RedisURL:          "localhost:6379",
		WorkerConcurrency: 5,
	}
	err := fs.Register("add-flow", Myflow)
	if err != nil {
		return
	}
	if err = fs.Start(); err != nil {
		panic(err)
	}

}
