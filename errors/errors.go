package errors

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func SecurityError(err error) {
	fmt.Println(aurora.Red("Stop being naughty. Error: "), err)
}
func NoKeyError() {
	fmt.Println(aurora.Red("No key found in parakeet.toml!"))
}
func UnknownError(err error) {
	fmt.Println(aurora.Red("Error: "), err)
}
