package pb_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"gitee.com/go-course/rpc-g7/protobuf/pb"
)

func TestMarshal(t *testing.T) {
	should := assert.New(t)

	str := &pb.String{Value: "hello"}

	// object --protobuf--> []byte
	pbBytes, err := proto.Marshal(str)
	if should.NoError(err) {
		fmt.Println(pbBytes)
	}

	// []byte --protobuf--> object
	obj := pb.String{}
	err = proto.Unmarshal(pbBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj.Value)
	}

	e := &pb.Event{}
	// 采用断言来判断时one of中的哪一种类型
	e.GetHeader()

	// 直接通过Get获取, 判断是否为nil
	e.GetProtobuf()
	e.GetJson()

	//      foo := &pb.Foo{...}
	//      any, err := anypb.New(foo)
	//      if err != nil {
	//        ...
	//      }
	//      ...
	//      foo := &pb.Foo{}
	//      if err := any.UnmarshalTo(foo); err != nil {
	//        ...
	//      }

	// any对象赋值
	es := &pb.ErrorStatus{Message: "test"}
	anyEs, err := anypb.New(es)
	if should.NoError(err) {
		fmt.Println(anyEs)
	}

	// any对象取值
	// json.Unmarshal(data, obj)
	objReceive := pb.ErrorStatus{}
	anyEs.UnmarshalTo(&objReceive)

}
