package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"os/exec"
)

func clear_myScreen() {
	out, _ := exec.Command("/usr/local/bin/clear").Output()
  os.Stdout.Write(out)
}

func get_myTime()  {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))	
}

func get_wrkDir() {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The current working directory is", cwd)
	}
}

func list_testDir() {
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}

func main() {
	
	
	// Create new parser object
	parser := argparse.NewParser("test", "Simple example of argparse flags")

	// Create input flag
	
	commandScan := parser.NewCommand("scan", "Returns random numbers")
	
	verbose := parser.Flag(
		"", "verbose", &argparse.Options{
			Help: "Verbose mode",
		},
	)

	scannerInput := commandScan.String(
		"i", "input", &argparse.Options{
			Required: true,
			Help: "I need Input file to scan against",
		},
	)

	scannerOutput := commandScan.String(
		"o", "", &argparse.Options{
			Help: "Please give me an output file",
		},
	)


	//obj_input := parser.NewCommand("i", "input", &argparse.Options{Help: "input"})
	//obj_output := parser.NewCommand("o", "output", &argparse.Options{Help: "output"})

	/*obj_input := commandinput.String(
		"i", "input", &argparse.Options{
			Required: true,
			Help: "Please enter input file to process",
		},
	)
	obj_output := commandoutput.String(
		"o", "output", &argparse.Options{
			Required: true,
			Help: "Please enter output file",
		},
	)*/


	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	switch {
	case commandScan.Happened():
		fmt.Println("scanner:", commandScan.Happened())
		fmt.Println("--input:", *scannerInput)
		fmt.Println("--output:", *scannerOutput)
		fmt.Println("--verbose:", *verbose)
	}

		/*
	// if verbose mode enabled, print debugging information
	if *obj_input {
		fmt.Println("Args: ", os.Args)
		clear_myScreen()
		list_testDir()
	}
	if *obj_output {
		fmt.Println("Output: ", os.Args)
	}
	get_wrkDir()
	get_myTime()*/
	
}