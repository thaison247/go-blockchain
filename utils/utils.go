package utils

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

const (
	EXPLORE_PAGE = "explorer.html"
)

var (
	ARR_TEMPLATES = []string{
		EXPLORE_PAGE,
	}
)