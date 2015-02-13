package hello

import (
	"github.com/ppalucki/hello"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(t *testing.T) {
	// got := Simple()
	// if got != "hello world" {
	// 	t.Fail()
	// }

	Convey("I want my hw example working", t, func() {
		So(Simple(), ShouldEqual, "hello world")
	})

	// I like asserts more
	assert.Equal(t, "hello world", Simple())
}

func TestSkipped(t *testing.T) {
	t.Skip()
}
