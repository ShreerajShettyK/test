package main

import "log"

type medium interface {
	sendMessage(userId string) error
}

type mediumImpl struct {
	senderID   string
	senderName string
}

func (ms *mediumImpl) sendMessage(userId string) error {
	ms.senderID = "12345"
	ms.senderName = "Relanto"
	log.Println(ms.senderID)
	log.Println(ms.senderName)
	return nil
}

func main() {
	var m medium = &mediumImpl{}
	userId := "user_001"
	m.sendMessage(userId)
}
