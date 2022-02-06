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
var zeroStringRegex = regexp.MustCompile(`[0]{1}`)

func newStorage() *storage {
	storage := storage{}
	storage.values = make(map[string]uint64)
	storage.expire = make(map[string]int64)
	return &storage
}

func (s *storage) count() int {
	return len(s.values)
}

func (s *storage) newItem(level int, lang string, ttl int64) *storageItem {
	s.mu.Lock()
	defer s.mu.Unlock()

	expireTime := time.Now()
	expireTime = expireTime.Add(time.Second * time.Duration(ttl))

	id := generateID()
	value := generateProblem(level)
	intlValue := convertUInt64ToString(value)

	if lang == "fa" {
		intlValue = message.NewPrinter(language.Persian).Sprintf("%v", number.Decimal(value, number.NoSeparator()))
	} else if lang == "ar" {
		intlValue = message.NewPrinter(language.Arabic).Sprintf("%v", number.Decimal(value, number.NoSeparator()))
	}
	item := storageItem{id: id, value: value, language: lang, intlValue: intlValue, level: level, expire: expireTime}
	s.expire[id] = expireTime.Unix()
	s.values[id] = value
	return &item
}

func (s *storage) validate(id string, value uint64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if exp, ok := s.expire[id]; ok {
		if exp >= time.Now().Unix() && s.values[id] == value {
			delete(s.expire, id)
			delete(s.values, id)
			prometheusValidTotal.Inc()
			return true
		}
	}
	prometheusInValidTotal.Inc()
	return false
}

func (s *storage) cleanUp() {
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

func generateID() string {
	b1 := make([]byte, 16)
	_, err := rand.Read(b1)

	if err != nil {
		panic(err.Error())
	}

	return safeStringRegex.ReplaceAllString(base64.StdEncoding.EncodeToString(b1), "")[0:12]
}

func generateProblem(level int) uint64 {
	var num int
	var min int
	var max int
	if level == levelEasy {
		min = 10000
		max = 99999
	} else if level == levelHard {
		min = 1000000
		max = 9999999
	} else { // medium is default
		min = 100000
		max = 999999
	}
	num = getRandomNumber(min, max)
	numberString := strconv.Itoa(num)
	numberString = zeroStringRegex.ReplaceAllStringFunc(numberString, func(m string) string {
		return strconv.Itoa(getRandomNumber(1, 9))
	})
	return convertStringToUInt64(numberString)
}

func getRandomNumber(min int, max int) int {
	math_rand.Seed(time.Now().UnixNano())
	return math_rand.Intn(max-min) + min
}

func convertStringToUInt64(str string) uint64 {
	value, e := strconv.ParseInt(str, 10, 64)
	if e != nil {
		return 0
	}
	return uint64(value)
}

func convertUInt64ToString(in uint64) string {
	return strconv.FormatUint(in, 10)
}

func getLevel(str string) int {
	if str == "1" || strings.ToUpper(str) == "EASY" {
		return levelEasy
	}
	if str == "2" || strings.ToUpper(str) == "HARD" {
		return levelHard
	}
	return levelMedium
}
