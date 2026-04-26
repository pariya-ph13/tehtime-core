package timesvc

import (
	"fmt"
	"time"
)

// ToJalaliString converts a time.Time to a Jalali (Solar Hijri) ISO-like string (YYYY-MM-DDTHH:MM:SSZ) in UTC.
func ToJalaliString(t time.Time) string {
	t = t.UTC()
	jy, jm, jd := gregorianToJalali(t.Year(), int(t.Month()), t.Day())
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ", jy, jm, jd, t.Hour(), t.Minute(), t.Second())
}

// gregorianToJalali converts a Gregorian date to Jalali.
// Implementation based on the algorithm commonly used for Persian calendar conversions.
func gregorianToJalali(gy, gm, gd int) (jy, jm, jd int) {
	g_d_m := [12]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	var gy2 int
	if gy > 1600 {
		jy = 979
		gy -= 1600
	} else {
		jy = 0
		gy -= 621
	}
	if gm > 2 {
		gy2 = gy + 1
	} else {
		gy2 = gy
	}
	days := 365*gy + (gy2+3)/4 - (gy2+99)/100 + (gy2+399)/400 - 80 + gd + g_d_m[gm-1]
	jy += 33 * (days / 12053)
	days %= 12053
	jy += 4 * (days / 1461)
	days %= 1461
	if days > 365 {
		jy += (days - 1) / 365
		days = (days - 1) % 365
	}
	if days < 186 {
		jm = 1 + days/31
		jd = 1 + days%31
	} else {
		jm = 7 + (days-186)/30
		jd = 1 + (days-186)%30
	}
	jy += 1 // Adjust to Solar Hijri year number
	return
}
