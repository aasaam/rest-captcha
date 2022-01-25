package main

import (
	"testing"
	"time"
)

func TestGenerateID(t *testing.T) {
	id1 := GenerateID()
	id2 := GenerateID()
	if id1 == id2 {
		t.Errorf("id must be unique")
	}
}
func TestConvertStringToUInt64(t *testing.T) {
	num := ConvertStringToUInt64("a")
	if num != 0 {
		t.Errorf("num must be zero")
	}
}
func TestGetLevel(t *testing.T) {
	if GetLevel("1") != LevelEasy {
		t.Errorf("invalid level")
	}
	if GetLevel("EaSy") != LevelEasy {
		t.Errorf("invalid level")
	}
	if GetLevel("2") != LevelHard {
		t.Errorf("invalid level")
	}
	if GetLevel("hArd") != LevelHard {
		t.Errorf("invalid level")
	}
	if GetLevel("else") != LevelMedium {
		t.Errorf("invalid level")
	}
}

func TestStorage1(t *testing.T) {
	storage := NewStorage()
	if storage.Count() != 0 {
		t.Errorf("count must be zero")
	}
	storage.CleanUp()
	item := storage.NewItem(LevelHard, "fa", 3)
	mustTrue := storage.Validate(item.ID, item.Value)
	if mustTrue != true {
		t.Errorf("item validation must be true")
	}
}

func TestStorage2(t *testing.T) {
	storage := NewStorage()
	item := storage.NewItem(LevelEasy, "ar", 3)
	mustTrue := storage.Validate(item.ID, item.Value)
	if mustTrue != true {
		t.Errorf("item validation must be true")
	}
	mustFalse := storage.Validate(item.ID, item.Value)
	if mustFalse != false {
		t.Errorf("item validation must be false, double checked")
	}
}

func TestStorage3(t *testing.T) {
	storage := NewStorage()
	item := storage.NewItem(LevelMedium, "en", 0)
	time.Sleep(1 * time.Second)
	mustFalse := storage.Validate(item.ID, item.Value)
	if mustFalse != false {
		t.Errorf("item validation must be false, expired")
	}
}

func TestStorage4(t *testing.T) {
	storage := NewStorage()
	item1 := storage.NewItem(LevelMedium, "en", 0)
	item2 := storage.NewItem(999, "fa", 3)
	time.Sleep(1 * time.Second)
	mustFalse := storage.Validate(item1.ID, item1.Value)
	if mustFalse != false {
		t.Errorf("item validation must be false, expired")
	}
	mustTrue := storage.Validate(item2.ID, item2.Value)
	if mustTrue != true {
		t.Errorf("item validation must be true")
	}
}

func TestStorage5(t *testing.T) {
	storage := NewStorage()
	storage.NewItem(LevelMedium, "en", 0)
	storage.NewItem(LevelMedium, "ar", 0)
	storage.CleanUp()
	storage.Count()
	if storage.Count() != 0 {
		t.Errorf("count must be zero")
	}
}

func BenchmarkGenerateID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateID()
	}
}

func BenchmarkGenerateProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateProblem(LevelMedium)

	}
}
