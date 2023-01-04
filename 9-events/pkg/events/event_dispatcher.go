package events

import (
	"errors"
	"sync"
)

var ErrorHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDisptacher struct {
	handlers map[string][]EventHandlerInteraface
}

func NewEventDispatcher() *EventDisptacher {
	return &EventDisptacher{
		handlers: make(map[string][]EventHandlerInteraface),
	}
}

func (ed *EventDisptacher) Register(eventName string, handler EventHandlerInteraface) error {

	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrorHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDisptacher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInteraface)
}

func (ed *EventDisptacher) Has(eventName string, handler EventHandlerInteraface) bool {

	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (ed *EventDisptacher) Dispatch(event EventInterface) error {

	if handlers, ok := ed.handlers[event.GetName()]; ok {
		waitGroup := &sync.WaitGroup{}
		for _, handler := range handlers {
			waitGroup.Add(1)
			go handler.Handle(event, waitGroup)
		}
		waitGroup.Wait()
	}

	return nil
}

func (ed *EventDisptacher) Remove(eventName string, handlerRef EventHandlerInteraface) error {

	if handlers, ok := ed.handlers[eventName]; ok {
		for index, handler := range handlers {
			if handler == handlerRef {
				ed.handlers[eventName] = append(ed.handlers[eventName][:index], ed.handlers[eventName][index+1:]...)
				return nil
			}
		}
	}

	return nil
}
