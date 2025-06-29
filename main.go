package main

import(
	"fmt"
	"net/http"
//	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my great site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w,"<h1>Contact Page</h1><p> To get in touch, email me at <a href=\"mailto:nerds@veryeasy.dev\">nerds@veryeasy.dev</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul>
  <li>
    <b>Is there a free version?</b>
    Yes! We offer a free trial for 30 days on any paid plans.
  </li>
  <li>
    <b>What are your support hours?</b>
    We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
  </li>
  <li>
    <b>How do I contact support?</b>
    Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
  </li>
</ul>
`)
}

func pathHandler(w http.ResponseWriter, r *http.Request){
	//fmt.Fprint(w, r.URL.Path)
	//if r.URL.Path == "/"{
	//	homeHandler(w,r)
	//} else if r.URL.Path == "/contact"{
	//	contactHandler(w,r)
	//}
	switch r.URL.Path{
	case "/":
	  homeHandler(w,r)
	case "/contact":
	  contactHandler(w,r)
	case "/faq":
	  faqHandler(w,r)
	default:
	 // w.WriteHeader(http.StatusNotFound)
	 // fmt.Fprint(w, "Page not found v1")
	  http.Error(w, "Page not found v2", http.StatusNotFound)
	}
}

//type Router struct {}

//func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request){
//        switch r.URL.Path{
//	case "/":
//	  homeHandler(w,r)
//	case "/contact":
//	  contactHandler(w,r)
//	default:
//	  http.Error(w, "Page not found v2", http.StatusNotFound)
//	}
//}
func main(){
	//var router http.HandlerFunc = pathHandler
	//var router Router
	//http.HandleFunc("/", pathHandler)
	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/contact", contactHandler)
	//fmt.Println("Starting the server on :3000...")
	//fmt.Fprintln(os.Stdout, "Open your browser, and go to localhost:3000 to view the page.")
	//http.ListenAndServe(":3000", nil)
	//http.ListenAndServe(":3000", pathHandler)
	//http.ListenAndServe(":3000", router)
	//http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/faq",faqHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
