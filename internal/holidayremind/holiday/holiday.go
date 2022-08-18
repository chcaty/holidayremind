package holiday

import (
	"errors"
	"github.com/golang-module/carbon"
	"holidayRemind/internal/pkg"
	"strconv"
	"time"
)

var nowCarbon = carbon.Time2Carbon(time.Now())

func GetDayProp(date string, calendar map[string]DayProperty, dateProperty *DayProperty) error {
	if value, ok := calendar[date]; ok {
		*dateProperty = value
		return nil
	}
	return errors.New("calendar not contain date")
}

func CreateCalendar(calendar *map[string]DayProperty) error {
	// 根据法定节假日设置日历
	config := holidayConfig{}
	err := getHolidayConfig(&config)
	if err != nil {
		return err
	}
	// 本年开始时间
	yearStartDate := nowCarbon.StartOfYear()
	// 本年结束时间
	yearEndDate := nowCarbon.EndOfYear()
	currentDate := yearStartDate
	for currentDate.DiffInDays(yearEndDate) > 0 {
		currentDateStr := currentDate.ToDateString()
		property := DayProperty{}
		err := setHoliday(&property, currentDateStr, config)
		if err != nil {
			return err
		}
		(*calendar)[currentDateStr] = property
		currentDate = currentDate.AddDay()
	}
	return nil
}

func setHoliday(property *DayProperty, currentDay string, config holidayConfig) error {
	current := carbon.Parse(currentDay)
	// 根据周几初始化日历
	property.IsHoliday = false
	property.Description = "工作"
	if current.IsSaturday() {
		flag := isBigWeek(current.Carbon2Time())
		if flag {
			property.IsHoliday = true
			property.Description = "休息"
		}
	} else if current.IsSunday() {
		property.IsHoliday = true
		property.Description = "休息"
	}

	if value, ok := config.Holidays[currentDay]; ok {
		property.Description = value
		property.IsHoliday = true
	}

	if value, ok := config.SpecialWorkingDays[currentDay]; ok {
		property.Description = value
		property.IsHoliday = false
	}
	return nil
}

func getHolidayConfig(holidayConfig *holidayConfig) error {
	fileName := "holiday" + strconv.Itoa(nowCarbon.Year())
	err := pkg.GetConfigByJson(fileName, holidayConfig)
	if err != nil {
		return err
	}
	return nil
}

func isBigWeek(currentDate time.Time) bool {
	// 挑选一个大周作为大周的判断日期
	flagDate := time.Date(2022, 8, 8, 0, 0, 0, 0, time.Local)
	weekCount := carbon.Time2Carbon(flagDate).DiffAbsInWeeks(carbon.Time2Carbon(currentDate))
	if weekCount%2 == 0 {
		return true
	}
	return false
}
