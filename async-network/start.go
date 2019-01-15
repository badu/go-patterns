package main

func start(endChan chan interface{}) {
	// creation channels for communication between modules
	sumChn := make(chan float64)
	fctrlChn := make(chan float64)
	lgChn := make(chan string, 2*count)

	// creation calculation modules
	nmbrs := Numbers{sumChn}

	// creation slice of Sum modules for asynchronous
	// calculation sum of previous integers
	sums := make([]*Sum, 0, 0)
	for i := 0; i < count; i++ {
		sum := &Sum{
			In:    sumChn,
			Out:   fctrlChn,
			ToLog: lgChn,
		}
		sums = append(sums, sum)
	}

	// creation slice of Factorial modules for asynchronous
	// calculation of factorials
	fctrls := make([]*Factorial, 0, 0)
	for i := 0; i < count; i++ {
		fctrl := &Factorial{
			In:    fctrlChn,
			ToLog: lgChn,
		}
		fctrls = append(fctrls, fctrl)
	}

	// creation logger
	lgr := Logger{
		In:  lgChn,
		Out: endChan,
	}

	// start goroutines containing functions for calculations
	go lgr.Start()
	for _, fctrl := range fctrls {
		go fctrl.Start()
	}
	for _, sum := range sums {
		go sum.Start()
	}
	go nmbrs.Start()
}
