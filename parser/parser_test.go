package parser_test

import (
	. "github.com/cthulhu/go-notes/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Context("All annotations", func() {
		It("returns information about annotations", func() {
			parser := New(false, false, false, "")
			Expect(parser.Parse("fixtures/file1.go")).NotTo(HaveOccurred())
			aggregated := parser.Aggregate()
			Expect(aggregated).To(ContainSubstring("TODO: 1"))
			Expect(aggregated).To(ContainSubstring("FIXME: 2"))
			Expect(aggregated).To(ContainSubstring("OPTIMIZE: 3"))
			Expect(aggregated).NotTo(ContainSubstring("BUG: 4"))
			Expect(aggregated).NotTo(ContainSubstring("BAD SMELL: 5"))
		})
	})
	Context("only todos", func() {
		It("returns information about annotations", func() {
			parser := New(false, true, false, "")
			Expect(parser.Parse("fixtures/file1.go")).NotTo(HaveOccurred())
			aggregated := parser.Aggregate()
			Expect(aggregated).To(ContainSubstring("TODO: 1"))
			Expect(aggregated).NotTo(ContainSubstring("FIXME: 2"))
			Expect(aggregated).NotTo(ContainSubstring("OPTIMIZE: 3"))
			Expect(aggregated).NotTo(ContainSubstring("BUG: 4"))
			Expect(aggregated).NotTo(ContainSubstring("BAD SMELL: 5"))
		})
	})
	Context("only fixme's", func() {
		It("returns information about annotations", func() {
			parser := New(true, false, false, "")
			Expect(parser.Parse("fixtures/file1.go")).NotTo(HaveOccurred())
			aggregated := parser.Aggregate()
			Expect(aggregated).NotTo(ContainSubstring("TODO: 1"))
			Expect(aggregated).To(ContainSubstring("FIXME: 2"))
			Expect(aggregated).NotTo(ContainSubstring("OPTIMIZE: 3"))
			Expect(aggregated).NotTo(ContainSubstring("BUG: 4"))
			Expect(aggregated).NotTo(ContainSubstring("BAD SMELL: 5"))
		})
	})
	Context("only optimizes", func() {
		It("returns information about annotations", func() {
			parser := New(false, false, true, "")
			Expect(parser.Parse("fixtures/file1.go")).NotTo(HaveOccurred())
			aggregated := parser.Aggregate()
			Expect(aggregated).NotTo(ContainSubstring("TODO: 1"))
			Expect(aggregated).NotTo(ContainSubstring("FIXME: 2"))
			Expect(aggregated).To(ContainSubstring("OPTIMIZE: 3"))
			Expect(aggregated).NotTo(ContainSubstring("BUG: 4"))
			Expect(aggregated).NotTo(ContainSubstring("BAD SMELL: 5"))
		})
	})
	Context("custom annotation", func() {
		It("returns information about BUG", func() {
			parser := New(false, false, false, "BUG")
			Expect(parser.Parse("fixtures/file1.go")).NotTo(HaveOccurred())
			aggregated := parser.Aggregate()
			Expect(aggregated).NotTo(ContainSubstring("TODO: 1"))
			Expect(aggregated).NotTo(ContainSubstring("FIXME: 2"))
			Expect(aggregated).NotTo(ContainSubstring("OPTIMIZE: 3"))
			Expect(aggregated).To(ContainSubstring("BUG: 4"))
			Expect(aggregated).NotTo(ContainSubstring("BAD SMELL: 5"))
		})
		It("returns information about BAD SMELL", func() {
			parser := New(false, false, false, "BUG")
			Expect(parser.Parse("fixtures/file1.go")).NotTo(HaveOccurred())
			aggregated := parser.Aggregate()
			Expect(aggregated).NotTo(ContainSubstring("TODO: 1"))
			Expect(aggregated).NotTo(ContainSubstring("FIXME: 2"))
			Expect(aggregated).NotTo(ContainSubstring("OPTIMIZE: 3"))
			Expect(aggregated).To(ContainSubstring("BUG: 4"))
			Expect(aggregated).NotTo(ContainSubstring("BAD SMELL: 5"))
		})
	})
})
