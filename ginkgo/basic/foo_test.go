package foo_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	//	. "github.com/onsi/gomega"

	. "github.com/fatsheep9146/mixture/ginkgo/basic"
)

var _ = Describe("Foo", func() {
	var f Foo

	BeforeEach(func() {
		f = Foo{
			Bar: "bar",
		}
	})

	Context("when it is at the weekend", func() {
		f.Bar = "dog"

		It("should be a fool", func() {
			f.Bar = "cat"
			fmt.Printf(Display(f))
		})

		It("should not work", func() {
			fmt.Printf(Display(f))
		})
	})

	Context("when it is in the weekday", func() {
		It("should not be a fool", func() {
			fmt.Printf(Display(f))
		})
	})

})
