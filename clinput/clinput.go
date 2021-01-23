/*

Copyright 2021 CJ Harries

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package clinput

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func GetStringInput(prompt string, source io.Reader) string {
	reader := bufio.NewReader(source)
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	return strings.Replace(input, "\n", "", -1)
}

func GetFloat64Input(prompt string, source io.Reader) float64 {
	stringResult := GetStringInput(prompt, source)
	float64result, _ := strconv.ParseFloat(stringResult, 64)
	return float64result
}
