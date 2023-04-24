package encoder

import "C"
import (
	"fmt"
	"github.com/tliron/py4go"
)

func Run() {
	// Initialize Python
	python.Initialize()
	defer python.Finalize()

	// Import Python code (foo.py)
	encoder, _ := python.Import("encoder")
	defer encoder.Release()

	// Get access to a Python function
	hello, _ := foo.GetAttr("hello")
	defer hello.Release()

	// Call the function with arguments
	r, _ := hello.Call("myargument")
	defer r.Release()
	fmt.Printf("Returned: %s\n", r.String())

	// Expose a Go function to Python via a C wrapper
	// (Just use "import api" from Python)
	api, _ := python.CreateModule("api")
	defer api.Release()
	api.AddModuleCFunctionNoArgs("my_function", C.api_my_function)
	api.EnableModule()
}
