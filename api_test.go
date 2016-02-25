package main_test

import (
	"github.com/jmagrippis/passwordApi"

	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"

	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode/utf8"
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

	Convey("Given I visit \"/generate/1?delimiter=^", t, func() {

		req, err := http.NewRequest("GET", "/generate/1?delimiter=^", nil)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()
		params := httprouter.Params{httprouter.Param{Key: "amount", Value: "1"}}

		Convey("I get a json array of one password delimited by \"^\"", func() {
			main.Generate(w, req, params)
			So(w.Code, ShouldEqual, 200)
			var response []string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			So(len(response), ShouldEqual, 1)
			So(len(strings.Split(response[0], "^")), ShouldEqual, 5)
		})

	})

	Convey("Given I visit \"/generate/1?delimiter=%2F-%2F", t, func() {

		req, err := http.NewRequest("GET", "/generate/1?delimiter=%2F-%2F", nil)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()
		params := httprouter.Params{httprouter.Param{Key: "amount", Value: "1"}}

		Convey("I get a json array of one password delimited by \"/-/\"", func() {
			main.Generate(w, req, params)
			So(w.Code, ShouldEqual, 200)
			var response []string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			So(len(response), ShouldEqual, 1)
			So(len(strings.Split(response[0], "/-/")), ShouldEqual, 5)
		})

	})

	Convey("Given I visit \"/generate/1?prefix=^", t, func() {

		req, err := http.NewRequest("GET", "/generate/1?prefix=^", nil)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()
		params := httprouter.Params{httprouter.Param{Key: "amount", Value: "1"}}

		Convey("I get a json array of one password prefixed by \"^\"", func() {
			main.Generate(w, req, params)
			So(w.Code, ShouldEqual, 200)
			var response []string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			So(len(response), ShouldEqual, 1)
			first, _ := utf8.DecodeRuneInString(response[0])
			So(string(first), ShouldEqual, "^")
		})

	})

	Convey("Given I visit \"/generate/1?prefix=%2F-%2F", t, func() {

		req, err := http.NewRequest("GET", "/generate/1?prefix=%2F-%2F", nil)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()
		params := httprouter.Params{httprouter.Param{Key: "amount", Value: "1"}}

		Convey("I get a json array of one password prefixed by \"/-/\"", func() {
			main.Generate(w, req, params)
			So(w.Code, ShouldEqual, 200)
			var response []string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			So(len(response), ShouldEqual, 1)
			prefixCharacterCount := 3
			var startingCharacters = make([]string, prefixCharacterCount)
			for i, character := range response[0] {
				startingCharacters[i] = string(character)
				if i == prefixCharacterCount-1 {
					break
				}
			}
			So(strings.Join(startingCharacters, ""), ShouldEqual, "/-/")
		})

	})

	Convey("Given I visit \"/generate/1?suffix=^", t, func() {

		req, err := http.NewRequest("GET", "/generate/1?suffix=^", nil)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()
		params := httprouter.Params{httprouter.Param{Key: "amount", Value: "1"}}

		Convey("I get a json array of one password suffixed by \"^\"", func() {
			main.Generate(w, req, params)
			So(w.Code, ShouldEqual, 200)
			var response []string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			So(len(response), ShouldEqual, 1)
			last, _ := utf8.DecodeLastRuneInString(response[0])
			So(string(last), ShouldEqual, "^")
		})

	})

	Convey("Given I visit \"/generate/1?suffix=%2F-%2F", t, func() {

		req, err := http.NewRequest("GET", "/generate/1?suffix=%2F-%2F", nil)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()
		params := httprouter.Params{httprouter.Param{Key: "amount", Value: "1"}}

		Convey("I get a json array of one password suffixed by \"/-/\"", func() {
			main.Generate(w, req, params)
			So(w.Code, ShouldEqual, 200)
			var response []string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			So(len(response), ShouldEqual, 1)
			suffixCharacterCount := 3
			var endingCharacters = string(response[0][len(response[0])-suffixCharacterCount:])
			So(endingCharacters, ShouldEqual, "/-/")
		})

	})

	Convey("Given I visit \"/generate/1?titleCase=true", t, func() {

		Convey("I get a json array of one password In Title Case", nil)

	})

	Convey("Given I visit \"/generate/1/safe", t, func() {

		Convey("I get a json array of one password that will satisfy most site password requirements", nil)

	})

}
