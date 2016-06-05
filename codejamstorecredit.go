//File:		codejamstorecredit.go
//Author:	Gary Bezet
//Desc:		This program is designed to solve Google Code Jam's qualification round for Africa 2010.  I did not compete in the codejam but I was having trouble figuring out what to code so I decided to give it a shot a test problem.  Also I'm learning Go.

package main

import (
		"fmt"
		"flag"
		"os"
	)

//global variables
var infileopt, outfileopt string
var infile, outfile *os.File


//main entry point
func main() {

	initflags()
	

	openFiles()

	defer infile.Close()
	defer outfile.Close()

	printProgramOptions()

	printSolution("test ", "fass")

}



//get the flags from command line
func initflags() {
	flag.StringVar(&infileopt, "if", "", "Input file (required)")
	flag.StringVar(&outfileopt, "of", "-", "Output file, defaults to stdout" )

	flag.Parse()

	if infileopt == "" {
		printErrln("You must supply an input file\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	

}



//prints program settings
func printProgramOptions() {
	printErrln("InFile:\t", infileopt)
	printErrln("OutFile:\t", outfileopt)
	

}



//open the in and out files
func openFiles() {
	
	var err error

	infile, err = os.Open(infileopt)

	if err != nil {
		printErrln( "Error:  Could not open:  ", infileopt)
		printErrln( "\tError: ", err  )
		os.Exit(2)
	}

	if outfileopt == "-"  {
		outfile = os.Stderr
		outfileopt = "Stderr"
	} else {
		outfile, err = os.Create(outfileopt)

		if err != nil {
			printErrln( "Error:  Could not create:  ", outfileopt)
			printErrln( "\tError: ", err  )
			os.Exit(3)
		}
	}

	
}



//print error to console
func printErrln( line ...interface{} ) {
	fmt.Fprintln( os.Stderr, line... )
}



//print a solution line either to stdout or file
func printSolution( line ...interface{} ) {
	fmt.Fprintln(outfile, line... )

}

