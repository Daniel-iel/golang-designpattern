package main

type television interface {
	volumeUp() int
	volumeDowns() int
	channelUp() int
	channelDown() int
	turnOn()
	turnOff()
	goToChannel(ch int)
}
