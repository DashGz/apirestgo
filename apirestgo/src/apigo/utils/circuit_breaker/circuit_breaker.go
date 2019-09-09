package circuit_breaker

import (
	"time"
	"github.com/sparrc/go-ping"
)

type CircuitBreaker struct {
	State            string
	Timeout          int
	FailureThreshold int
	FailureCount     int
	ChState          chan string
}




func (cb *CircuitBreaker) StartCB(state string, timeout int, failure_threshold int) {

	//initial state
	cb.State = "CLOSED"
	//timeout for the API request
	cb.Timeout = timeout
	// Number of failures we receive from the depended service before we change the state to 'OPEN'
	cb.FailureThreshold = failure_threshold
	cb.FailureCount = 0



	go func(ChState chan string) {

		select {
		case currentSate := <- ChState:
			if currentSate == "OPEN" {
				StartTimeout(cb)
			} else if currentSate == "HALF-OPEN" {
				TestConnection(cb)
			} else {
				Reset(cb)
			}

		}
	}(cb.ChState)
}

func Reset (cb *CircuitBreaker) {
	cb.State = "CLOSED"
	cb.FailureCount = 0
}

func (cb *CircuitBreaker) RecordFailure()  {
	cb.FailureCount++
	if cb.FailureCount > cb.FailureThreshold{
		cb.State = "OPEN"
		cb.ChState <- cb.State
	}
}

func StartTimeout(cb *CircuitBreaker)  {
	time.Sleep(time.Second * time.Duration(cb.Timeout))
	cb.State = "HALF-OPEN"
	cb.ChState <- cb.State
}

func TestConnection(cb *CircuitBreaker) {

	var c [3]string

	c[0] = "http://localhost:8082/sites/"
	c[1] = "http://localhost:8082/users/"
	c[2] = "http://localhost:8082/countries/"


	for i := 0; i< 3; i++ {
		if testResponse(c[i]) == false {
			cb.State = "OPEN"
			cb.ChState <- cb.State
			return
		}
	}

	cb.State = "CLOSE"
	cb.ChState <- cb.State
}

func testResponse(url string) (bool){

	response := true
	pinger, err := ping.NewPinger(url)
	if err != nil {
		panic(err)
	}
	pinger.Count = 1
	pinger.Run()
	pinger.OnFinish = func(statistics *ping.Statistics) {
		if  statistics.PacketsSent - statistics.PacketsRecv != 0{
			response = false
		}
	}

	return response
}



