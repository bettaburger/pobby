package tui

import (
	"charm.land/lipgloss/v2"
)


// title style
var titleStyle = 
	lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#64ff71"))

// search style
var searchStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#507fcb")).
	Bold(true)

// delete style
var delStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#cd3b3b")). 
	Bold(true)

// footer style
var footStyle = 
	lipgloss.NewStyle().
	Foreground(lipgloss.Color("#34383e"))

// selection style
var selectStyle = 
	lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffff")).
	Background(lipgloss.Color("#282b30")).
	Bold(true)
