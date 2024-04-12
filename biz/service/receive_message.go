package service

type ReceiveMessage struct {
	BeiDouCard
	Location
}

func (r ReceiveMessage) ProcessingMessage(originatorCardID int, destinationCardID int, timestamp int64, messageContent string) error {
	//TODO implement me
	return nil
}
