package memoryhook

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type MemoryHook struct {
	entries []*logrus.Entry
	mu      sync.Mutex
}

func New() *MemoryHook {
	return &MemoryHook{
		entries: make([]*logrus.Entry, 0),
	}
}

func (h *MemoryHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *MemoryHook) Fire(entry *logrus.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.entries = append(h.entries, entry)
	return nil
}

func (h *MemoryHook) Entries() []*logrus.Entry {
	h.mu.Lock()
	defer h.mu.Unlock()

	return h.entries
}
