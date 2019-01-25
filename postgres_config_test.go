package flag_test

import (
	"github.com/concourse/flag"
	"time"

	. "github.com/onsi/ginkgo"
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
	Describe("Timeout", func() {
		It("check config Timeout correctly", func() {
			pgConfig := flag.PostgresConfig{
				Host:           "1.2.3.4",
				Port:           5432,
				User:           "some user",
				Password:       "password",
				SSLMode:        "verify-full",
				Database:       "atc",
				ConnectTimeout: 20 * time.Second,
				ReadTimeout:    15 * time.Second,
				WriteTimeout:   30 * time.Second,
			}
			Expect(int(pgConfig.ConnectTimeout.Seconds())).To(Equal(20))
			Expect(int(pgConfig.ReadTimeout.Seconds())).To(Equal(15))
			Expect(int(pgConfig.WriteTimeout.Seconds())).To(Equal(30))
		})
	})
})
