
package main

import ("fmt"; "errors"; "os"; "strconv")


/*
  FUNCTION numberOfAlerts

  DESCRIPTION:

    A compliance system monitors incoming and outbound calls.

    It sends an alert whenever the average calls exceed a threshold over a
    trailing number of minutes.

    If the trailing minutes to consider is 5 at time T, average the call volumes for
    times T-(5-1), T-(5-2) ...T.


  RETURNS:

    integer that represents the number of alerts sent over the timeframe.


  PARAMETERS:

    precedingMinutes integer:
      trailing number of minutes to consider.

    alertThreshold integer:
      maximum number of calls allowed without triggering an alert.

    numCalls[numCalls[0]..numCalls[n-1]] integer array:
      numCalls[i]: number of calls per minute; i is minutes index


  CONSTRAINTS:

      1 ≤ n                ≤ 105
      1 ≤ precedingMinutes ≤ N
      1 ≤ alertThreshold   ≤ 105
      0 ≤ numCalls[i]      ≤ 105


  EXAMPLE:

     n = 8
     numCalls         = [2, 2, 2, 2, 5, 5, 5, 8] 
     alertThreshold   =  4 
     precedingMinutes =  3 
 
     No alerts will be sent until at least T = 3 because there are not enough
     values to consider. 

     When T = 3; the average calls = (2 + 2 + 2)/3 = 2.

     Additionally, average calls from T = 3 to T = 8 are 2, 2, 3, 4, 5, and 6.

     A total of two alerts are sent during the last two periods.

     Given the data as described, determine the number of alerts sent at the end
     of the timeframe. 
*/

// CONSTRAINT CONSTs from comment above
const Nmax      = 105 // N MAX precedingMinutes alertThreshold numCalls[i]
const Nmin      =   1 // N MIN precedingMinutes alertThreshold
const NZro      =   1 // N MIN                                 numCalls[i]

const SizAry    =   8 // Size of Array of calls per minute from example above

var   aNumCalls = [SizAry] int{2, 2, 2, 2, 5, 5, 5, 8}  // from example above


func
numberOfAlerts(preceedingMinutes int, alertThreshold int, numCalls int) (int, error){

	if   1 > preceedingMinutes {
		return -1, errors.New("preceedingMinutes less than Nmin!")
  }

	if 105 < preceedingMinutes {
		return -2, errors.New("preceedingMinutes greater than Nmax!")
	}

	if   1 > alertThreshold {
		return -3, errors.New("alertThreshold less than Nmin!")
  }

	if 105 < alertThreshold {
		return -4, errors.New("alertThreshold greater than Nmax!")
  }

	if   0 > numCalls {
		return -5, errors.New("numCalls less than Nzro!!")
	}

	if 105 < numCalls {
		return -6, errors.New("numCalls greater than Nmax!")
	}

	// the accumulator for the number of alerts generated for the call dataset
	// and the sum of the given slice of the NumCalls array
	var numAlert, sumSlc int

	// In a real world scenario these probably should be done with gonum
	// https://www.gonum.org
	// but i would have to install it and do other things that would likely take
	// longer than just doing it the old fashioned slow way for this exerise.
  //
	// TD: ponder how this works with a massive array.
  //
	// IMPLEMENTATION
	//
  // When T = 3; the average calls = (2 + 2 + 2)/3 = 2.
  // Additionally, average calls from T = 3 to T = 8 are 2, 2, 3, 4, 5, and 6.
  //
	// using slice and range for this cause it's interesting!
	// 0. interate thru aNumCalls array and select slices
	// 1. add up the number of calls in each slice; accumulate in SumSlc
	// 2. divide SumSlc by preceeding minutes - cast to floats to not drop info.
	// 3. check result of 2 against alertThreshold
	// 4. increment numAlert if result of step 3 is greater than alertThreshold
	// 5. reset SumSlc; continue iteration to completion of for loop
	// 6. return numAlert to the caller

	for i := preceedingMinutes; i < SizAry; i++ {
		for _, cntCall := range aNumCalls[i-preceedingMinutes:i]{
			sumSlc += cntCall
		}
		if (float64(alertThreshold)) <= (float64(sumSlc))/(float64(preceedingMinutes)){
			numAlert+=1
		}
		sumSlc = 0
	}

  return numAlert, nil
}


func
main(){

	// error variables to use down in the code.
	var eAlrt, eAtoi         error

	var iCntAlert            int

	// a value that might be altered on the command line using os.Args
	var argPreceedingMinutes int = 3
	
	// TODO, these could be passed in as arguments too, thus the name, but that's
	// not an interesting thing to do ATM

	var argAlertThreshold    int = 4             // from the example above
	var argNumCalls          int = 8             // from the example above
	
	// print out the default value as an aid to the user so that if they passed in
	// an argument then they can see that it has changed.
	
	fmt.Printf("argPreceedingMinutes=%d\n", argPreceedingMinutes)

  if len(os.Args) == 2 {
		
		// command line argument is optional; only read it if something is there
		// only use the value if there is no error from reading the argument.

		if argPreceedingMinutes, eAtoi = strconv.Atoi(os.Args[1]); eAtoi == nil {
			fmt.Printf("command line argument passed in argPreceedingMinutes=%d\n",
				         argPreceedingMinutes)
		}
	}

	// run the function of interest
	
	if iCntAlert, eAlrt = numberOfAlerts(argPreceedingMinutes, argAlertThreshold,
		                                   argNumCalls); eAlrt != nil {
		fmt.Printf("ERR:\n numberOfAlerts():\n %s Bad thing happened\n", eAlrt)
		return
	}

	fmt.Printf("Number of Alerts: %d\n", iCntAlert)
}

//		fmt.Printf("ERR:\n strconv.Atoi(os.Args[1])\n %s\n", eAtoi,
//			argPreceedingMinutes)
