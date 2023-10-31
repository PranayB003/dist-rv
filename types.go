package main

// 3 truth values
var _unknown string = "unknown"
var _true    string = "true"
var _false   string = "false"

// time interval [start, end).
type Interval struct {
	Start float64
	End   float64
}

// graph structure node
type Node struct {
	Formula     string
	Interval    Interval
	TruthValue  string
	Guards      []*Guard
	Triggers    []*Trigger
}

// node guard
type Guard struct {
	Precondition *Node
	AnchorNode   *Node
	Continuation *Node
}

// guard trigger
type Trigger struct {
	FromGuard *Guard
	ToNode    *Node
}

type Verdict struct {
	Value string
	Time  float64
}

type NotifyMessage struct {
	Component string
	Timestamp float64
	Sequence  int
}

type ReportMessage struct {
	Proposition string
	TruthValue  string
	Timestamp   float64
}
