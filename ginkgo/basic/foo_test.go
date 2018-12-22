package foo_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	//	. "github.com/onsi/gomega"

	. "github.com/fatsheep9146/mixture/ginkgo/basic"
)

var _ = Describe("Foo", func() {
	var f Foo
	var g Foo
	BeforeEach(func() {
		f = Foo{
			Bar: "bar",
		}
	})

	g = Foo{
		Bar: "gbar0",
	}

	Context("when it is at the weekend", func() {
		BeforeEach(func() {
			f = Foo{
				Bar: "dog",
			}
		})
		g.Bar = "gbar1"

		It("should be a fool", func() {
			f.Bar = "cat"
			fmt.Printf("Foo f: %v", Display(f))
			fmt.Printf("Foo g: %v", Display(g))
		})

		It("should not work", func() {
			fmt.Printf("Foo f: %v", Display(f))
			fmt.Printf("Foo g: %v", Display(g))
		})
	})

	Context("when it is in the weekday", func() {
		It("should not be a fool", func() {
			fmt.Printf("Foo f: %v", Display(f))
			fmt.Printf("Foo g: %v", Display(g))
		})
	})

})
