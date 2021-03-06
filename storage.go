package main

import (
	"crypto/rand"
	"encoding/base64"
	math_rand "math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

var safeStringRegex = regexp.MustCompile(`[^a-zA-Z0-9]`)

// NewStorage will return new storage pool
func NewStorage() *Storage {
	storage := Storage{}
	storage.values = make(map[string]uint64)
	storage.expire = make(map[string]int64)
	return &storage
}

// Count will return number of
func (s *Storage) Count() int {
	return len(s.values)
}

// NewItem will register new StorageItem and return it
func (s *Storage) NewItem(level int, lang string, ttl int64) *StorageItem {
	s.mu.Lock()
	defer s.mu.Unlock()

	expireTime := time.Now()
	expireTime = expireTime.Add(time.Second * time.Duration(ttl))

	id := GenerateID()
	value := GenerateProblem(level)
	intlValue := ConvertUInt64ToString(value)

	if lang == "fa" {
		intlValue = message.NewPrinter(language.Persian).Sprintf("%v", number.Decimal(value, number.NoSeparator()))
	} else if lang == "ar" {
		intlValue = message.NewPrinter(language.Arabic).Sprintf("%v", number.Decimal(value, number.NoSeparator()))
	}
	item := StorageItem{ID: id, Value: value, Language: lang, IntlValue: intlValue, Level: level, Expire: expireTime}
	s.expire[id] = expireTime.Unix()
	s.values[id] = value
	return &item
}

// Validate id via input value
func (s *Storage) Validate(id string, value uint64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if exp, ok := s.expire[id]; ok {
		if exp >= time.Now().Unix() && s.values[id] == value {
			delete(s.expire, id)
			delete(s.values, id)
			PrometheusValidTotal.Inc()
			return true
		}
	}
	PrometheusInValidTotal.Inc()
	return false
}

// CleanUp remove expired
func (s *Storage) CleanUp() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().Unix()
	for id, exp := range s.expire {
		if exp <= now {
			delete(s.expire, id)
			delete(s.values, id)
		}
	}
}

// GenerateID return random id for each captcha
func GenerateID() string {
	b1 := make([]byte, 16)
	_, err := rand.Read(b1)

	if err != nil {
		panic(err.Error())
	}

	return safeStringRegex.ReplaceAllString(base64.StdEncoding.EncodeToString(b1), "")[0:12]
}

// GenerateProblem generate number problem
func GenerateProblem(level int) uint64 {
	var num int
	var min int
	var max int
	if level == LevelEasy {
		min = 10000
		max = 99999
	} else if level == LevelHard {
		min = 1000000
		max = 9999999
	} else { // medium is default
		min = 100000
		max = 999999
	}
	num = GetRandomNumber(min, max)
	randomNoneZero := GetRandomNumber(1, 9)
	numberString := strconv.Itoa(num)
	numberString = strings.ReplaceAll(numberString, "0", strconv.Itoa(randomNoneZero))
	return ConvertStringToUInt64(numberString)
}

// GetRandomNumber random number between min and max
func GetRandomNumber(min int, max int) int {
	math_rand.Seed(time.Now().UnixNano())
	return math_rand.Intn(max-min) + min
}

// ConvertStringToUInt64 string to uint64
func ConvertStringToUInt64(str string) uint64 {
	value, e := strconv.ParseInt(str, 10, 64)
	if e != nil {
		return 0
	}
	return uint64(value)
}

// ConvertUInt64ToString uint64 to string
func ConvertUInt64ToString(in uint64) string {
	return strconv.FormatUint(in, 10)
}

// GetLevel parse level string
func GetLevel(str string) int {
	if str == "1" || strings.ToUpper(str) == "EASY" {
		return LevelEasy
	}
	if str == "2" || strings.ToUpper(str) == "HARD" {
		return LevelHard
	}
	return LevelMedium
}
