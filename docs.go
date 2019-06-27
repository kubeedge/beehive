/*
#  #############################################
#  Copyright (c) 2019-2039 All rights reserved.
#  #############################################
#
#  Name:  docs.go
#  Date:  2019-06-27 15:09
#  Author:   zhangjie
#  Email:   iamzhangjie0619@163.com
#  Desc:
#
*/

package beehive

//go:generate  protoc  -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf   --gogo_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:./ ./pkg/core/model/message.proto


