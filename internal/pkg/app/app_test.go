// Copyright 2020 Buf Technologies Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvContainer(t *testing.T) {
	envContainer := NewEnvContainer(
		map[string]string{
			"foo1": "bar1",
			"foo2": "bar2",
			"foo3": "",
		},
	)
	assert.Equal(t, "bar1", envContainer.Env("foo1"))
	assert.Equal(t, "bar2", envContainer.Env("foo2"))
	assert.Equal(t, "", envContainer.Env("foo3"))
	assert.Equal(
		t,
		[]string{
			"foo1=bar1",
			"foo2=bar2",
		},
		Environ(envContainer),
	)

	envContainer, err := newEnvContainerForEnviron(
		[]string{
			"foo1=bar1",
			"foo2=bar2",
			"foo3=",
		},
	)
	require.NoError(t, err)
	assert.Equal(t, "bar1", envContainer.Env("foo1"))
	assert.Equal(t, "bar2", envContainer.Env("foo2"))
	assert.Equal(t, "", envContainer.Env("foo3"))
	assert.Equal(
		t,
		[]string{
			"foo1=bar1",
			"foo2=bar2",
		},
		Environ(envContainer),
	)

	_, err = newEnvContainerForEnviron(
		[]string{
			"foo1=bar1",
			"foo2=bar2",
			"foo3",
		},
	)
	require.Error(t, err)
}

func TestArgContainer(t *testing.T) {
	args := []string{"foo", "bar", "baz"}
	assert.Equal(t, args, Args(NewArgContainer(args...)))
}
