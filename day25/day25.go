package day25

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"strconv"
)

func toSnafu(number int) string {
    snafu := ""

    for number != 0 {
        rem := number % 5
        number = number / 5

        if rem <= 2 {
            snafu = strconv.Itoa(rem) + snafu
        } else {
            snafu = string("   =-"[rem]) + snafu
            number += 1
        }
    }

    return snafu
}

func fromSnafu(snafu string) int {
	num := 0.0

	size := len(snafu)
	for i := size - 1; i >= 0; i-- {
		ch := snafu[i]
		digit := 0
		switch ch {
		case '-':
			digit = -1
		case '=':
			digit = -2
		default:
			var err error
			digit, err = strconv.Atoi(string(ch))
			if err != nil {
				log.Fatalln("NaN")
			}
		}

        if i == size - 1 {
            num = float64(digit)
        } else {
            j := size - i - 1
            num += float64(digit) * math.Pow(5, float64(j))
        }
	}

	return int(num)
}

func Day25() {
	contents := utils.GetFileContents("day25/input")

	sum := 0
	for _, line := range contents {
		sum += fromSnafu(line)
	}

	fmt.Println("Results part 1:", toSnafu(sum))
}
