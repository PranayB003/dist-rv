package main

// NewTimePoint updates the monitor's state when a new time point is identified.
func NewTimePoint(phi *Node, timePoint float64) {
	// Create new nodes for the split interval [0, timePoint), {timePoint}, and (timePoint, ∞)
	node1 := &Node{Formula: phi.Formula, 
                Interval: Interval{Start: 0, End: timePoint}, 
                TruthValue: _unknown, 
                Guards: []*Guard{}, 
                Triggers: []*Trigger{}}
	node2 := &Node{Formula: phi.Formula, 
                Interval: Interval{Start: timePoint, End: timePoint}, 
                TruthValue: _unknown, 
                Guards: []*Guard{}, 
                Triggers: []*Trigger{}}
	node3 := &Node{Formula: phi.Formula,
                Interval: Interval{Start: timePoint, End: math.Inf(1)}, 
                TruthValue: _unknown, 
                Guards: []*Guard{}, 
                Triggers: []*Trigger{}}

	// Update guards and triggers of the new nodes based on the original node (phi, J)
	for _, guard := range phi.Guards {
		precondition := guard.Precondition
		anchorNode := guard.AnchorNode
		continuation := guard.Continuation

		// Update guards for node1, node2, and node3
		newGuard1 := &Guard{Precondition: precondition, 
                        AnchorNode: anchorNode, 
                        Continuation: continuation}
		newGuard2 := &Guard{Precondition: precondition,
                        AnchorNode: anchorNode, 
                        Continuation: continuation}
		newGuard3 := &Guard{Precondition: precondition, 
                        AnchorNode: anchorNode, 
                        Continuation: continuation}

		// Create triggers for the new guards
		trigger1 := &Trigger{FromGuard: newGuard1, ToNode: node2}
		trigger2 := &Trigger{FromGuard: newGuard2, ToNode: node3}

		// Add the new guards and triggers to the corresponding nodes
		node1.Guards = append(node1.Guards, newGuard1)
		node1.Triggers = append(node1.Triggers, trigger1)

		node2.Guards = append(node2.Guards, newGuard2)
		node3.Triggers = append(node3.Triggers, trigger2)
	}

  // TODO: Delete the original node (phi, J)
	// Note: This assumes that the graph structure is maintained externally, and nodes can be deleted safely.

  // TODO: Add the new nodes to the graph structure
	// Note: This also assumes that the graph structure is maintained externally.

  // TODO: Perform any additional actions or updates as needed
}

// SetTruthValue sets the truth value for a node and propagates it through the triggers.
func SetTruthValue(node *Node, value string) {
	node.TruthValue = value

	// Propagate truth value through triggers
	for _, trigger := range node.Triggers {
		fromGuard := trigger.FromGuard
		toNode := trigger.ToNode

		// Update truth value of the successor node based on the trigger
		if value == _unknown {
			SetTruthValue(toNode, _unknown)
		} else if len(toNode.Guards) == 0 {
			SetTruthValue(toNode, _unknown)
		} else {
			SetTruthValue(toNode, _true)
		}
	}
}

// NoTimePoint updates the monitor's state when a nonsingular interval J is complete.
func NoTimePoint(phi *Node, J Interval) {
	// Check if interval J is complete
	isComplete := true
	for _, component := range NewCompleteIntervals(nil) {
		if J.Start <= component.Start && J.End >= component.End {
			isComplete = false
			break
		}
	}

	// Perform actions based on completeness
	if isComplete {
    // TODO: Delete nodes associated with interval J
		// Note: This assumes that the graph structure is maintained externally.

		// Update triggers if necessary
		for _, guard := range phi.Guards {
			if guard.AnchorNode.Interval == J {
        // TODO: Update triggers based on the completion of interval J
				// Note: This assumes that the graph structure is maintained externally.
			}
		}

    // TODO: Call SetTruthValue for relevant nodes to propagate truth values
		// Note: This assumes that the graph structure is maintained externally.
		SetTruthValue(phi, _false)
	}
}

