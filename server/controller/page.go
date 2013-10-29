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
	
	if this.GetSession(c, r) == "" {
		view.Login()
	} else {
		view.TimeLine()
	}
}
