package main

import (
	_ "embed"
)

// this file auto generate at 2021/07/17 21:21
// total files: 15
// total size: 124KB

//go:embed captcha-fonts/destination/ar_f535e32deed575c0be7f5de4bb26c582.ttf
var arFont0 []byte

//go:embed captcha-fonts/destination/ar_80c6577cb53980614d802892a47f06bc.ttf
var arFont1 []byte

//go:embed captcha-fonts/destination/ar_dc090a198e485cb26f9d1ae3151b22c4.ttf
var arFont2 []byte

//go:embed captcha-fonts/destination/ar_a352855c46d3976810828d59c0336e97.ttf
var arFont3 []byte

//go:embed captcha-fonts/destination/ar_9a6f6c585e8437bf82b12c334553c22a.ttf
var arFont4 []byte

//go:embed captcha-fonts/destination/fa_a9d08c2fe9dff20dd104ed181b86d50e.ttf
var faFont0 []byte

//go:embed captcha-fonts/destination/fa_70d7ba6093fd78d0682a485613f4d8d7.ttf
var faFont1 []byte

//go:embed captcha-fonts/destination/fa_9231ae413b4f4f2090d9ec221c314dee.ttf
var faFont2 []byte

//go:embed captcha-fonts/destination/fa_1f3d7e42443fa231634fb7f3eeafa697.ttf
var faFont3 []byte

//go:embed captcha-fonts/destination/fa_0be92be905af42ad6c6e1d7976cc9005.ttf
var faFont4 []byte

//go:embed captcha-fonts/destination/en_d95cec57d0707759c507ca199d09538f.ttf
var enFont0 []byte

//go:embed captcha-fonts/destination/en_b81eef48f35316d96d5ee558f06400f9.ttf
var enFont1 []byte

//go:embed captcha-fonts/destination/en_cfbe2858223b5d6f0fead9583b07f3fd.ttf
var enFont2 []byte

//go:embed captcha-fonts/destination/en_b949029e82c593c78cc0c0d8aeac7f71.ttf
var enFont3 []byte

//go:embed captcha-fonts/destination/en_57826df94499d69a08a017760adaafbd.ttf
var enFont4 []byte

// CaptchaFontsAr list of embed fonts for Ar
var CaptchaFontsAr = [][]byte{arFont0, arFont1, arFont2, arFont3, arFont4}

// CaptchaFontsFa list of embed fonts for Fa
var CaptchaFontsFa = [][]byte{faFont0, faFont1, faFont2, faFont3, faFont4}

// CaptchaFontsEn list of embed fonts for En
var CaptchaFontsEn = [][]byte{enFont0, enFont1, enFont2, enFont3, enFont4}
