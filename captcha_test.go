package main

import (
	"testing"
)

func TestGenerateCaptcha(t *testing.T) {
	storage := NewStorage()
	storage.CleanUp()

	itemFa := storage.NewItem(LevelEasy, "fa", 3)
	itemAr := storage.NewItem(LevelMedium, "ar", 3)
	itemEn := storage.NewItem(LevelHard, "en", 3)

	GenerateCaptcha(itemFa, 30)
	GenerateCaptcha(itemAr, 30)
	GenerateCaptcha(itemEn, 10)
}

func BenchmarkGenerateCaptcha(b *testing.B) {
	storage := NewStorage()
	for i := 0; i < b.N; i++ {
		GenerateCaptcha(storage.NewItem(LevelHard, "fa", 3), 20)
	}
}
