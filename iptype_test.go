package iputility

import (
	"math/rand"
	"net"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ipType      Ip
	tstEndpoint string
)

var _ = Describe("ip type module", func() {

	Describe("GetType", func() {

		BeforeEach(func() {
			rand.Seed(time.Now().UTC().UnixNano())
		})

		JustBeforeEach(func() {
			ipType = GetType(tstEndpoint)
		})

		Context("when input is IP address", func() {

			BeforeEach(func() {
				tstEndpoint = "10.11.12.14"
			})

			It("returns the correct object", func() {
				Expect(ipType.Type).To(Equal(IPTYPE_ADDRESS))
				Expect(ipType.Endpoint).To(Equal(tstEndpoint))
			})
		})

		Context("when input is IP cidr", func() {

			BeforeEach(func() {
				tstEndpoint = "10.11.12.14/16"
			})

			It("returns the correct object", func() {
				Expect(ipType.Type).To(Equal(IPTYPE_CIDR))
				Expect(ipType.Endpoint).To(Equal(tstEndpoint))
			})

			Context("when CIDR is /32", func() {
				BeforeEach(func() {
					tstEndpoint = "10.11.12.14/32"
				})

				It("returns the correct object", func() {
					Expect(ipType.Type).To(Equal(IPTYPE_ADDRESS))
					Expect(ipType.Endpoint).To(Equal("10.11.12.14"))
				})
			})
		})

		Context("when input is IP range", func() {

			BeforeEach(func() {
				tstEndpoint = "10.11.12.14-10.11.12.255"
			})

			It("returns the correct object", func() {
				Expect(ipType.Type).To(Equal(IPTYPE_RANGE))
				Expect(ipType.Endpoint).To(Equal(tstEndpoint))
			})

			Context("when range lo equals range hi", func() {
				BeforeEach(func() {
					tstEndpoint = "10.11.12.14-10.11.12.14"
				})

				It("returns the correct object", func() {
					Expect(ipType.Type).To(Equal(IPTYPE_ADDRESS))
					Expect(ipType.Endpoint).To(Equal("10.11.12.14"))
				})
			})
		})

		Context("when input is FQDN", func() {

			BeforeEach(func() {
				tstEndpoint = "github.com"
			})

			It("returns the correct object", func() {
				Expect(ipType.Type).To(Equal(IPTYPE_FQDN))
				Expect(ipType.Endpoint).To(Equal(tstEndpoint))
			})
		})

		Context("when input is bad", func() {

			BeforeEach(func() {
				tstEndpoint = "badcom"
			})

			It("returns the correct object", func() {
				Expect(ipType.Type).To(Equal(IPTYPE_UNDEFINED))
				Expect(ipType.Endpoint).To(Equal(""))
			})
		})
	})

	Describe("type In", func() {

		var t, t1 Ip
		var e, e1 string
		var isIn bool

		JustBeforeEach(func() {
			t = GetType(e)
			t1 = GetType(e1)

			isIn = t.In(t1)
		})

		Context("when t equals lower limit of t1 but its hi is less than hi of t1", func() {

			BeforeEach(func() {
				e = "1.2.3.0-1.2.3.254"
				e1 = "1.2.3.0/24"
			})

			It("returns true", func() {
				Expect(isIn).To(Equal(true))
			})
		})

		Context("when t greater than lower limit of t1 but its hi is equal to hi of t1", func() {

			BeforeEach(func() {
				e = "1.2.3.1-1.2.3.255"
				e1 = "1.2.3.0/24"
			})

			It("returns true", func() {
				Expect(isIn).To(Equal(true))
			})
		})

		Context("when t equals t1", func() {
			BeforeEach(func() {
				e = "1.2.3.0-1.2.3.255"
				e1 = "1.2.3.0/24"
			})

			It("returns false", func() {
				Expect(isIn).To(Equal(false))
			})
		})

		Context("when type is not supported", func() {

			It("returns false when FQDN", func() {
				e = "google.com"
				e1 = "1.1.2.1"
				t = GetType(e)
				t1 = GetType(e1)
				Expect(t.In(t1)).To(Equal(false))
			})

			It("returns false when invalid", func() {
				e = "invalid_endpoint"
				e1 = "1.1.2.1"
				t = GetType(e)
				t1 = GetType(e1)
				Expect(t.In(t1)).To(Equal(false))
			})
		})
	})

	Describe("Equals", func() {

		var t, t1 Ip
		var e, e1 string
		var isEqual bool

		JustBeforeEach(func() {
			t = GetType(e)
			t1 = GetType(e1)

			isEqual = t.Equals(t1)
		})

		Context("when t equals t1", func() {
			BeforeEach(func() {
				e = "1.2.3.0-1.2.3.255"
				e1 = "1.2.3.0/24"
			})

			It("returns true", func() {
				Expect(isEqual).To(Equal(true))
			})
		})

		Context("when t is in t1", func() {
			BeforeEach(func() {
				e = "1.2.3.0/24"
				e1 = "1.2.3.0/23"
			})

			It("returns false", func() {
				Expect(isEqual).To(Equal(false))
			})
		})

		Context("when t1 is in t", func() {
			BeforeEach(func() {
				e = "1.2.3.0/23"
				e1 = "1.2.3.0/24"
			})

			It("returns false", func() {
				Expect(isEqual).To(Equal(false))
			})
		})

		Context("when type is not supported", func() {

			It("returns false when FQDN", func() {
				e = "google.com"
				e1 = "1.1.2.1"
				t = GetType(e)
				t1 = GetType(e1)
				Expect(t.Equals(t1)).To(Equal(false))
			})

			It("returns false when invalid", func() {
				e = "invalid_endpoint"
				e1 = "1.1.2.1"
				t = GetType(e)
				t1 = GetType(e1)
				Expect(t.Equals(t1)).To(Equal(false))
			})
		})
	})

	Describe("ip to byte conversion", func() {

		It("does it right", func() {

			ip := net.ParseIP("255.255.255.1")

			b := toUint(ip)

			var i uint64
			for i = 0; i < uint64(len(ip)); i++ {

				idx := len(ip) - 1 - int(i)

				Expect(ip[idx]).To(Equal(byte(b >> (i * 8))))
			}
		})

		It("return 0 when ip is nil", func() {
			Expect(toUint(nil)).To(Equal(uint64(0)))
		})
	})

	Describe("is Ip type check", func() {

		It("returns true when type is IPTYPE_ADDRESS", func() {
			e := "1.2.1.1"
			t := GetType(e)
			Expect(t.isIpType()).To(Equal(true))
		})

		It("returns true when type is IPTYPE_CIDR", func() {
			e := "1.2.1.1/24"
			t := GetType(e)
			Expect(t.isIpType()).To(Equal(true))
		})

		It("returns true when type is IPTYPE_RANGE", func() {
			e := "1.2.1.1-2.2.2.2"
			t := GetType(e)
			Expect(t.isIpType()).To(Equal(true))
		})

		It("returns false when type is IPTYPE_FQDN", func() {
			e := "google.com"
			t := GetType(e)
			Expect(t.isIpType()).To(Equal(false))
		})

		It("returns false when type is IPTYPE_UNDEFINED", func() {
			e := "nuffincom"
			t := GetType(e)
			Expect(t.isIpType()).To(Equal(false))
		})
	})
})
