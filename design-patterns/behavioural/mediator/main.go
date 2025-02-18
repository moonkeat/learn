package main

import "fmt"

// Usage:
// Hard to modify classes because they are tightly coupled
// GUI, Chat applications, Control System, System integration

// Pros:
// Decoupling classes
// Centralized control

// Cons:
// Complexity of mediator
// Single point of failure
// Debugging through mediator is hard

func main() {
	tower := &Tower{canLand: true}

	plane1 := Airplane{id: "123", tower: tower}
	plane2 := Airplane{id: "456", tower: tower}

	plane1.RequestLanding()
	plane2.RequestLanding()

	plane1.NotifyTakeOff()
	plane2.RequestLanding()
}

type Tower struct {
	canLand bool
}

type Airplane struct {
	id    string
	tower *Tower
}

func (a *Airplane) RequestLanding() {
	if a.tower.canLand {
		fmt.Printf("Plane %s is landing\n", a.id)
		a.tower.canLand = false
	} else {
		fmt.Printf("Plane %s is waiting for landing\n", a.id)
	}
}

func (a *Airplane) NotifyTakeOff() {
	fmt.Printf("Plane %s is taking off\n", a.id)
	a.tower.canLand = true
}
