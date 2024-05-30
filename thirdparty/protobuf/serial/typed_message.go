package serial

import (
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
)

const TypeURLPrefix = "types.lolico.me/"

// ToAnypbMessage converts a proto.Message into anypb.Any.
func ToAnypbMessage(message proto.Message) *anypb.Any {
	if message == nil {
		return nil
	}
	bytes, _ := proto.Marshal(message)
	return &anypb.Any{
		TypeUrl: TypeURLPrefix + GetMessageType(message),
		Value:   bytes,
	}
}

// ToProtoMessage converts an anypb.Any into proto.Message
func ToProtoMessage(v *anypb.Any) (proto.Message, error) {
	if v == nil {
		return nil, nil
	}
	instance, err := GetInstance(internalType(v))
	if err != nil {
		return nil, err
	}
	protoMessage := instance.(proto.Message)
	if err := proto.Unmarshal(v.Value, protoMessage); err != nil {
		return nil, err
	}
	return protoMessage, nil // equals to anypb.UnmarshalNew(v, proto.UnmarshalOptions{})
}

// GetMessageType returns the name of this proto Message.
func GetMessageType(message proto.Message) string {
	return string(proto.MessageName(message))
}

// GetInstance creates a new instance of the message with messageType.
func GetInstance(messageType string) (interface{}, error) {
	mt, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(messageType))
	if err != nil {
		return nil, err
	}
	return mt.New().Interface(), nil
}

func internalType(any *anypb.Any) string {
	return strings.TrimPrefix(any.TypeUrl, TypeURLPrefix)
}
