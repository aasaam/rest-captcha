package main

import (
	"testing"
)

func TestGenerateCaptcha(t *testing.T) {
	storage := newStorage(true)
	storage.cleanUp()

	itemFa := storage.newItem(levelEasy, "fa", 3)
	itemAr := storage.newItem(levelMedium, "ar", 3)
	itemEn := storage.newItem(levelHard, "en", 3)

	generateCaptcha(itemFa, 30)
	generateCaptcha(itemAr, 30)
	generateCaptcha(itemEn, 10)
}

func BenchmarkGenerateCaptcha(b *testing.B) {
	storage := newStorage(true)
	for i := 0; i < b.N; i++ {
		generateCaptcha(storage.newItem(levelHard, "fa", 3), 20)
	}
}
