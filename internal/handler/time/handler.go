package timehandler

import (
	"strings"
	timepkg "time"

	"github.com/gofiber/fiber/v2"
	ptime "github.com/yaa110/go-persian-calendar"

	timesvc "github.com/TehranTime/tehtime-core/internal/service/time"
)

type Handler struct {
	Svc timesvc.Service
}

func NewHandler(svc timesvc.Service) *Handler {
	return &Handler{Svc: svc}
}

func (h *Handler) GetTime(c *fiber.Ctx) error {
	now := h.Svc.Now().UTC()
	cal := strings.ToLower(c.Query("calendar"))
	var ts, calOut string
	if cal == "gregorian" {
		ts = now.Format(timepkg.RFC3339)
		calOut = "gregorian"
	} else {
		pt := ptime.New(now)
		ts = pt.Format("2006-01-02T15:04:05Z07:00")
		calOut = "solar"
	}
	return c.JSON(fiber.Map{
		"time":     ts,
		"calendar": calOut,
	})
}
