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
			Bar: "fbar0",
		}
	})

	g = Foo{
		Bar: "gbar0",
	}

	Context("when it is at the weekend", func() {
		By("Fix Foo f")
		BeforeEach(func() {
			f = Foo{
				Bar: "fbar1",
			}
		})
		g.Bar = "gbar1"

		It("should be a fool", func() {
			f.Bar = "fbar2"
			fmt.Printf("Foo f: %v\n", Display(f)) // Foo f: fbar2
			fmt.Printf("Foo g: %v\n", Display(g)) // Foo g: gbar1
		})

		It("should not work", func() {
			fmt.Printf("Foo f: %v\n", Display(f)) // Foo f: fbar1
			fmt.Printf("Foo g: %v\n", Display(g)) // Foo g: gbar1
		})
	})

	Context("when it is in the weekday", func() {
		It("should not be a fool", func() {
			fmt.Printf("Foo f: %v\n", Display(f)) // Foo f: fbar0
			fmt.Printf("Foo g: %v\n", Display(g)) // Foo g: gbar1
		})
	})

	Describe("test nested describe", func() {
		Context("when test nested describe", func() {
			It("should display like what", func() {

			})
		})
	})

})
