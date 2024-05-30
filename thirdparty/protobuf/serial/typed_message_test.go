package serial

import (
	"fmt"
	"github.com/c3b2a7/HelloGo/thirdparty/protobuf"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"
)

func TestAnypbMessageAndProtoMessageCustomizeConverts(t *testing.T) {
	req := &protobuf.Request{
		Id:   0,
		Type: protobuf.Type_PING,
		Data: nil,
	}
	fmt.Println(GetMessageType(req))

	// converts req into anypb.Any
	anypbMessage := ToAnypbMessage(req)
	fmt.Println(GetMessageType(anypbMessage))

	// converts anypbMessage into proto.Message
	protoMessage, _ := ToProtoMessage(anypbMessage)
	fmt.Println(GetMessageType(protoMessage))
	if !proto.Equal(req, protoMessage) {
		t.Error()
	}
}

func TestAnypbMessageAndProtoMessageOfficalConverts(t *testing.T) {
	req := &protobuf.Request{
		Id:   0,
		Type: protobuf.Type_PING,
		Data: nil,
	}
	fmt.Println(GetMessageType(req))

	// converts req into anypb.Any
	anypbMessage, _ := anypb.New(req)
	fmt.Println(GetMessageType(anypbMessage))

	// converts anypbMessage into proto.Message
	protoMessage, _ := anypb.UnmarshalNew(anypbMessage, proto.UnmarshalOptions{})
	fmt.Println(GetMessageType(protoMessage))
	if !proto.Equal(req, protoMessage) {
		t.Error()
	}
}

func TestAnypbMessageAndProtoMessageMixedConverts(t *testing.T) {
	req := &protobuf.Request{
		Id:   0,
		Type: protobuf.Type_PING,
		Data: nil,
	}
	fmt.Println(GetMessageType(req))

	// converts req into anypb.Any
	anypbMessage := ToAnypbMessage(req)
	fmt.Println(GetMessageType(anypbMessage))

	// converts anypbMessage into proto.Message
	protoMessage, _ := anypb.UnmarshalNew(anypbMessage, proto.UnmarshalOptions{})
	fmt.Println(GetMessageType(protoMessage))
	if !proto.Equal(req, protoMessage) {
		t.Error()
	}
}
