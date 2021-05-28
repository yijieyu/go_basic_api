package reload

import (
	"strings"
)

type Event map[string]struct{}

func NewEvent(es []string) Event {
	e := make(map[string]struct{}, len(es))
	for _, v := range es {
		e[v] = struct{}{}
	}

	return e
}

func (e Event) Exist(event string) bool {
	_, ok := e[event]
	return ok
}

func (e Event) Empty() bool {
	return len(e) == 0
}

func (e Event) String() string {
	buf := strings.Builder{}
	buf.WriteByte(',')
	for name := range e {
		buf.WriteString(name)
	}

	return buf.String()[1:]
}
