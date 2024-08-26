package main

const ( // actions
	// Top action
	ActionTop = iota

	//Bot action
	ActionBot

	//Left action
	ActionLeft

	// Right action
	ActionRight
)

// Action Struct link enum action to name : string
type Action struct {
	Name  string
	Value int
}

// None for no action -> first turn
var ActionNone = Action{
	"None",
	-1,
}

// L array of actions
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
