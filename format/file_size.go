package format

import "fmt"

const bInKb = 1_000
const bInMb = 1_000_000
const bInGb = 1_000_000_000

func FileSize(size int64) string {
	// gb
	if size >= bInGb {
		return fmt.Sprintf("%dG", size/bInGb)
	}
	// mb
	if size >= bInMb {
		return fmt.Sprintf("%dM", size/bInMb)
	}
	// kb
	if size >= bInKb {
		return fmt.Sprintf("%dk", size/bInKb)
	}
	// b
	return fmt.Sprintf("%d", size)
}
