package datacoroner

import (
   "golang.org/x/net/html"
)

func GetFields(nodes []*html.Node, needles []HtmlCoronerT) map[string]string {
   var i, j               int
   var flds    map[string]string
   var needle             HtmlCoronerT
   var needlesRemaining []HtmlCoronerT
   var node              *html.Node

   flds = map[string]string{}
   Goose.Get.Logf(6,"Entering Getfields")
   needlesRemaining = make([]HtmlCoronerT,len(needles))
   for i, needle = range needles {
      needlesRemaining[i] = needle
   }

   for i=0; i<len(nodes); i++ {
      Goose.Get.Logf(6,"Getfields i=%d",i)
      if (nodes[i]!=nil) && (nodes[i].FirstChild!=nil) {
         Goose.Get.Logf(6,"Getfields (nodes[i]!=nil) && (nodes[i].FirstChild!=nil)")
         for node=nodes[i].FirstChild; node !=nil; node=node.NextSibling {
            Goose.Get.Logf(6,"Getfields node:%#v",node)
            for j, needle = range needlesRemaining {
               Goose.Get.Logf(6,"Getfields j=%d, needle=%#v",j,needle)
               if needle.Test(node) {
                  Goose.Get.Logf(6,"Getfields needle.Test(node)==true")
                  flds[needle.GetId()] = needle.GetVal(node)
                  Goose.Get.Logf(6,"Got: %d=%s",j,flds[needle.GetId()])
                  if len(needles) == 0 {
                     return flds
                  }
                  Goose.Get.Logf(6,"Needle: %s",needle.GetId())
                  needlesRemaining = append(needlesRemaining[:j],needlesRemaining[j+1:]...)
                  break
               }
            }
         }
      }
   }
   return flds
}

