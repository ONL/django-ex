package main

import (
	"net/http"
)

type AfrikaContent struct {
  score int
  cat1 string
  cat2 string
  cat3 string
  cat4 string
  cat5 string
  cat6 string
  isauthenticated string
}
  

func afrikaKlimaHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
  content := &AfrikaContent{
    score: 0,
    cat1: "none",
    cat2: "none",
    cat3: "none",
    cat4: "none",
    cat5: "none",
    cat6: "none",
    isauthenticated: "false"}
  
  if http.MethodGet == r.Method {
        content.score = 5
        return renderTemplate(w, "afrika_klima", "base", content)
  } else {
        if "1" == r.PostForm.Get("afgnr") {
            if "posted" == r.PostForm.Get("status") {
                if "subtrop" == r.PostForm.Get("cat1") {
                    content.score = content.score + 1
                }
                if "passat" == r.PostForm.Get("cat2") {
                    content.score = content.score + 1
                }
                if "wechsel" == r.PostForm.Get("cat3") {
                    content.score = content.score + 1
                }
                if "aequatorial" == r.PostForm.Get("cat4") {
                    content.score = content.score + 1
                }
                content.cat1 = r.PostForm.Get("cat1")
                content.cat2 = r.PostForm.Get("cat2")
                content.cat3 = r.PostForm.Get("cat3")
                content.cat4 = r.PostForm.Get("cat4")
                return renderTemplate(w, "afrika_klima", "base", content)
            } else {
                content.score = 5
                return renderTemplate(w, "afrika_klima", "base", content)
            }
        } else {
            content.score = 5
            return renderTemplate(w, "afrika_klima", "base", content)
        }
    }
}


func afrikaVegetationHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
  content := &AfrikaContent{
    score: 0,
    cat1: "none",
    cat2: "none",
    cat3: "none",
    cat4: "none",
    cat5: "none",
    cat6: "none",
    isauthenticated: "false" }
  
  if http.MethodGet == r.Method {
	content.score = 7
        return renderTemplate(w, "afrika_vegetation", "base", content)
  } else {
        if "1" == r.PostForm.Get("afgnr") {
            if "posted" == r.PostForm.Get("status") {
                if "hartlaub" == r.PostForm.Get("cat1") {
                    content.score = content.score + 1
                }
                if "wuste" == r.PostForm.Get("cat2") {
                    content.score = content.score + 1
                }
                if "dornenstrauch" == r.PostForm.Get("cat3") {
                    content.score = content.score + 1
                }
                if "trockensavanne" == r.PostForm.Get("cat4") {
                    content.score = content.score + 1
                }
                if "feuchtsavanne" == r.PostForm.Get("cat5") {
                    content.score = content.score + 1
                }
                if "tropR" == r.PostForm.Get("cat6") {
                    content.score = content.score + 1
                }
                content.cat1 = r.PostForm.Get("cat1")
                content.cat2 = r.PostForm.Get("cat2")
                content.cat3 = r.PostForm.Get("cat3")
                content.cat4 = r.PostForm.Get("cat4")
                content.cat5 = r.PostForm.Get("cat5")
                content.cat6 = r.PostForm.Get("cat6")
                return renderTemplate(w, "afrika_vegetation", "base", content)
	    } else {
                content.score = 7
                return renderTemplate(w, "afrika_vegetation", "base", content)
	    }
        } else {
            content.score = 7
            return renderTemplate(w, "afrika_vegetation", "base", content)
       }
   }
}
