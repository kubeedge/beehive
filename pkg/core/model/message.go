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
	"time"

	"github.com/kubeedge/beehive/pkg/core/typeurl"
	"github.com/pkg/errors"

	"github.com/satori/go.uuid"
)

//Constants for database operations and resource type settings
const (
	InsertOperation        = "insert"
	DeleteOperation        = "delete"
	QueryOperation         = "query"
	UpdateOperation        = "update"
	ResponseOperation      = "response"
	ResponseErrorOperation = "error"
)

//BuildRouter sets route and resource operation in message
// res interface need MarshalAny type successfully
func (msg *Message) BuildRouter(source, group string, res interface{}, opr string) *Message {
	msg.SetRoute(source, group)
	msg.SetResourceOperation(res, opr)
	return msg
}

//SetResourceOperation sets router resource and operation in message
// res interface need MarshalAny type successfully
func (msg *Message) SetResourceOperation(res interface{}, opr string) *Message {
	real, err := typeurl.MarshalAny(res)
	if err != nil {
		panic(errors.Errorf("Marshal resource to any type error %v", err))
	}
	if msg.Router == nil {
		msg.Router = new(MessageRoute)
	}
	msg.Router.Resource = real
	msg.Router.Operation = opr
	return msg
}

//SetRoute sets router source and group in message
func (msg *Message) SetRoute(source, group string) *Message {
	if msg.Router == nil {
		msg.Router = new(MessageRoute)
	}
	msg.Router.Source = source
	msg.Router.Group = group
	return msg
}

// IsSync : msg.Header.Sync will be set in sendsync
func (msg *Message) IsSync() bool {
	return msg.Header.Sync
}

//GetResource returns message route resource
func (msg *Message) GetResource() interface{} {
	res, err := typeurl.UnmarshalAny(msg.Router.Resource)
	if err != nil {
		panic(errors.Errorf("Unmarshal any type resource to interface error %v", err))
	}
	return res
}

//GetOperation returns message route operation string
func (msg *Message) GetOperation() string {
	return msg.Router.Operation
}

//GetSource returns message route source string
func (msg *Message) GetSource() string {
	return msg.Router.Source
}

//GetGroup returns message route group
func (msg *Message) GetGroup() string {
	return msg.Router.Group
}

//GetID returns message ID
func (msg *Message) GetID() string {
	return msg.Header.ID
}

//GetParentID returns message parent id
func (msg *Message) GetParentID() string {
	return msg.Header.ParentID
}

//GetTimestamp returns message timestamp
func (msg *Message) GetTimestamp() int64 {
	return msg.Header.Timestamp
}

//GetContent returns message content
func (msg *Message) GetContent() interface{} {
	content, err := typeurl.UnmarshalAny(msg.Data)
	if err != nil {
		panic(errors.Errorf("Unmarshal Any type data to interface error %v", err))
	}
	return content
}

//UpdateID returns message object updating its ID
func (msg *Message) UpdateID() *Message {
	msg.Header.ID = uuid.NewV4().String()
	return msg
}

// BuildHeader builds message header. You can also use for updating message header
func (msg *Message) BuildHeader(ID, parentID string, timestamp int64) *Message {
	if msg.Header == nil {
		msg.Header = new(MessageHeader)
	}
	msg.Header.ID = ID
	msg.Header.ParentID = parentID
	msg.Header.Timestamp = timestamp
	return msg
}

//FillBody fills message  content that you want to send
// content interface need Marshal Any type successfully
func (msg *Message) FillBody(content interface{}) *Message {
	real, err := typeurl.MarshalAny(content)
	if err != nil {
		panic(errors.Errorf("Marshal content interface to any type error %v", err))
	}
	msg.Data = real
	return msg
}

// NewRawMessage returns a new raw message:
// model.NewRawMessage().BuildHeader().BuildRouter().FillBody()
func NewRawMessage() *Message {
	return &Message{
		Header: new(MessageHeader),
		Router: new(MessageRoute),
	}
}

// NewMessage returns a new basic message:
// model.NewMessage().BuildRouter().FillBody()
func NewMessage(parentID string) *Message {
	msg := NewRawMessage()
	msg.Header.ID = uuid.NewV4().String()
	msg.Header.ParentID = parentID
	msg.Header.Timestamp = time.Now().UnixNano() / 1e6
	return msg
}

// Clone a message
// only update message id
func (msg *Message) Clone(message *Message) *Message {
	msgID := uuid.NewV4().String()
	return NewRawMessage().BuildHeader(msgID, message.GetParentID(), message.GetTimestamp()).
		BuildRouter(message.GetSource(), message.GetGroup(), message.GetResource(), message.GetOperation()).
		FillBody(message.GetContent())
}

// NewRespByMessage returns a new response message by a message received
func (msg *Message) NewRespByMessage(message *Message, content interface{}) *Message {
	return NewMessage(message.GetID()).SetRoute(message.GetSource(), message.GetGroup()).
		SetResourceOperation(message.GetResource(), ResponseOperation).
		FillBody(content)
}

// NewErrorMessage returns a new error message by a message received
func NewErrorMessage(message *Message, errContent interface{}) *Message {
	return NewMessage(message.Header.ParentID).
		SetResourceOperation(message.Router.Resource, ResponseErrorOperation).
		FillBody(errContent)
}
