package web

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jalaali/go-jalaali"
)

type CalendarData struct {
	Calendar   string
	Year       int
	Month      int
	Weeks      [][]*DayCell
	PrevLink   string
	NextLink   string
}

type DayCell struct {
	Y   int
	M   int
	D   int
	InMonth bool
}

func NewWebHandler() *Handler { return &Handler{} }

type Handler struct{}

// GET /  -> full page
// GET /calendar -> partial month grid (htmx)
func (h *Handler) Index(c *fiber.Ctx) error {
	cd := buildCalendarFromQuery(c)
	return c.Render("index", cd)
}

func (h *Handler) CalendarPartial(c *fiber.Ctx) error {
	cd := buildCalendarFromQuery(c)
	return c.Render("_calendar", cd)
}

func buildCalendarFromQuery(c *fiber.Ctx) CalendarData {
	cal := c.Query("calendar")
	if cal == "" { cal = "solar" }
	y, _ := strconv.Atoi(defaultIfEmpty(c.Query("year"), "0"))
	m, _ := strconv.Atoi(defaultIfEmpty(c.Query("month"), "0"))

	now := time.Now().UTC()
	if cal == "gregorian" {
		if y == 0 { y = now.Year() }
		if m == 0 { m = int(now.Month()) }
		weeks := monthGridGregorian(y, m)
		py, pm := prevMonthGregorian(y, m)
		ny, nm := nextMonthGregorian(y, m)
		return CalendarData{
			Calendar: cal,
			Year: y, Month: m,
			Weeks: weeks,
			PrevLink: link("/calendar", cal, py, pm),
			NextLink: link("/calendar", cal, ny, nm),
		}
	}
	// solar
	jy, jm, _ := jalaali.ToJalaali(now.Year(), int(now.Month()), now.Day())
	if y == 0 { y = jy }
	if m == 0 { m = jm }
	weeks := monthGridJalali(y, m)
	py, pm := prevMonthJalali(y, m)
	ny, nm := nextMonthJalali(y, m)
	return CalendarData{
		Calendar: cal,
		Year: y, Month: m,
		Weeks: weeks,
		PrevLink: link("/calendar", cal, py, pm),
		NextLink: link("/calendar", cal, ny, nm),
	}
}

func link(path, cal string, y, m int) string {
	q := url.Values{}
	q.Set("calendar", cal)
	q.Set("year", fmt.Sprintf("%d", y))
	q.Set("month", fmt.Sprintf("%d", m))
	return path + "?" + q.Encode()
}

func defaultIfEmpty(s, d string) string { if s == "" { return d }; return s }

func monthGridGregorian(y, m int) [][]*DayCell {
	first := time.Date(y, time.Month(m), 1, 0,0,0,0, time.UTC)
	w := int(first.Weekday()) // 0=Sun
	days := daysInGregorianMonth(y, m)
	var grid [][]*DayCell
	row := make([]*DayCell, 0, 7)
	// leading blanks from previous month
	py, pm := prevMonthGregorian(y, m)
	pdays := daysInGregorianMonth(py, pm)
	for i:=0; i<w; i++ { row = append(row, &DayCell{Y:py, M:pm, D:pdays-w+i+1, InMonth:false}) }
	for d:=1; d<=days; d++ {
		row = append(row, &DayCell{Y:y,M:m,D:d,InMonth:true})
		if len(row)==7 { grid = append(grid, row); row = make([]*DayCell,0,7) }
	}
	// trailing cells from next month
	ny, nm := nextMonthGregorian(y,m)
	for len(row)>0 && len(row)<7 { row = append(row, &DayCell{Y:ny,M:nm,D:len(row),InMonth:false}) }
	if len(row)>0 { grid = append(grid, row) }
	return grid
}

func daysInGregorianMonth(y, m int) int {
	if m==2 { if (y%4==0 && y%100!=0) || (y%400==0) { return 29 }; return 28 }
	if m==4 || m==6 || m==9 || m==11 { return 30 }
	return 31
}

func prevMonthGregorian(y, m int) (int,int) { if m==1 { return y-1, 12 }; return y, m-1 }
func nextMonthGregorian(y, m int) (int,int) { if m==12 { return y+1, 1 }; return y, m+1 }

func monthGridJalali(y, m int) [][]*DayCell {
	gy, gm, _ := jalaali.ToGregorian(y, m, 1)
	first := time.Date(gy, time.Month(gm), 1, 0,0,0,0, time.UTC)
	w := int(first.Weekday()) // anchor by Sunday to match Fiber default locale
	days := daysInJalaliMonth(y, m)
	var grid [][]*DayCell
	row := make([]*DayCell, 0, 7)
	py, pm := prevMonthJalali(y, m)
	pdays := daysInJalaliMonth(py, pm)
	for i:=0; i<w; i++ { row = append(row, &DayCell{Y:py, M:pm, D:pdays-w+i+1, InMonth:false}) }
	for d:=1; d<=days; d++ {
		row = append(row, &DayCell{Y:y,M:m,D:d,InMonth:true})
		if len(row)==7 { grid = append(grid, row); row = make([]*DayCell,0,7) }
	}
	ny, nm := nextMonthJalali(y,m)
	for len(row)>0 && len(row)<7 { row = append(row, &DayCell{Y:ny,M:nm,D:len(row),InMonth:false}) }
	if len(row)>0 { grid = append(grid, row) }
	return grid
}

func daysInJalaliMonth(y, m int) int {
	if m<=6 { return 31 }
	if m<=11 { return 30 }
	// Esfand (12): 29 normally, 30 in leap years
	if jalaali.IsLeapJalali(y) { return 30 }
	return 29
}

func prevMonthJalali(y, m int) (int,int) { if m==1 { return y-1, 12 }; return y, m-1 }
func nextMonthJalali(y, m int) (int,int) { if m==12 { return y+1, 1 }; return y, m+1 }
