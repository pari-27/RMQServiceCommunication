package service

import "RMQServiceCommunication/db"

type Dependencies struct {
	Store db.Storer
	// define other service dependencies
}
