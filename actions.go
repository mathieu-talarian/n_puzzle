package main

const ( // actions
	ActionTop = iota

	ActionBot

	ActionLeft

	ActionRight
)

// Action Struct link enum action to name : string
type Action struct {
	Name  string
	Value int
}

var ActionNone = Action{
	"None",
	-1,
}

var ActionsList = [4]Action{
	{
		"Top",
		ActionTop,
	},
	{
		"Bot",
		ActionBot,
	},
	{
		"Left",
		ActionLeft,
	},
	{
		"Right",
		ActionRight,
	},
}
