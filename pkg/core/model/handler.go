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

// Package model is
package model

import (
	"errors"
	"reflect"
	"sync"

	"github.com/gogo/protobuf/proto"
)

type MessageResource interface {
	Type() string
	proto.Message
}

var lock sync.Mutex
var messageResourceMap map[string]struct{}

var (
	ResourceWrongTypeError    = errors.New("Resource Type Wrong")
	ResourceTypeNotExistError = errors.New("Resource Type Not Exist")
	ResourceTypeExistError    = errors.New("Resource Type Has Exist")
	ResourceNilError          = errors.New("Resource is nil")
	ExpectPointerError        = errors.New("Expect a pointer")
)

func init() {
	lock = sync.Mutex{}
	messageResourceMap = make(map[string]struct{})
}

func RegisterMessageResource(rs ...MessageResource) {
	lock.Lock()
	defer lock.Unlock()
	for _, r := range rs {
		if _, ok := messageResourceMap[r.Type()]; ok {
			panic(ResourceTypeExistError)
		}
		messageResourceMap[r.Type()] = struct{}{}
	}
}

func MarshalResource(r MessageResource) (*Resource, error) {
	data, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	return &Resource{
		ResourceType: r.Type(),
		Data:         data,
	}, nil
}

func UnmarshalResource(r *Resource, h MessageResource) error {

	if r == nil {
		return ResourceNilError
	}

	if _, ok := messageResourceMap[r.GetResourceType()]; !ok {
		return ResourceTypeNotExistError
	}

	if h.Type() != r.GetResourceType() {
		return ResourceWrongTypeError
	}

	if reflect.ValueOf(h).Kind() != reflect.Ptr {
		return ExpectPointerError
	}

	return proto.Unmarshal(r.GetData(), h)
}
