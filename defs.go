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

type DatacoronerG  struct {
   Get goose.Alert
}



var Goose DatacoronerG
