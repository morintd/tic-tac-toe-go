package game

type IPlayer interface {
	Send(action *GameNotificationAction)
	Receiver() chan PlayerAction
}
