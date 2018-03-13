package network

import "reflect"

type MessageHandler interface {
	RegisterMessage(eventName string, callback func(*Event))
	OnEvent(event *Event)
}

type MessageHandlerImplement struct {
	messageMap map[string][]func(*Event)
}

func (Self *MessageHandlerImplement) RegisterMessage(eventName string, callback func(*Event)) {
	if Self.messageMap == nil{
		Self.messageMap = make(map[string][]func(*Event))
	}
	if v, ok := Self.messageMap[eventName]; ok {
		v = append(v, callback)
	} else {
		v := []func(*Event){callback}
		Self.messageMap[eventName] = v
	}
}

func (Self *MessageHandlerImplement) OnEvent(event *Event)  {
	if Self.messageMap == nil {
		return
	}

	eventName := reflect.ValueOf(event.Msg).Elem().Type().String()
	if v, ok := Self.messageMap[eventName]; ok {
		for i := 0; i < len(v); i++{
			v[i](event)
		}
	}
}