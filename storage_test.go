package main

import (
	"testing"
	"time"
)

func TestGenerateID(t *testing.T) {
	id1 := generateID()
	id2 := generateID()
	if id1 == id2 {
		t.Errorf("id must be unique")
	}
}
func TestConvertStringToUInt64(t *testing.T) {
	num := convertStringToUInt64("a")
	if num != 0 {
		t.Errorf("num must be zero")
	}
}
func TestGetLevel(t *testing.T) {
	if getLevel("1") != levelEasy {
		t.Errorf("invalid level")
	}
	if getLevel("EaSy") != levelEasy {
		t.Errorf("invalid level")
	}
	if getLevel("2") != levelHard {
		t.Errorf("invalid level")
	}
	if getLevel("hArd") != levelHard {
		t.Errorf("invalid level")
	}
	if getLevel("else") != levelMedium {
		t.Errorf("invalid level")
	}
}

func TestStorage1(t *testing.T) {
	storage := newStorage()
	if storage.count() != 0 {
		t.Errorf("count must be zero")
	}
	storage.cleanUp()
	item := storage.newItem(levelHard, "fa", 3)
	mustTrue := storage.validate(item.id, item.value)
	if mustTrue != true {
		t.Errorf("item validation must be true")
	}
}

func TestStorage2(t *testing.T) {
	storage := newStorage()
	item := storage.newItem(levelEasy, "ar", 3)
	mustTrue := storage.validate(item.id, item.value)
	if mustTrue != true {
		t.Errorf("item validation must be true")
	}
	mustFalse := storage.validate(item.id, item.value)
	if mustFalse != false {
		t.Errorf("item validation must be false, double checked")
	}
}

func TestStorage3(t *testing.T) {
	storage := newStorage()
	item := storage.newItem(levelMedium, "en", 0)
	time.Sleep(1 * time.Second)
	mustFalse := storage.validate(item.id, item.value)
	if mustFalse != false {
		t.Errorf("item validation must be false, expired")
	}
}

func TestStorage4(t *testing.T) {
	storage := newStorage()
	item1 := storage.newItem(levelMedium, "en", 0)
	item2 := storage.newItem(999, "fa", 3)
	time.Sleep(1 * time.Second)
	mustFalse := storage.validate(item1.id, item1.value)
	if mustFalse != false {
		t.Errorf("item validation must be false, expired")
	}
	mustTrue := storage.validate(item2.id, item2.value)
	if mustTrue != true {
		t.Errorf("item validation must be true")
	}
}

func TestStorage5(t *testing.T) {
	storage := newStorage()
	storage.newItem(levelMedium, "en", 0)
	storage.newItem(levelMedium, "ar", 0)
	storage.cleanUp()
	storage.count()
	if storage.count() != 0 {
		t.Errorf("count must be zero")
	}
}

func BenchmarkGenerateID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateID()
	}
}

func BenchmarkGenerateProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateProblem(levelMedium)

	}
}
