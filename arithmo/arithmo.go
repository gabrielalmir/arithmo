package arithmo

import (
	"fmt"
	"strconv"
	"sync"
)

type Storage struct {
	items sync.Map
}

func (s *Storage) Set(key string, value any) {
	s.items.Store(key, value)
}

func (s *Storage) Get(key string) (any, bool) {
	val, ok := s.items.Load(key)
	if !ok {
		return nil, false
	}
	return val, true
}

func (s *Storage) Del(key string) bool {
	_, ok := s.items.Load(key)
	s.items.Delete(key)
	return ok
}

func (s *Storage) Exists(key string) bool {
	_, ok := s.items.Load(key)
	return ok
}

func (s *Storage) Type(key string) string {
	val, ok := s.items.Load(key)
	if !ok {
		return "none"
	}
	switch val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func (s *Storage) LPush(key string, values ...any) (int, error) {
	val, _ := s.items.Load(key)
	var list []any

	if val != nil {
		switch v := val.(type) {
		case []any:
			list = v
		default:
			return s.Count(), fmt.Errorf("ERR value is not a list")
		}
	}

	list = append(values, list...)
	s.items.Store(key, list)

	return s.Count(), nil
}

func (s *Storage) RPop(key string) (any, error) {
	val, ok := s.items.Load(key)
	if !ok {
		return nil, fmt.Errorf("ERR no such key")
	}

	list, ok := val.([]any)
	if !ok {
		return nil, fmt.Errorf("ERR value is not a list")
	}

	if len(list) == 0 {
		return nil, fmt.Errorf("ERR list is empty")
	}

	lastElement := list[len(list)-1]
	list = list[:len(list)-1]
	s.items.Store(key, list)

	return lastElement, nil
}

func (s *Storage) Count() int {
	length := 0
	s.items.Range(func(_, _ any) bool {
		length++
		return true
	})

	return length
}

func (s *Storage) Incr(key string) (int, error) {
	val, ok := s.items.Load(key)
	if !ok {
		val = 0
	}

	switch v := val.(type) {
	case int:
		newVal := v + 1
		s.items.Store(key, newVal)
		return newVal, nil
	case string:
		num, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("ERR value is not an integer")
		}
		newVal := num + 1
		s.items.Store(key, newVal)
		return newVal, nil
	default:
		return 0, fmt.Errorf("ERR value is not an integer")
	}
}

func (s *Storage) Decr(key string) (int, error) {
	val, ok := s.items.Load(key)
	if !ok {
		val = 0
	}

	switch v := val.(type) {
	case int:
		newVal := v - 1
		s.items.Store(key, newVal)
		return newVal, nil
	case string:
		num, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("ERR value is not an integer")
		}
		newVal := num - 1
		s.items.Store(key, newVal)
		return newVal, nil
	default:
		return 0, fmt.Errorf("ERR value is not an integer")
	}
}
