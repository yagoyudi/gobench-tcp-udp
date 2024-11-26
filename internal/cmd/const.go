package cmd

const (
	oneMB         = 1 * 1024 * 1024
	tenMB         = 10 * oneMB
	hundredMB     = 100 * oneMB
	fiveHundredMB = 500 * oneMB
	oneGB         = 1 * 1024 * 1024 * 1024
	fiveGB        = 5 * oneGB
	tenGB         = 10 * oneGB
)

func parseTotalFlag(total string) int {
	switch total {
	case "100mb":
		return hundredMB
	case "500mb":
		return fiveHundredMB
	case "1gb":
		return oneGB
	case "5gb":
		return fiveGB
	case "10gb":
		return tenGB
	default:
		return tenMB
	}
}
