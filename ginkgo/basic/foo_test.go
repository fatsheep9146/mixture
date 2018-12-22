package foo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fatsheep9146/mixture/ginkgo/basic"
)

var _ = Describe("Foo", func() {

	Context("when it is at the weekend", func() {
		It("should be a fool", func() {
		})

		It("should not work", func() {

		})
	})

	Context("when it is in the weekday", func() {
		It("should not be a fool", func() {
		})
	})

})
