package tui

import (
	"charm.land/lipgloss/v2"
)

// table styling 
var (
	baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#424549")).
	Padding(0, 1).
	Margin(1, 2)
)

// title style
var titleStyle = 
	lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#64ff71"))

// search style
var searchStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#507fcb")).
	Bold(true).
	Faint(true)

// delete style
var delStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#cd3b3b")). 
	Bold(true).
	Faint(true)

// footer style
var footStyle = 
	lipgloss.NewStyle().
	Foreground(lipgloss.Color("#34383e")).
	Underline(true).
	Faint(true)
   
// selection style
var selectStyle = 
	lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffff")).
	Background(lipgloss.Color("#282b30")).
	Bold(true)
