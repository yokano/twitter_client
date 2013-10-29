// セッション管理用
package model

import (
	"appengine"
)

type Model struct {
	c appengine.Context
}

// モデルのインスタンスを作成する。
// モデルを使う前にかならず実行すること。
func NewModel(c appengine.Context) *Model {
	model := new(Model)
	model.c = c
	return model
}
