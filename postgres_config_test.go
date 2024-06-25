package flag_test

import (
	"github.com/concourse/flag/v2"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PostgresConfig", func() {
	Describe("ConnectionString", func() {
		It("escapes values correctly", func() {
			Expect(flag.PostgresConfig{
				Host: "1.2.3.4",
				Port: 5432,

				User:     "some user",
				Password: "password \\ with ' funny ! chars",

				SSLMode: "verify-full",

				Database: "atc",
			}.ConnectionString()).To(Equal("dbname='atc' host='1.2.3.4' password='password \\\\ with \\' funny ! chars' port=5432 sslmode='verify-full' user='some user'"))
		})
	})
})

var _ = Describe("PostgresConfig", func() {
	Describe("ConnectionString", func() {
		It("adds binary_parameters correctly", func() {
			Expect(flag.PostgresConfig{
				Host: "1.2.3.4",
				Port: 5432,

				User:     "some user",
				Password: "not-so-important",

				BinaryParameters: true,

				Database: "atc",
			}.ConnectionString()).To(Equal("binary_parameters='yes' dbname='atc' host='1.2.3.4' password='not-so-important' port=5432 sslmode='' user='some user'"))
		})
	})
})
