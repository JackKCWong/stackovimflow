package cs

import (
	"context"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearch(t *testing.T) {
	Convey("Can Search", t, func() {
		c, err := NewClient(context.Background(),
			WithApiKey(os.Getenv("PCS_API_KEY")),
			WithEngineID(os.Getenv("PCS_CX")),
		)

		if err != nil {
			t.Fatal(err)
		}

		res, err := c.Search("How to read a file in Rust")
		if err != nil {
			t.Fatal(err)
		}

		So(len(res.Items), ShouldEqual, 5)
	})
}
