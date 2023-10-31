package main

import (
	"fmt"
	"math"
	"time"
)

// Global variable to store verdicts.
var verdicts []Verdict

// NewTimePoint updates the monitor's state when a new time point is identified.
func NewTimePoint(phi *Node, timePoint float64) {
}

// SetTruthValue sets the truth value for a node and propagates it through triggers.
func SetTruthValue(node *Node, value string) {
}

// NoTimePoint updates the monitor's state when a nonsingular interval J is complete.
func NoTimePoint(phi *Node, J Interval) {
}

// NewCompleteIntervals returns new complete intervals based on the received message.
func NewCompleteIntervals(msg interface{}) []Interval {
	return nil
}

// Monitor is the main loop of the monitoring algorithm.
func Monitor(phi *Node, msg interface{}) {
	verdicts = []Verdict{} // Reset verdicts for each message

	switch m := msg.(type) {
	case NotifyMessage:
		NewTimePoint(phi, m.Timestamp)
	case ReportMessage:
		NewTimePoint(phi, m.Timestamp)
		SetTruthValue(m.Proposition, m.TruthValue)
		for _, interval := range NewCompleteIntervals(m) {
			NoTimePoint(phi, interval)
		}
    // TODO: Add additional handling if needed
	}

  // TODO: Return verdicts computed in this iteration
	fmt.Println(verdicts)
}

func main() {
	// test usage
	phi := &Node{
		Formula:    "ExampleFormula",
		Interval:   Interval{Start: 0, End: math.Inf(1)},
		TruthValue: _unknown,
		Guards:     []*Guard{},
		Triggers:   []*Trigger{},
	}

	msg1 := NotifyMessage{Component: "C", Timestamp: 2.0, Sequence: 2}
	msg2 := ReportMessage{Proposition: "p", TruthValue: _false, Timestamp: 2.0}

	Monitor(phi, msg1)
	Monitor(phi, msg2)
	// more messages after this
}
