package douban

import (
	"holidayRemind/internal/holidayremind/douban"
	"testing"
)

func TestGetMovieWeeklyBest(t *testing.T) {
	resp := douban.CollectionResponse{}
	params := douban.CollectionParams{
		Start:     0,
		Count:     10,
		ItemsOnly: 0,
		ForMobile: 1,
	}
	err := douban.GetWeeklyBestByType(douban.Movie, params, &resp)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}

func TestGetChineseTVWeeklyBest(t *testing.T) {
	resp := douban.CollectionResponse{}
	params := douban.CollectionParams{
		Start:     0,
		Count:     10,
		ItemsOnly: 0,
		ForMobile: 1,
	}
	err := douban.GetWeeklyBestByType(douban.ChineseTV, params, &resp)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}
