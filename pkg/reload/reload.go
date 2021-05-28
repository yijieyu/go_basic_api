package reload

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type ReloadPlugin interface {
	OnReload(event Event) error
	Events() []string
}

type OnReloadFunc func(event Event) error

type reloadPlugin struct {
	name   string
	events Event
	f      OnReloadFunc
}

type OnReload struct {
	plugins map[string]reloadPlugin
}

func New() *OnReload {
	return &OnReload{
		plugins: map[string]reloadPlugin{},
	}
}

func (r *OnReload) Register(name string, f OnReloadFunc, events ...string) {
	if _, ok := r.plugins[name]; ok {
		panic("duplicate register reload components by " + name)
	}

	r.plugins[name] = reloadPlugin{name: name, f: f, events: NewEvent(events)}
}

func (r *OnReload) Run(events ...string) error {
	for _, p := range r.plugins {
		e := p.intersection(events)
		if e.Empty() {
			return nil
		}

		if err := p.f(e); err != nil {
			return fmt.Errorf("reload %s component fail", p.name)
		}

		logrus.WithFields(logrus.Fields{
			"events":        events,
			"plugin":        p.name,
			"reload_events": e.String(),
		}).Infof("%s reload success", p.name)
	}

	return nil
}

func (p *reloadPlugin) intersection(es []string) Event {
	// 事件列表是空的，表示要刷新所有缓存
	if len(es) == 0 {
		return p.events
	}

	e := make(Event)
	for _, v := range es {
		if p.events.Exist(v) {
			e[v] = struct{}{}
		}
	}

	return e
}
