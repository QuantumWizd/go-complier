package chat

import "sync"

const MaxHistory = 50

type History struct {
	mu       sync.RWMutex
	messages []Message
}

func NewHistory() *History {
	return &History{
		messages: make([]Message, 0, MaxHistory),
	}
}

func (h *History) Add(m Message) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.messages) >= MaxHistory {
		h.messages = h.messages[1:]
	}
	h.messages = append(h.messages, m)
}

func (h *History) GetAll() []Message {
	h.mu.RLock()
	defer h.mu.RUnlock()

	result := make([]Message, len(h.messages))
	copy(result, h.messages)
	return result
}