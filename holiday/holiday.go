package holiday

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

func CreateCalendar() {
	now := time.Now()
	var holidayConfig = make(map[string]interface{})
	file, err := os.ReadFile("./json/holiday" + strconv.Itoa(now.Year()) + ".json")
	if err != nil {
		fmt.Printf("get holiday%s.json file fail. error: %s", strconv.Itoa(now.Year()), err.Error())
	}
	err = json.Unmarshal(file, &holidayConfig)
	if err != nil {
		fmt.Printf("get holidayConfig struct fail. error: %s", err.Error())
	}
	fmt.Printf("holidayConfig Struct: %#v\n", holidayConfig)
	var name = holidayConfig["2022-01-01"]
	fmt.Println(name)
}
