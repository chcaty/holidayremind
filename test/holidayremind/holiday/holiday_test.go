package holiday

import (
	"holidayRemind/internal/holidayremind/holiday"
	"testing"
)

func TestCreateCalendar(t *testing.T) {
	var calendar = map[string]holiday.DayProperty{}
	err := holiday.CreateCalendar(&calendar)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(calendar)
}
