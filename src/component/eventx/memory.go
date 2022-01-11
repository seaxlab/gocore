package eventx

// MemoryEventStore is a event store for sync operations
type MemoryEventStore struct {
	async     bool
	listeners map[string][]Listener
	manager   *EventManager
}

// NewMemoryEventStore create a sync event store
func NewMemoryEventStore(async bool) *MemoryEventStore {
	return &MemoryEventStore{
		async:     async,
		listeners: make(map[string][]Listener),
	}
}

// Listen add a listener to a event
func (eventStore *MemoryEventStore) Listen(evtType string, listener Listener) {
	if _, ok := eventStore.listeners[evtType]; !ok {
		eventStore.listeners[evtType] = make([]Listener, 0)
	}

	eventStore.listeners[evtType] = append(eventStore.listeners[evtType], listener)
}

// Publish :publish a event
func (eventStore *MemoryEventStore) Publish(evtType string, evt interface{}) {
	if listeners, ok := eventStore.listeners[evtType]; ok {
		for _, listener := range listeners {
			if eventStore.async {
				go eventStore.manager.Call(evt, listener)
			} else {
				eventStore.manager.Call(evt, listener)
			}
		}
	}
}

// SetManager event manager
func (eventStore *MemoryEventStore) SetManager(manager *EventManager) {
	eventStore.manager = manager
}
