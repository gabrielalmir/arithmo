package arithmo_test

import (
	"testing"

	"github.com/gabrielalmir/arithmo/arithmo"
)

func TestStorageSetGet(t *testing.T) {
	storage := &arithmo.Storage{}
	storage.Set("key1", "value1")

	val, ok := storage.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("expected value1, got %v", val)
	}
}

func TestStorageDel(t *testing.T) {
	storage := &arithmo.Storage{}
	storage.Set("key1", "value1")

	ok := storage.Del("key1")
	if !ok {
		t.Errorf("expected true, got false")
	}

	_, ok = storage.Get("key1")
	if ok {
		t.Errorf("expected false, got true")
	}
}

func TestStorageExists(t *testing.T) {
	storage := &arithmo.Storage{}
	storage.Set("key1", "value1")

	ok := storage.Exists("key1")
	if !ok {
		t.Errorf("expected true, got false")
	}

	storage.Del("key1")
	ok = storage.Exists("key1")
	if ok {
		t.Errorf("expected false, got true")
	}
}

func TestStorageType(t *testing.T) {
	storage := &arithmo.Storage{}
	storage.Set("key1", "value1")
	storage.Set("key2", 123)

	if typ := storage.Type("key1"); typ != "string" {
		t.Errorf("expected string, got %v", typ)
	}

	if typ := storage.Type("key2"); typ != "int" {
		t.Errorf("expected int, got %v", typ)
	}

	if typ := storage.Type("key3"); typ != "none" {
		t.Errorf("expected none, got %v", typ)
	}
}

func TestStorageIncr(t *testing.T) {
	storage := &arithmo.Storage{}
	storage.Set("key1", 1)

	newVal, err := storage.Incr("key1")
	if err != nil || newVal != 2 {
		t.Errorf("expected 2, got %v, err: %v", newVal, err)
	}

	storage.Set("key2", "1")
	newVal, err = storage.Incr("key2")
	if err != nil || newVal != 2 {
		t.Errorf("expected 2, got %v, err: %v", newVal, err)
	}

	storage.Set("key3", "not an int")
	_, err = storage.Incr("key3")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestStorageDecr(t *testing.T) {
	storage := &arithmo.Storage{}
	storage.Set("key1", 1)

	newVal, err := storage.Decr("key1")
	if err != nil || newVal != 0 {
		t.Errorf("expected 0, got %v, err: %v", newVal, err)
	}

	storage.Set("key2", "1")
	newVal, err = storage.Decr("key2")
	if err != nil || newVal != 0 {
		t.Errorf("expected 0, got %v, err: %v", newVal, err)
	}

	storage.Set("key3", "not an int")
	_, err = storage.Decr("key3")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
