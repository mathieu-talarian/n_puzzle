package main

const ( // actions
	// Top action
	Top = iota

	//Bot action
	Bot

	//Left action
	Left

	// Right action
	Right
)

// Action Struct link enum action to name : string
type Action struct {
	Name  string
	Value int
}

// None for no action -> first turn
var None = Action{
	"None",
	-1,
}

// L array of actions
var L = [4]Action{
	{
		"Top",
		Top,
	},
	{
		"Bot",
		Bot,
	},
	{
		"Left",
		Left,
	},
	{
		"Right",
		Right,
	},
}
