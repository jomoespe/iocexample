package main

import (
	"github.com/jomoespe/iocexample/receiver"
	"github.com/jomoespe/iocexample/sender"
)

// Bind all the "context"
// dependency injection or service locator must be used to set the receiver implementation in sender
var theSender = &sender.Sender{
	Receiver: &receiver.SpecificReceiver{},
}

func main() {
	theSender.DoSomethig()
}
