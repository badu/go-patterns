This is example of asynchronous calculation network ([asynchronous network pattern](https://medium.com/@vague.capitan/async-network-pattern-ae3202ff9462)) powered by golang goroutines.<br/>
We have initial data as ten integers from 0 to 9. We need calculate for each of them sum of it and all previous integers and then find factorial of resulting number. Also we want log results. Construct computational network. The first module simulate initial data and produce numbers from 0 to 9. Through the channel it sends numbers to calculation layer which finds previous sum for received integer.  It consists of ten modules, one for each initial number. Thus every module takes its own number from channel and calculations run in paralel. All of them send results by channel to second layer. Second calculation layer finds factorials. It also consists of ten modules, one for each summation result. Also there is logger module. It gets results from both layers and prints it. The scheme is depicted on figure.<br/>
<br/>
<br/>
<img src="https://github.com/geneva-lake/async-network/blob/master/network.png"/>
<br/>
<br/>
Initial element is realized as Numbers struct

```
type Numbers struct {
    Out chan float64
}

func (n Numbers) Start() {
    for _, i := range n.DoWork() {
        n.Out <- i
    }
}
```

DoWork function produce slice of ten integers. 

Modules in first layer are realized as Sum struct. It has incoming channel for receiving data, out coming channel for result and functions. Function DoWork find sum of previous integers and function Start gets data from channel, calls DoWork and send result to next computation layer and string to logger. 

```
type Sum struct {
    In chan float64
    Out chan float64
    ToLog chan string
}

func (s Sum) Start() {
    for nmbr := range s.In {
        sum := s.DoWork(nmbr)
        s.Out <- sum
        s.ToLog <- fmt.Sprintf("Sum number: %s", 
            strconv.FormatFloat(sum, 'f', 0, 64))
    }
}
```

Second calculation level organized similarly. 

```
type Factorial struct {
    In chan float64
    ToLog chan string
}

func (f Factorial) Start() {
    for sum := range f.In {
        fctrl := f.DoWork(sum)
        f.ToLog <- fmt.Sprintf("Factorial of %s is: %s", 
            strconv.FormatFloat(sum, 'f', 0, 64), strconv.FormatFloat(fctrl, 'e', 0, 64))
    }
}
```


So we need make channels and create modules

```
sumChn := make(chan float64)
fctrlChn := make(chan float64)
lgChn := make(chan string, 20)
nmbrs := Numbers{sumChn}

sums := make([]*Sum, 0, 0)
for i := 0; i < count; i++ {
    sum := &Sum{
        In:    sumChn,
        Out:   fctrlChn,
        ToLog: lgChn,
    }
    sums = append(sums, sum)
}

fctrls := make([]*Factorial, 0, 0)
for i := 0; i < 10; i++ {
    fctrl := &Factorial{
        In: fctrlChn,
        ToLog: lgChn,
    }
    fctrls = append(fctrls, fctrl)
}

lgr := Logger{
    In: lgChn,
    Out: endChan,
}
```


And run Start functions in goroutines from end to beginning

```
go lgr.Start()
for _, fctrl := range fctrls {
    go fctrl.Start()
}
for _, sum := range sums {
    go sum.Start()
}
go nmbrs.Start()
```