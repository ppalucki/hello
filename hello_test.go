package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(t *testing.T) {
	// got := simple()
	// if got != "hello world" {
	// 	t.Fail()
	// }

	Convey("I want my hw example working", t, func() {
		So(simple(), ShouldEqual, "hello world")
	})

	// I like asserts more
	assert.Equal(t, "hello world", simple())
}

func TestSkipped(t *testing.T) {
	t.Skip()
}
