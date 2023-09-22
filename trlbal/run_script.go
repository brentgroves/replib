package trlbal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// https://www.sohamkamani.com/golang/exec-shell-command/
func RunScript(script string) string {
	// fmt.Printf("Path Separator: '%c'\n", os.PathSeparator)

	fmt.Printf("/home/brent/src/reports/volume/go/runner/main\n")
	//  ./test.sh dev reports11 30011 1 reports11 30311 reports
	// /home/brent/src/reports/volume/go/runner/main/test.sh dev reports11 30011 1 reports11 30311 reports
	cmd6 := exec.Command("/home/brent/src/reports/volume/go/runner/main/test.sh", "dev", "reports11", "30011", "1", "reports11", "30311", "reports")
	cmd6.Stdout = os.Stdout

	err6 := cmd6.Run()

	if err6 != nil {
		log.Fatal(err6)
	}

	return script
}
