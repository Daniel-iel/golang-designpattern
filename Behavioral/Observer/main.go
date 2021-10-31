package main

func main() {
	listner1 := DataListener{
		Name: "listner1",
	}

	listner2 := DataListener{
		Name: "listner2",
	}

	subj := &DataSubject{}

	subj.registerObserver(listner1)
	subj.registerObserver(listner2)

	subj.ChangeItem("Monday!")
	subj.ChangeItem("Wednesday!")
	subj.ChangeItem("Friday!")

	subj.unregisterObserver(listner1)
	subj.ChangeItem("Friday!")
}
