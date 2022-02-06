package main

import (
	"sync"
	"time"
)

type storage struct {
	mu     sync.Mutex
	values map[string]uint64
	expire map[string]int64
}

type storageItem struct {
	id        string
	value     uint64
	level     int
	intlValue string
	language  string
	expire    time.Time
}

type config struct {
	returnValue bool
	testImage   bool
}

const (
	levelMedium = 0
	levelEasy   = 1
	levelHard   = 2
)
