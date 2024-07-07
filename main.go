package main

import (
	"github.com/JFKongphop/go-json-render/json-render"
)

func main() {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"name":    "John",
			"age":     30,
			"hobbies": []string{"reading", "swimming"},
		},
		"city": "New York",
	}

	jsonrender.Iterate(data, 0)
}
