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


type getRawIntWithBitSizeFixture struct {
	input string
	bitSize int
	result int64
}

var getRawIntWithBitSizeFixtures = []getRawIntWithBitSizeFixture{
	{"1", 0, 1},
	{"4.2",0,  0},
	{"150.64", 0, 0},
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

func (s *ClinputSuite) TestGetRawIntWithBitSize(c *C) {
	for _, fixture := range getRawIntWithBitSizeFixtures {
		input := getRawIntWithBitSize("", strings.NewReader(fixture.input), fixture.bitSize)
		c.Assert(input, Equals, fixture.result)
	}
}

func (s *ClinputSuite) TestGetIntInput(c *C) {
	input := GetIntInput("", strings.NewReader("1"))
	c.Assert(input, Equals, int(1))
	c.Assert(input, FitsTypeOf, int(1))
}

func (s *ClinputSuite) TestGetInt8Input(c *C) {
	input := GetInt8Input("", strings.NewReader("1"))
	c.Assert(input, Equals, int8(1))
	c.Assert(input, FitsTypeOf, int8(1))
}

func (s *ClinputSuite) TestGetInt16Input(c *C) {
	input := GetInt16Input("", strings.NewReader("1"))
	c.Assert(input, Equals, int16(1))
	c.Assert(input, FitsTypeOf, int16(1))
}

func (s *ClinputSuite) TestGetInt32Input(c *C) {
	input := GetInt32Input("", strings.NewReader("1"))
	c.Assert(input, Equals, int32(1))
	c.Assert(input, FitsTypeOf, int32(1))
}

func (s *ClinputSuite) TestGetInt64Input(c *C) {
	input := GetInt64Input("", strings.NewReader("1"))
	c.Assert(input, Equals, int64(1))
	c.Assert(input, FitsTypeOf, int64(1))
}
