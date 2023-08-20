package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFetch(t *testing.T) {
	Convey("it can render html to md without code block highlight", t, func() {
		f, err := os.Open("./test_render.html")
		So(err, ShouldBeNil)
		defer f.Close()

		txt := render(f)
		So(txt, ShouldContainSubstring, "```\n")
	})
}
