/*
Copyright 2019 The KubeEdge Authors.

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

// Package typeurl is
package typeurl

import (
	"testing"
)

func TestMarshalAny(t *testing.T) {
	type Node struct {
		ID string
	}

	RegisterType(Node{}, "node")
	RegisterType(map[string]Node{}, "node.map")

	t.Run("IS ", func(t *testing.T) {
		flag := false
		any, err := MarshalAny(flag)
		if err != nil {
			t.Error(err)
		}
		var flag1 bool
		if !IS(any, flag1) {
			t.Errorf("need true")
		}
	})
	t.Run("string", func(t *testing.T) {
		str := "test"

		any, err := MarshalAny(str)
		if err != nil {
			t.Error(err)
		}

		v, err := UnmarshalAny(any)
		if err != nil {
			t.Error(err)
		}
		t.Logf("string:%v", v.(*string))
	})
	t.Run("*string", func(t *testing.T) {
		str := "test"

		any, err := MarshalAny(&str)
		if err != nil {
			t.Error(err)
		}

		v, err := UnmarshalAny(any)
		if err != nil {
			t.Error(err)
		}
		t.Logf("string:%v", v.(*string))
	})

	t.Run("struct", func(t *testing.T) {

		node := Node{ID: "test_id"}

		any, err := MarshalAny(node)
		if err != nil {
			t.Error(err)
		}

		v, err := UnmarshalAny(any)
		if err != nil {
			t.Error(err)
		}
		t.Logf("node:%v", v.(*Node))
	})

	t.Run("struct point", func(t *testing.T) {
		node := Node{ID: "test_id"}
		any, err := MarshalAny(&node)
		if err != nil {
			t.Error(err)
		}

		v, err := UnmarshalAny(any)
		if err != nil {
			t.Error(err)
		}
		t.Logf("node:%v", v.(*Node))
	})

	t.Run("node_map", func(t *testing.T) {
		node := Node{ID: "test_id"}
		cache := make(map[string]Node)
		cache[node.ID] = node
		any, err := MarshalAny(&cache)
		if err != nil {
			t.Error(err)
		}

		v, err := UnmarshalAny(any)
		if err != nil {
			t.Error(err)
		}
		t.Logf("node map:%v", v.(*map[string]Node))
	})
}
