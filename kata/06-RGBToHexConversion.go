package kata

import (
	"fmt"
)

/*
	The rgb function is incomplete. Complete it so that passing in RGB decimal values will result in a hexadecimal representation being returned. Valid decimal values for RGB are 0 - 255.
	Any values that fall out of that range must be rounded to the closest valid value.

	Note: Your answer should always be 6 characters long, the shorthand with 3 will not work here.

	The following are examples of expected output values:

	kata.rgb(255, 255, 255) -- returns FFFFFF
	kata.rgb(255, 255, 300) -- returns FFFFFF
	kata.rgb(0, 0, 0) -- returns 000000
	kata.rgb(148, 0, 211) -- returns 9400D3
*/

func RGB(r, g, b int) string {
	rgbArr := []int{r, g, b}

	for i, v := range rgbArr {
		fmt.Println(v)
		if v < 0 {
			v = 0
		}

		if v > 255 {
			v = 255
		}

		rgbArr[i] = v
	}

	hex := fmt.Sprintf("%02X%02X%02X", rgbArr[0], rgbArr[1], rgbArr[2])

	return hex
}
