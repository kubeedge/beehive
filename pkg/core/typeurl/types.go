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

// Package typeurl is Marushal For any type
package typeurl

import (
	"encoding/json"
	"path"
	"reflect"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
)

var registry map[reflect.Type]string
var revRegistry map[string]reflect.Type
var lock sync.Mutex

var (
	// WrongTypeError stand for custom type error
	WrongTypeError = errors.Errorf("Wrong Type Error")
)

func init() {
	registry = make(map[reflect.Type]string)
	revRegistry = make(map[string]reflect.Type)
	lock = sync.Mutex{}
	// Register sting
	RegisterType("", "string.content")
	// Register boolean
	RegisterType(true, "bool.content")
	// register slice
	RegisterType([]interface{}{}, "slice.interface.content")
}

func tryStrip(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Ptr:
		// if kind is Prt, get it`s realy type
		for t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		panic(WrongTypeError)
	}
	return t
}

// RegisterType is Register a type with the base url of the type
func RegisterType(v interface{}, args ...string) {
	t := tryStrip(v)
	param := []string{"kubeedge.io"}
	param = append(param, args...)
	url := path.Join(param...)

	lock.Lock()
	defer lock.Unlock()
	if u, ok := registry[t]; ok {
		panic(errors.Errorf("type %v has registerd , it's path is %v", t.String(), u))
	}
	if t, ok := revRegistry[url]; ok {
		panic(errors.Errorf("path %v has registerd, it`s type is %v", url, t.String()))
	}
	registry[t] = url
	revRegistry[url] = t
}

// TypeURL is get the typeurl string from interface
func TypeURL(v interface{}) (string, error) {
	lock.Lock()
	url, ok := registry[tryStrip(v)]
	lock.Unlock()
	if ok {
		return url, nil
	}
	panic(errors.Errorf("Can not get typeurl, Type %v not register", tryStrip(v)))
}

// IS returns true if the type of the Any is the same as v
func IS(any *types.Any, v interface{}) bool {
	url, err := TypeURL(v)
	if err != nil {
		return false
	}
	return any.GetTypeUrl() == url
}

// MarshalAny takes interface and encodes it into google.protobuf.Any.
func MarshalAny(v interface{}) (*types.Any, error) {
	var marshal func(v interface{}) ([]byte, error)
	switch t := v.(type) {
	case *types.Any:
		return t, nil
	case proto.Message:
		// do not support Proto.Message, if v is Proto.Message, please use types.UnmarshalAny()
		panic(WrongTypeError)
	default:
		marshal = json.Marshal
	}
	url, err := TypeURL(v)
	if err != nil {
		return nil, err
	}
	data, err := marshal(v)
	if err != nil {
		return nil, err
	}
	return &types.Any{
		TypeUrl: url,
		Value:   data,
	}, nil
}

// UnmarshalAny parses the interface representation in a google.protobuf.Any
// message and  return interface, It returns an error if type of
// contents of Any message does not match type of interface.
// the kind of interface will be pointer
func UnmarshalAny(any *types.Any) (interface{}, error) {
	t, ok := revRegistry[any.TypeUrl]
	if !ok {
		return nil, errors.Errorf("Any %v url not found", *any)
	}
	v := reflect.New(t).Interface()
	err := json.Unmarshal(any.Value, &v)
	return v, err
}
