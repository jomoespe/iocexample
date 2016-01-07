package receiver

import (
    "fmt"
)

type SpecificReceiver struct {}

func (r *SpecificReceiver) ReceiveThis() {
    // do something interesting
    fmt.Println("ReceiveThis() in SpecificReceiver")
}
