package cmd

const (
	oneMB                = 1 * 1024 * 1024
	twoHundredAndFiftyMB = 250 * oneMB
	fiveHundredMB        = 500 * oneMB
	oneGB                = 1 * 1024 * 1024 * 1024
	twoGB                = 2 * oneGB
	fourGB               = 4 * oneGB
)

func parseTotalFlag(total string) int {
	switch total {
	case "250mb":
		return twoHundredAndFiftyMB
	case "500mb":
		return fiveHundredMB
	case "1gb":
		return oneGB
	case "2gb":
		return twoGB
	case "4gb":
		return fourGB
	default:
		return twoHundredAndFiftyMB
	}
}
