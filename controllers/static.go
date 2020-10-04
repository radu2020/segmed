package controllers

import "../views"

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
		ShowTable: views.NewView("bootstrap", "views/static/showtable.gohtml"),
	}
}

type Static struct {
	Home  *views.View
	Contact *views.View
	ShowTable *views.View
}

