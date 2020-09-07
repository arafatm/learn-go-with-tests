package learngowithtests

const repeatCount = 5

// Repeat a string 5 times
func Repeat(input string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += input
	}
	return repeated
}
