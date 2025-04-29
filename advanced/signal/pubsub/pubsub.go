package pubsub

import (
	"errors"
	"github.com/c3b2a7/HelloGo/advanced/signal/done"
	"github.com/c3b2a7/HelloGo/advanced/task"
	"sync"
	"time"
)

type Subscriber struct {
	buffer chan interface{}
	done   *done.Instance
}

func (s *Subscriber) push(msg interface{}) {
	select {
	case s.buffer <- msg:
	default:
	}
}

func (s *Subscriber) Wait() <-chan interface{} {
	return s.buffer
}

func (s *Subscriber) Close() error {
	return s.done.Close()
}

func (s *Subscriber) IsClosed() bool {
	return s.done.Done()
}

type Service struct {
	lock  sync.RWMutex
	subs  map[string][]*Subscriber
	ctask *task.Periodic
}

func NewService() *Service {
	s := &Service{
		subs: make(map[string][]*Subscriber),
	}
	s.ctask = &task.Periodic{
		Execute:  s.Cleanup,
		Interval: time.Second * 30,
	}
	return s
}

// Cleanup cleans up internal caches of subscribers.
// Visible for testing only.
func (s *Service) Cleanup() error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.subs) == 0 {
		return errors.New("nothing to do")
	}

	for name, subs := range s.subs {
		newSub := make([]*Subscriber, 0, len(s.subs))
		for _, sub := range subs {
			if !sub.IsClosed() {
				newSub = append(newSub, sub)
			}
		}
		if len(newSub) == 0 {
			delete(s.subs, name)
		} else {
			s.subs[name] = newSub
		}
	}

	if len(s.subs) == 0 {
		s.subs = make(map[string][]*Subscriber)
	}
	return nil
}

func (s *Service) Subscribe(name string) *Subscriber {
	sub := &Subscriber{
		buffer: make(chan interface{}, 16),
		done:   done.New(),
	}
	s.lock.Lock()
	s.subs[name] = append(s.subs[name], sub)
	s.lock.Unlock()
	if err := s.ctask.Start(); err != nil {
		panic(err)
	}
	return sub
}

func (s *Service) Publish(name string, message interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, sub := range s.subs[name] {
		if !sub.IsClosed() {
			sub.push(message)
		}
	}
}
