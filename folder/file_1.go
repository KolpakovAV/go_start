package folder

import "fmt"

func GetStr(s string) string {

	return fmt.Sprintf("GET /%s", s)
}
