package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInteraface interface {
	Handle(event EventInterface, waitGroup *sync.WaitGroup)
}

type EventDisptacherInterface interface {
	Register(eventName string, handler EventHandlerInteraface) error
	Disptach(event EventInterface) error
	Remove(eventName string, handler EventHandlerInteraface) error
	Has(eventName string, handler EventHandlerInteraface) bool
	Clear() error
}
