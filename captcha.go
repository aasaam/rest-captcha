package main

import (
	"bytes"
	"encoding/base64"
	"image/color"
	"image/jpeg"
	"math/rand"
	"time"

	"github.com/afocus/captcha"
)

// GetRandomCaptchaFont return random captcha font by language
func GetRandomCaptchaFont(lang string) []byte {
	rand.Seed(time.Now().UnixNano())
	if lang == "fa" {
		n := rand.Int() % len(CaptchaFontsFa)
		return CaptchaFontsFa[n]
	} else if lang == "ar" {
		n := rand.Int() % len(CaptchaFontsAr)
		return CaptchaFontsAr[n]
	}
	n := rand.Int() % len(CaptchaFontsEn)
	return CaptchaFontsEn[n]
}

// GenerateCaptcha return captcha number and base64 encoded image
func GenerateCaptcha(item *StorageItem, qualityInput int) string {
	cap := captcha.New()
	cap.SetSize(512, 128)
	cap.AddFontFromBytes(GetRandomCaptchaFont(item.Language))

	quality := minMaxDefault(qualityInput, 10, 95)

	cap.SetFrontColor(color.RGBA{uint8(GetRandomNumber(0, 64)), uint8(GetRandomNumber(0, 64)), uint8(GetRandomNumber(0, 64)), 0xff})
	cap.SetBkgColor(color.RGBA{uint8(GetRandomNumber(192, 256)), uint8(GetRandomNumber(192, 256)), uint8(GetRandomNumber(192, 256)), 0xff})

	if item.Level == LevelEasy {
		cap.SetDisturbance(32)
	} else if item.Level == LevelHard {
		cap.SetDisturbance(192)
	} else {
		cap.SetDisturbance(64)
	}

	img := cap.CreateCustom(item.IntlValue)
	buf := new(bytes.Buffer)

	jpeg.Encode(buf, img, &jpeg.Options{
		Quality: quality,
	})

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
