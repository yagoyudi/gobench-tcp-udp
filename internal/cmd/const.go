package cmd

import "fmt"

const (
	oneMB         = 1024 * 1024 * 1024
	hundredMB     = 100 * oneMB
	fiveHundredMB = 500 * oneMB
	oneGB         = 1024 * 1024 * 1024
	fiveGB        = 5 * oneGB
	tenGB         = 10 * oneGB
)

func parseTotalFlag(total string) (int, error) {
	switch total {
	case "100mb":
		return hundredMB, nil
	case "500mb":
		return fiveHundredMB, nil
	case "1gb":
		return oneGB, nil
	case "5gb":
		return fiveGB, nil
	case "10gb":
		return tenGB, nil
	}
	return 0, fmt.Errorf("invalid amount")
}
