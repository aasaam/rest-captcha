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

func getRandomCaptchaFont(lang string) []byte {
	rand.Seed(time.Now().UnixNano())
	if lang == "fa" {
		n := rand.Int() % len(captchaFontsFa)
		return captchaFontsFa[n]
	} else if lang == "ar" {
		n := rand.Int() % len(captchaFontsAr)
		return captchaFontsAr[n]
	}
	n := rand.Int() % len(captchaFontsEn)
	return captchaFontsEn[n]
}

func generateCaptcha(item *storageItem, qualityInput int) string {
	cap := captcha.New()
	cap.SetSize(512, 128)
	cap.AddFontFromBytes(getRandomCaptchaFont(item.language))

	quality := minMaxDefault(qualityInput, 5, 85)

	cap.SetFrontColor(color.RGBA{uint8(getRandomNumber(0, 64)), uint8(getRandomNumber(0, 64)), uint8(getRandomNumber(0, 64)), 0xff})
	cap.SetBkgColor(color.RGBA{uint8(getRandomNumber(192, 256)), uint8(getRandomNumber(192, 256)), uint8(getRandomNumber(192, 256)), 0xff})

	if item.level == levelEasy {
		cap.SetDisturbance(32)
	} else if item.level == levelHard {
		cap.SetDisturbance(192)
	} else {
		cap.SetDisturbance(64)
	}

	img := cap.CreateCustom(item.intlValue)
	buf := new(bytes.Buffer)

	jpeg.Encode(buf, img, &jpeg.Options{
		Quality: quality,
	})

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
