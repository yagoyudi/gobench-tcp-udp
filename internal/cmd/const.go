package cmd

import "fmt"

const (
	tenMB         = 10 * 1024 * 1024
	hundredMB     = 100 * 1024 * 1024
	fiveHundredMB = 500 * 1024 * 1024
	oneGB         = 1024 * 1024 * 1024
)

func parseTotalFlag(total string) (int, error) {
	switch total {
	case "10mb":
		return tenMB, nil
	case "100mb":
		return hundredMB, nil
	case "500mb":
		return fiveHundredMB, nil
	case "1gb":
		return oneGB, nil
	}
	return 0, fmt.Errorf("invalid amount")
}
