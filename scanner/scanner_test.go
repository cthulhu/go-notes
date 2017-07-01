package scanner_test

import (
	"context"
	"time"

	. "github.com/cthulhu/go-notes/scanner"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scanner", func() {
	Context("On set of dirrectories and files", func() {
		It("returns go files", func() {
			ctx := context.Background()
			paths, errors := New(ctx, []string{"fixtures"})
			Expect(<-paths).To(Equal("fixtures/folder1/file1.go"))
			Expect(<-paths).To(Equal("fixtures/folder2/file2.go"))
			time.Sleep(10 * time.Millisecond)
			Expect(errors).To(BeClosed())
			Expect(paths).To(BeClosed())
		})
	})
	Context("On empty array", func() {
		It("returns nothing", func() {
			ctx := context.Background()
			paths, errors := New(ctx, []string{})
			time.Sleep(10 * time.Millisecond)
			Expect(paths).To(BeClosed())
			Expect(errors).To(BeClosed())
		})
	})
})
