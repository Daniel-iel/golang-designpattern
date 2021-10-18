package main

import "fmt"

func main() {
	var bldr = newNotificationBuilder().
		SetTitle("New Notification").
		SetIcon("Icon.png").
		SetImage("image.png").
		SetPriority(10).
		SetMessage("This is a basic Notifiction with some text in it").
		SetType("Alert")

	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creatinh the notification")
	}

	fmt.Printf("Notification: %+v", notif)
}
