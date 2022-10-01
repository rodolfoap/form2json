package main
import("html/template";)

type Form struct {
	tmpl *template.Template
}

type Instance struct {
	Edit		bool	`json:"-"` // No need to store the mode
	Id		string	`json:"-"` // No need to store the id, it is the filename
	Scope		string
	Type		string
	Definition	string
	Campaign	string
	Notes		string
}

type Scenario struct {
	Edit		bool	`json:"-"` // No need to store the mode
	Id		string	`json:"-"` // No need to store the id, it is the filename
	Constraint	string
}
