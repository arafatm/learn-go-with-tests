package learngowithtests

// Repeat a string 5 times
func Repeat(input string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + input
	}
	return repeated
}
