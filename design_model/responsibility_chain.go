package main

import (
	"errors"
	"fmt"
)

func main() {
	intercepts := make([]intercept, 0)
	intercepts = append(intercepts, processA)
	intercepts = append(intercepts, processB)
	intercepts = append(intercepts, processC)

	processIn := func(n int) error { return nil }

	chain := interceptChain(intercepts...)
	err := chain(2, processIn)
	fmt.Println(err)
}

type process func(n int) error
type intercept func(n int, p process) error

func processA(i int, p process) error {
	if i == 1 {
		return errors.New("1")
	}
	return p(i)
}

func processB(i int, p process) error {
	if i == 2 {
		return errors.New("2")
	}
	return p(i)
}

func processC(i int, p process) error {
	if i == 3 {
		return errors.New("3")
	}
	return p(i)
}

func interceptChain(intercepts ...intercept) intercept {
	build := func(n int, i intercept, p process) process {
		return func(n int) error {
			return i(n, p)
		}
	}
	return func(n int, p process) error {
		chain := p
		for i := len(intercepts) - 1; i >= 0; i-- {
			chain = build(n, intercepts[i], chain)
		}
		return chain(n)
	}
}
