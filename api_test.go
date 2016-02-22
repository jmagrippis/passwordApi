package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenerator(t *testing.T) {

	Convey("Given I visit \"/generate/:amount\"", t, func() {

		Convey("It returns a json array of :amount passwords", nil)

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
