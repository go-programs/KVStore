package kvstore

type EventType byte 

const (
	EventGet EventType = iota
	EventPut
	EventDel        
)

type Event struct {
	Sequence    uint64
	EventType   EventType
	Key         string
	Value       string
}

type EventLogger interface {
	writeDelete(key string) 
	writePut(key, value string)
	writeGet(key string)
}