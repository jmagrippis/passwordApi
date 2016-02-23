package main_test

import (
	"github.com/jmagrippis/passwordApi"

	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"

	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerator(t *testing.T) {

	Convey("Given I visit \"/generate/5\"", t, func() {

		req, err := http.NewRequest("GET", "/generate/5", nil)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()
		params := httprouter.Params{httprouter.Param{Key: "amount", Value: "5"}}

		Convey("It returns a json array of 5 passwords", func() {
			main.Generate(w, req, params)
			So(w.Code, ShouldEqual, 200)
			var response []string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			So(len(response), ShouldEqual, 5)
		})

	})

	Convey("Given I visit \"/generate/12\"", t, func() {

		req, err := http.NewRequest("GET", "/generate/12", nil)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()
		params := httprouter.Params{httprouter.Param{Key: "amount", Value: "12"}}

		Convey("It returns a json array of 12 passwords", func() {
			main.Generate(w, req, params)
			So(w.Code, ShouldEqual, 200)
			var response []string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			So(len(response), ShouldEqual, 12)
		})

	})

	Convey("Given I visit \"/generate/1?delimiter=x", t, func() {

		Convey("I get a json array of one password delimited by \"x\"", nil)

	})

	Convey("Given I visit \"/generate/1?prefix=x", t, func() {

		Convey("I get a json array of one password prefixed by \"x\"", nil)

	})

	Convey("Given I visit \"/generate/1?suffix=x", t, func() {

		Convey("I get a json array of one password suffixed by \"x\"", nil)

	})

	Convey("Given I visit \"/generate/1?titleCase=true", t, func() {

		Convey("I get a json array of one password In Title Case", nil)

	})

	Convey("Given I visit \"/generate/1/safe", t, func() {

		Convey("I get a json array of one password that will satisfy most site password requirements", nil)

	})

}
