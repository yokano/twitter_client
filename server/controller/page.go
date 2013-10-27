package controller

import (
	"net/http"
	"appengine"
	. "server/view"
)

// ログインページの表示
func (this *Controller) Top(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	view := NewView(c, w)
	view.Login()
}

// タイムラインの表示
func (this *Controller) TimeLine(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	view := NewView(c, w)
	view.TimeLine()
}