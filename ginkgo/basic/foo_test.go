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
		By("Context level [By] 0")
		BeforeEach(func() {
			f = Foo{
				Bar: "fbar1",
			}
		})
		g.Bar = "gbar1"

		It("should be a fool", func() {
			By("It level [By] 0-0")
			f.Bar = "fbar2"
			fmt.Printf("Foo f: %v\n", Display(f)) // Foo f: fbar2
			fmt.Printf("Foo g: %v\n", Display(g)) // Foo g: gbar1
		})

		It("should not work", func() {
			By("It level [By] 0-1")
			fmt.Printf("Foo f: %v\n", Display(f)) // Foo f: fbar1
			fmt.Printf("Foo g: %v\n", Display(g)) // Foo g: gbar1
		})
	})

	Context("when it is in the weekday", func() {
		By("Context level [By] 1")
		It("should not be a fool", func() {
			By("It level [By] 1-0")
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

/*
STEP: Context level [By] 0
STEP: Context level [By] 1
Running Suite: Basic Suite
==========================
Random Seed: 1545446770
Will run 4 of 4 specs

Foo when it is at the weekend
  should be a fool
  /root/go/src/github.com/fatsheep9146/mixture/ginkgo/basic/foo_test.go:33
STEP: It level [By] 0-0
Foo f: fbar2
Foo g: gbar1
•
------------------------------
Foo when it is at the weekend
  should not work
  /root/go/src/github.com/fatsheep9146/mixture/ginkgo/basic/foo_test.go:40
STEP: It level [By] 0-1
Foo f: fbar1
Foo g: gbar1
•
------------------------------
Foo when it is in the weekday
  should not be a fool
  /root/go/src/github.com/fatsheep9146/mixture/ginkgo/basic/foo_test.go:49
STEP: It level [By] 1-0
Foo f: fbar0
Foo g: gbar1
•
------------------------------
Foo test nested describe when test nested describe
  should display like what
  /root/go/src/github.com/fatsheep9146/mixture/ginkgo/basic/foo_test.go:58
•
Ran 4 of 4 Specs in 0.000 seconds
SUCCESS! -- 4 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS
*/
