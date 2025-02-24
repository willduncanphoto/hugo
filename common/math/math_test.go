// Copyright 2018 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package math

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestDoArithmetic(t *testing.T) {
	t.Parallel()
	c := qt.New(t)

	for _, test := range []struct {
		a      any
		b      any
		op     rune
		expect any
	}{
		{3, 2, '+', int64(5)},
		{0, 0, '+', int64(0)},
		{3, 2, '-', int64(1)},
		{3, 2, '*', int64(6)},
		{3, 2, '/', int64(1)},
		{3.0, 2, '+', float64(5)},
		{0.0, 0, '+', float64(0.0)},
		{3.0, 2, '-', float64(1)},
		{3.0, 2, '*', float64(6)},
		{3.0, 2, '/', float64(1.5)},
		{3, 2.0, '+', float64(5)},
		{3, 2.0, '-', float64(1)},
		{3, 2.0, '*', float64(6)},
		{3, 2.0, '/', float64(1.5)},
		{3.0, 2.0, '+', float64(5)},
		{0.0, 0.0, '+', float64(0.0)},
		{3.0, 2.0, '-', float64(1)},
		{3.0, 2.0, '*', float64(6)},
		{3.0, 2.0, '/', float64(1.5)},
		{uint(3), uint(2), '+', uint64(5)},
		{uint(0), uint(0), '+', uint64(0)},
		{uint(3), uint(2), '-', uint64(1)},
		{uint(3), uint(2), '*', uint64(6)},
		{uint(3), uint(2), '/', uint64(1)},
		{uint(3), 2, '+', uint64(5)},
		{uint(0), 0, '+', uint64(0)},
		{uint(3), 2, '-', uint64(1)},
		{uint(3), 2, '*', uint64(6)},
		{uint(3), 2, '/', uint64(1)},
		{3, uint(2), '+', uint64(5)},
		{0, uint(0), '+', uint64(0)},
		{3, uint(2), '-', uint64(1)},
		{3, uint(2), '*', uint64(6)},
		{3, uint(2), '/', uint64(1)},
		{uint(3), -2, '+', int64(1)},
		{uint(3), -2, '-', int64(5)},
		{uint(3), -2, '*', int64(-6)},
		{uint(3), -2, '/', int64(-1)},
		{-3, uint(2), '+', int64(-1)},
		{-3, uint(2), '-', int64(-5)},
		{-3, uint(2), '*', int64(-6)},
		{-3, uint(2), '/', int64(-1)},
		{uint(3), 2.0, '+', float64(5)},
		{uint(0), 0.0, '+', float64(0)},
		{uint(3), 2.0, '-', float64(1)},
		{uint(3), 2.0, '*', float64(6)},
		{uint(3), 2.0, '/', float64(1.5)},
		{3.0, uint(2), '+', float64(5)},
		{0.0, uint(0), '+', float64(0)},
		{3.0, uint(2), '-', float64(1)},
		{3.0, uint(2), '*', float64(6)},
		{3.0, uint(2), '/', float64(1.5)},
		{"foo", "bar", '+', "foobar"},
		{3, 0, '/', false},
		{3.0, 0, '/', false},
		{3, 0.0, '/', false},
		{uint(3), uint(0), '/', false},
		{3, uint(0), '/', false},
		{-3, uint(0), '/', false},
		{uint(3), 0, '/', false},
		{3.0, uint(0), '/', false},
		{uint(3), 0.0, '/', false},
		{3, "foo", '+', false},
		{3.0, "foo", '+', false},
		{uint(3), "foo", '+', false},
		{"foo", 3, '+', false},
		{"foo", "bar", '-', false},
		{3, 2, '%', false},
	} {
		result, err := DoArithmetic(test.a, test.b, test.op)

		if b, ok := test.expect.(bool); ok && !b {
			c.Assert(err, qt.Not(qt.IsNil))
			continue
		}

		c.Assert(err, qt.IsNil)
		c.Assert(test.expect, qt.Equals, result)
	}
}
