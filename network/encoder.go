package network

import (
	"proto/command"
	"proto/logic"
	"reflect"
	"github.com/golang/protobuf/proto"
	"log"
	"strings"
)

type Encoder interface {
	Encode(data interface{})(raw []byte)
	Decode(raw []byte) (msgId uint32, data interface{})
}

type EncoderImplement struct {
}

func (Self *EncoderImplement) Encode(data interface{})(raw []byte){
	switch t := data.(type){
	case *command.CMD_ERROR_SC:
		return Self.EncodeCommand(data)
	case *command.CMD_AUTH_SC:
		return Self.EncodeCommand(data)
	case *command.CMD_REPLAY_SC:
		return Self.EncodeCommand(data)
	default:
		log.Println("t = ", t)
		return Self.EncodeLogic(data)
	}
}

func (Self *EncoderImplement) Decode(raw []byte) (msgId uint32, msg interface{}){
	msgObj := &command.CMD_BASE_CS{}
	err := proto.Unmarshal(raw, msgObj)
	if err == nil {
		id := command.CommandType(msgObj.Cmd)
		body := msgObj.Body

		if id == command.CommandType_CMD_LOGIC {
			logicObj := &command.CMD_LOGIC_CS{}
			err := proto.Unmarshal(body, logicObj)
			if err == nil {
				id := int32(logicObj.LogicType)
				msgName := "logic." + logic.LogicType_name[int32(id)] + "_CS"
				logicType := proto.MessageType(msgName)
				obj := reflect.New(logicType).Interface()
				err := proto.Unmarshal(logicObj.LogicPkg, obj.(proto.Message))
				if err == nil {
					msgId = uint32(id)
					msg = obj
				}
			}
		} else {
			msgName := "command." + command.CommandType_name[int32(id)] + "_CS"
			commandType := proto.MessageType(msgName).Elem()
			obj := reflect.New(commandType).Interface()
			err := proto.Unmarshal(body, obj.(proto.Message))
			if err == nil {
				msgId = uint32(id)
				msg = obj
			}
		}
	}
	return
}

func (Self *EncoderImplement) EncodeCommand(msg interface{})(raw []byte){
	data, err := proto.Marshal((msg).(proto.Message))
	if err == nil {
		commandName := proto.MessageName(msg.(proto.Message))
		commandName = strings.TrimSuffix(commandName, "_SC")
		commandName = strings.TrimPrefix(commandName, "command.")
		baseObj := command.CMD_BASE_SC{}
		baseObj.Cmd = command.CommandType(command.CommandType_value[commandName])
		baseObj.Body = data
		raw, err = proto.Marshal(&baseObj)
		if err == nil {
			return
		}
	}

	return nil
}

func (Self *EncoderImplement) EncodeLogic(msg interface{})(raw []byte){
	data, err := proto.Marshal(msg.(proto.Message))
	if err == nil {
		logicName := proto.MessageName(msg.(proto.Message))
		logicName = strings.TrimSuffix(logicName, "_SC")
		logicName = strings.TrimPrefix(logicName, "logic.")
		logicObj := command.CMD_LOGIC_SC{}
		logicObj.LogicType = uint32(logic.LogicType_value[logicName])
		logicObj.LogicPkg = data
		return Self.EncodeCommand(logicObj)
	}

	return nil
}
