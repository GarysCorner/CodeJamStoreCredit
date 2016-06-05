//File:		codejamstorecredit.go
//Author:	Gary Bezet
//Desc:		This program is designed to solve Google Code Jam's qualification round for Africa 2010.  I did not compete in the codejam but I was having trouble figuring out what to code so I decided to give it a shot a test problem.  Also I'm learning Go.
//Problem:	https://code.google.com/codejam/contest/351101/dashboard#s=p0

package main

import (
		"fmt"
		"flag"
		"os"
		"bufio"
		"strings"
		"strconv"
		"time"
		"sort"
	)

//global variables
var infileopt, outfileopt string
var infile, outfile *os.File

var n int //number of cases
var testcases []testcase

var starttime time.Time = time.Now()





//structures
type testcase struct {
	num, credit, items int
	priceshigh, priceslow []int
	
	
}


//main entry point
func main() {

	
	
	initflags()
	

	openFiles()

	defer infile.Close()
	defer outfile.Close()

	defer printStats() //must run before close of outfile in case outfile is stdout

	printProgramOptions()
	
	printCases() //this is just for testing

	

}


//print the stats after progra ends
func printStats() {
	printErrln( "\nTotal time:  ", time.Since( starttime ))
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

	inputFile( bufio.NewReader(infile) )  //process the infile contents

	
}



//print error to console
func printErrln( line ...interface{} ) {
	fmt.Fprintln( os.Stderr, line... )
}



//print a solution line either to stdout or file
func printSolution( line ...interface{} ) {
	fmt.Fprintln(outfile, line... )

}


//process input file
func inputFile(reader *bufio.Reader) {

	
	var err error

	

	//get N or number of cases
	var nline string
	nline, err = reader.ReadString('\n')
	if err != nil {
		printErrln("Couldn't read first line from:  ", infileopt)
		printErrln("\tError:  ", err )
		os.Exit(4)
		
	}

	n, err = strconv.Atoi( strings.TrimSpace( nline  ) )
	if err != nil  { //if error reading number of cases
		printErrln("Couldn't read case numbers from:  ", infileopt)
		printErrln("\tError:  ", err )
		os.Exit(5)
	}


	//create testcase array
	testcases = make([]testcase, n)

	var casenum = 0;

	for i := 0; i < n*3; i++   {

		var line string
		

		line, err = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		
		switch i % 3 {

			case 0:  //process credit line and casenum
				testcases[casenum].num = casenum + 1
				testcases[casenum].credit, err = strconv.Atoi( line )  				

			case 1: //process itemnum line
				testcases[casenum].items, err = strconv.Atoi(line)

			case 2:  //process items price line and increment casenum

				itemstring := strings.Split(line, " ")

				if len(itemstring) != testcases[casenum].items {  //sanity check
					printErrln("Error:  items #s didn't match for case# ", casenum)
					os.Exit(6)
				}
				testcases[casenum].priceshigh = make([]int, 0, len(itemstring) )
				testcases[casenum].priceslow = make([]int, 0, len(itemstring) )

				halfwayexact := false //make sure halway values get in piceslow then priceshigh

				for itemnum := 0; itemnum < len(itemstring); itemnum++ {
					
					price, err := strconv.Atoi(itemstring[itemnum])
					
					if err != nil {
						printErrln("Error:  Could not convert item#", itemnum, " case#", testcases[casenum].num, " to int")
						printErrln( "\tError: ", err)
						os.Exit(7)
					}
					
					halfway := testcases[casenum].credit / 2
					
					
					switch {
						
						case price > testcases[casenum].credit:  //ignore price if too high
							
							printErrln("Case#", casenum+1, " item#", itemnum+1, " price too high!")  //for testing
						
						case price > halfway:
							testcases[casenum].priceshigh = append( testcases[casenum].priceshigh, price )
							
						case price < halfway:
							testcases[casenum].priceslow = append( testcases[casenum].priceslow, price )
					
						
						case price == halfway:  //special case may be better way, put one in low and the rest in high
							
							if halfwayexact == true {
								testcases[casenum].priceshigh = append( testcases[casenum].priceshigh, price )
							} else {
								testcases[casenum].priceslow = append( testcases[casenum].priceslow, price )
								halfwayexact = true
							} 
							
					}
					
				
				}
				
				sort.Ints(testcases[casenum].priceshigh )
				sort.Ints(testcases[casenum].priceslow )

				casenum++
				

		}

		if( err != nil ) {
			printErrln( "Failed to process infile:  ", infileopt)
			printErrln("\tError:  ", err)
			os.Exit(5)
		}
		

	}

	

}

//testprint data structure
func printCases() {

	for _, v := range testcases {
		printErrln("Case#:\t" , v.num)
		printErrln("Credit:\t", v.credit)
		printErrln("# of Items\t", v.items)
		printErrln("PriceLow:\t", v.priceslow)
		printErrln("PriceHigh:\t", v.priceshigh, "\n")	
	}	

}


