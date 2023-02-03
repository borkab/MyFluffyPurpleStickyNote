//csinálj egy hello world ot válaszoló http handlert _/
//a válasz body tartalma legyen hello world,
//a válasz http kód legyen http I'm a tea pot (418)

package vpostit

import (
	"log"
	"net/http"
	"strconv"
)

type FluffyHandler struct { //create a handler struct
}

// implement `ServeHTTP` method on `FluffyHandler` struct
func (f FluffyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//beallitom a header map tartalmat
	res.Header().Set("X-Foo", "bar")
	//megadom h mi legyen a statuscode
	res.WriteHeader(http.StatusTeapot)
	//megadom a response Body tartalmat
	_, _ = res.Write([]byte("Hello World!\n<3"))

}

type BuzzLightyearsLaserHandLER struct {
}

func (b BuzzLightyearsLaserHandLER) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		//resp.Header().Set("foo", "bar")
		resp.WriteHeader(http.StatusCreated)
		_, err := resp.Write([]byte("bar\n"))
		if err != nil {
			log.Println("error", err.Error())
			return
		}

	}
	if req.Method == "GET" {
		_, err := resp.Write([]byte("foo\n"))
		if err != nil {
			log.Println("error", err.Error())
			return
		}
		resp.WriteHeader(http.StatusOK)
	}

}

// a teszt arról szóljon hogy hívod a handlert egy requesttel amiben a X-Foo headernek valamilyen értéket adtál.
// az erre kapott valasz Response Body -ja pontosan ugyan az az érték legyen mint amit a request header X-Foo ban küldtél el
type Mandragora struct{}

func (m Mandragora) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		rw.Header().Set("X=Foo", "Mandragora's scream")
		_, err := rw.Write([]byte(r.Header.Get("X-Foo")))
		if err != nil {
			log.Println("error", err.Error())
			return
		}
		rw.WriteHeader(http.StatusOK)
	}
}

//csinalj egy handlert ami kepes az url query stringekbol kivenni a foo, bar es baz kulcsok ertekeit,
//es a baz eseten kepes akar tobb erteket is kezelni
//POST /my/path?foo=oof&bar=123&baz=hello&baz=world

//-> your server takes the values out from the query and makes the following:
//foo -> “off” (string)
//bar -> 123 (int)
//baz -> []string{“hello”, “world”}

//GET /my/path?my=query
//a request object .URl.Query() alatt éred el a parsolt query értékeket

// /*
type QHandler struct{}

type MyQuerysDTO struct {
	Foo string `json:"Foo"`
	Bar int    `json:"Bar"`
	Baz string `json:"Baz"`
}

func (h QHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	myQ := MyQuerysDTO{}

	myQ.Foo = r.URL.Query().Get("foo") //igy kapom meg a "foo" kulcs ertekparjat a querybol
	myQ.Bar, _ = strconv.Atoi(r.URL.Query().Get("bar"))
	myQ.Baz = r.URL.Query().Get("baz")

}

//*/
