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
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func TestRootClinput(t *testing.T) { TestingT(t) }

type ClinputSuite struct {
	getInputReader *strings.Reader
}

const getStringInputInput string = "test\n"
const getStringInputOutput string = "test"

type getFloat64InputFixture struct {
	input string
	result float64
}

var getFloat64InputFixtures = []getFloat64InputFixture{
	{"0", 0},
	{"4.2", 4.2},
	{"150.64", 150.64},
}

var _ = Suite(&ClinputSuite{
	getInputReader: strings.NewReader(getStringInputInput),
})

func (s *ClinputSuite) TestClinput(c *C) {

}

func (s *ClinputSuite) TestGetStringInput(c *C) {
	input := GetStringInput("", s.getInputReader)
	c.Assert(input, Equals, getStringInputOutput)
}

func (s *ClinputSuite) TestGetFloat64Input(c *C) {
	for _, fixture := range getFloat64InputFixtures {
		input := GetFloat64Input("", strings.NewReader(fixture.input))
		c.Assert(input, Equals, fixture.result)
	}
}
