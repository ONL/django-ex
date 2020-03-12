package main

import (
	"net/http"
)

type AfrikaContent struct {
  Score int
  Cat1 string
  Cat2 string
  Cat3 string
  Cat4 string
  Cat5 string
  Cat6 string
  Isauthenticated string
}
  

func afrikaKlimaHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
  content := &AfrikaContent{
    Score: 0,
    Cat1: "none",
    Cat2: "none",
    Cat3: "none",
    Cat4: "none",
    Cat5: "none",
    Cat6: "none",
    Isauthenticated: "false"}
  
  if http.MethodGet == r.Method {
        content.Score = 5
        return renderTemplate(w, "afrika_klima", "base", content)
  } else {
	r.ParseForm()
        if "1" == r.PostForm.Get("afgnr") {
            if "posted" == r.PostForm.Get("status") {
                if "subtrop" == r.PostForm.Get("cat1") {
                    content.Score = content.Score + 1
                }
                if "passat" == r.PostForm.Get("cat2") {
                    content.Score = content.Score + 1
                }
                if "wechsel" == r.PostForm.Get("cat3") {
                    content.Score = content.Score + 1
                }
                if "aequatorial" == r.PostForm.Get("cat4") {
                    content.Score = content.Score + 1
                }
                content.Cat1 = r.PostForm.Get("cat1")
                content.Cat2 = r.PostForm.Get("cat2")
                content.Cat3 = r.PostForm.Get("cat3")
                content.Cat4 = r.PostForm.Get("cat4")
                return renderTemplate(w, "afrika_klima", "base", content)
            } else {
                content.Score = 5
                return renderTemplate(w, "afrika_klima", "base", content)
            }
        } else {
            content.Score = 5
            return renderTemplate(w, "afrika_klima", "base", content)
        }
    }
}


func afrikaVegetationHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
  content := &AfrikaContent{
    Score: 0,
    Cat1: "none",
    Cat2: "none",
    Cat3: "none",
    Cat4: "none",
    Cat5: "none",
    Cat6: "none",
    Isauthenticated: "false" }
  
  if http.MethodGet == r.Method {
	content.Score = 7
        return renderTemplate(w, "afrika_vegetation", "base", content)
  } else {
	r.ParseForm()
        if "1" == r.PostForm.Get("afgnr") {
            if "posted" == r.PostForm.Get("status") {
                if "hartlaub" == r.PostForm.Get("cat1") {
                    content.Score = content.Score + 1
                }
                if "wuste" == r.PostForm.Get("cat2") {
                    content.Score = content.Score + 1
                }
                if "dornenstrauch" == r.PostForm.Get("cat3") {
                    content.Score = content.Score + 1
                }
                if "trockensavanne" == r.PostForm.Get("cat4") {
                    content.Score = content.Score + 1
                }
                if "feuchtsavanne" == r.PostForm.Get("cat5") {
                    content.Score = content.Score + 1
                }
                if "tropR" == r.PostForm.Get("cat6") {
                    content.Score = content.Score + 1
                }
                content.Cat1 = r.PostForm.Get("cat1")
                content.Cat2 = r.PostForm.Get("cat2")
                content.Cat3 = r.PostForm.Get("cat3")
                content.Cat4 = r.PostForm.Get("cat4")
                content.Cat5 = r.PostForm.Get("cat5")
                content.Cat6 = r.PostForm.Get("cat6")
                return renderTemplate(w, "afrika_vegetation", "base", content)
	    } else {
                content.Score = 7
                return renderTemplate(w, "afrika_vegetation", "base", content)
	    }
        } else {
            content.Score = 7
            return renderTemplate(w, "afrika_vegetation", "base", content)
       }
   }
}
