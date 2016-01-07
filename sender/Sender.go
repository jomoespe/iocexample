package sender

type Sender struct {
	Receiver Receiver
}

func (s *Sender) DoSomethig() {
    s.Receiver.ReceiveThis()
}
