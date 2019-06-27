package model

import (
	"time"

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
func (msg *Message) BuildRouter(source, group string, res *Resource, opr string) *Message {
	msg.SetRoute(source, group)
	msg.SetResourceOperation(res, opr)
	return msg
}

//SetResourceOperation sets router resource and operation in message
func (msg *Message) SetResourceOperation(res *Resource, opr string) *Message {
	msg.Router.Resource = res
	msg.Router.Operation = opr
	return msg
}

//SetRoute sets router source and group in message
func (msg *Message) SetRoute(source, group string) *Message {
	msg.Router.Source = source
	msg.Router.Group = group
	return msg
}

// IsSync : msg.Header.Sync will be set in sendsync
func (msg *Message) IsSync() bool {
	return msg.Header.Sync
}

//GetResource returns message route resource
func (msg *Message) GetResource() *Resource {
	return msg.Router.Resource
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

//UpdateID returns message object updating its ID
func (msg *Message) UpdateID() *Message {
	msg.Header.ID = uuid.NewV4().String()
	return msg
}

// BuildHeader builds message header. You can also use for updating message header
func (msg *Message) BuildHeader(ID, parentID string, timestamp int64) *Message {
	msg.Header.ID = ID
	msg.Header.ParentID = parentID
	msg.Header.Timestamp = timestamp
	return msg
}

//FillBody fills message  content that you want to send
func (msg *Message) FillBody(content []byte) *Message {
	msg.Content = content
	return msg
}

// NewRawMessage returns a new raw message:
// model.NewRawMessage().BuildHeader().BuildRouter().FillBody()
func NewRawMessage() *Message {
	return &Message{}
}

// NewMessage returns a new basic message:
// model.NewMessage().BuildRouter().FillBody()
func NewMessage(parentID string) *Message {
	msg := &Message{}
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
func (msg *Message) NewRespByMessage(message *Message, content []byte) *Message {
	return NewMessage(message.GetID()).SetRoute(message.GetSource(), message.GetGroup()).
		SetResourceOperation(message.GetResource(), ResponseOperation).
		FillBody(content)
}

// NewErrorMessage returns a new error message by a message received
func NewErrorMessage(message *Message, errContent []byte) *Message {
	return NewMessage(message.Header.ParentID).
		SetResourceOperation(message.Router.Resource, ResponseErrorOperation).
		FillBody(errContent)
}
