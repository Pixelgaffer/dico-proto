package dicoprotos

import "github.com/golang/protobuf/proto"

func GetMessageType(m proto.Message) SelfDescribingMessage_MessageType {
	switch m.(type) {
	case *Handshake:
		return SelfDescribingMessage_HANDSHAKE
	case *DoTask:
		return SelfDescribingMessage_DO_TASK
	case *TaskStatus:
		return SelfDescribingMessage_TASK_STATUS
	case *TaskResult:
		return SelfDescribingMessage_TASK_RESULT
	case *SubmitTask:
		return SelfDescribingMessage_SUBMIT_TASK
	}
	panic("invalid msg type")
}

func GetSDMMessageType(msg *SelfDescribingMessage) proto.Message {
	switch msg.GetType() {
	case SelfDescribingMessage_HANDSHAKE:
		return new(Handshake)
	case SelfDescribingMessage_DO_TASK:
		return new(DoTask)
	case SelfDescribingMessage_TASK_STATUS:
		return new(TaskStatus)
	case SelfDescribingMessage_TASK_RESULT:
		return new(TaskResult)
	case SelfDescribingMessage_SUBMIT_TASK:
		return new(SubmitTask)
	}
	panic("invalid msg type")
}

func DecodeUnknownMessage(buf []byte) (proto.Message, error) {
	sdm := new(SelfDescribingMessage)
	err := proto.Unmarshal(buf, sdm)
	if err != nil {
		return nil, err
	}
	msg := GetSDMMessageType(sdm)
	err = proto.Unmarshal(sdm.GetData(), msg)
	return msg, err
}

func WrapMessage(msg proto.Message) proto.Message {
	msgtype := GetMessageType(msg)
	data, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	sdm := &SelfDescribingMessage{
		Type: &msgtype,
		Data: data,
	}
	return sdm
}
