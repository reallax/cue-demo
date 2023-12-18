package main

import (
	"fmt"
	casbin "github.com/casbin/casbin/v2"
)

func main() {
	fmt.Println("Hello, World!")

	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		// handle err
		panic(err)
	}

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// handle err
		panic(err)
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println("permit alice to read data1")
	} else {
		// deny the request, show an error
		fmt.Println("deny the request, show an error")
	}
}
