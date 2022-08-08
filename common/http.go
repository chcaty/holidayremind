package common

import (
	"fmt"
	"io"
)

func CloseBody(body io.ReadCloser) {
	func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Request Body Close error: %s", err.Error())
		}
	}(body)
}
