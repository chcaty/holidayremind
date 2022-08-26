package moyuduck

import (
	"holidayRemind/internal/holidayremind/moyuduck"
	"testing"
)

func TestGetHoliday(t *testing.T) {
	holiday := moyuduck.Response[moyuduck.Holiday]{}
	err := moyuduck.GetHoliday(&holiday)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(holiday)
}

func TestGetImage(t *testing.T) {
	imageUrl := moyuduck.Response[string]{}
	err := moyuduck.GetImage(&imageUrl)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(imageUrl)
}

func TestGetHotTop(t *testing.T) {
	topSite := moyuduck.Response[moyuduck.HotTopSite]{}
	err := moyuduck.GetHotTop(&topSite)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(topSite.Msg)
}
