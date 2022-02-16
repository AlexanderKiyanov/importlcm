package helpers

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func StartImport(utility, importXml, logFile string) {

	//utility = fmt.Sprintf("%q", utility)
	importXml = fmt.Sprintf("%q", importXml)
	logFile = fmt.Sprintf("%q", logFile)

	args := fmt.Sprintf("%s %s >> %s", utility, importXml, logFile)
	cmd := exec.Command("powershell.exe", "/C", args)

	fmt.Printf("Execute command: %s\n\n", args)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error: cmd execution failed with %s and output is:\n %s\n", err, output)
		time.Sleep(14 * time.Second)
		log.Fatalf("error: cmd execution failed with %s and output is:\n %s\n", err, output)
	}

	fmt.Println("Import was successfully executed.")

}
