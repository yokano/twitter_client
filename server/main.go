// Twitter Client
package main

import (
	. "server/controller"
)

// エントリポイント
func init() {
	controller := NewController()
	controller.Handle()
}