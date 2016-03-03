package datacoroner

import (
//   "strings"
   "golang.org/x/net/html"
   "github.com/luisfurquim/goose"
)

type HtmlCoronerT interface {
   Test(node *html.Node) bool
   GetId() string
   GetVal(node *html.Node) string
}

var Goose struct {
   Get goose.Alert
}

