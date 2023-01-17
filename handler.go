//csinálj egy hello world ot válaszoló http handlert _/
//a válasz body tartalma legyen hello world,
//a válasz http kód legyen http I'm a tea pot (418)

package vpostit

import (
	"net/http"
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
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

	}
	if req.Method == "GET" {
		_, err := resp.Write([]byte("foo\n"))
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		resp.WriteHeader(http.StatusOK)
	}

}
