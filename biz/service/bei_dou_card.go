package service

type BeiDouCard struct {
}

func (b BeiDouCard) CheckStation(originatorID string, station int) error {
	//TODO implement me
	return nil
}

func (b BeiDouCard) ParseBeiDouCard(cardID string) (shipID int, err error) {
	//TODO implement me
	panic("implement me")
}

func (b BeiDouCard) CheckMessageType(cardId string, messageType string) error {
	//TODO implement me
	return nil
}
