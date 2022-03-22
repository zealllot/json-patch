package jsonpatch

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrintJson(data []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out.String())
}
