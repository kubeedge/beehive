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

// Package beehive is a messaging framework based on go-channels for communication between modules of KubeEdge. A module registered with beehive can communicate with other beehive modules if the name with which other beehive module is registered or the name of the group of the module is known. Beehive supports following module operations:
//
// Add Module
// Add Module to a group
// CleanUp (remove a module from beehive core and all groups)
//
// Beehive supports following message operations:
//
// Send to a module/group
// Receive by a module
// Send Sync to a module/group
// Send Response to a sync message
package beehive

//go:generate  protoc  -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf   --gogo_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:./ ./pkg/core/model/message.proto
