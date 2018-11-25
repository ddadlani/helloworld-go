package target

import "os"

func SetTarget(target string) {
	os.Setenv("TARGET", target)
}
