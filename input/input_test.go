package input_test

import (
	"bufio"
	"io"

	"github.com/mo-work/go-technical-test-for-claudia/input"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Input", func() {
	var (
		buffer *gbytes.Buffer
		i      input.Input
	)

	BeforeEach(func() {
		buffer = gbytes.NewBuffer()
		scanner := bufio.NewScanner(buffer)
		i = input.New(scanner)
	})

	Describe("GetImageSize", func() {
		It("receives the image size from stdin", func() {
			_, err := io.WriteString(buffer, "I 5 7")
			Expect(err).NotTo(HaveOccurred())

			x, y, err := i.GetImageSize()
			Expect(err).NotTo(HaveOccurred())
			Expect(x).To(Equal(5))
			Expect(y).To(Equal(7))
		})

		Context("if the argument for the x azis cannot be translated into an integer", func() {
			It("fails", func() {
				_, err := io.WriteString(buffer, "I x 7")
				Expect(err).NotTo(HaveOccurred())

				_, _, err = i.GetImageSize()
				Expect(err).To(MatchError("could not parse non-integer 'x'"))
			})
		})

		Context("if the argument for the y azis cannot be translated into an integer", func() {
			It("fails", func() {
				_, err := io.WriteString(buffer, "I 7 y")
				Expect(err).NotTo(HaveOccurred())

				_, _, err = i.GetImageSize()
				Expect(err).To(MatchError("could not parse non-integer 'y'"))
			})
		})

		Context("if the argument for the x axis size is less than the min value", func() {
			It("fails", func() {
				_, err := io.WriteString(buffer, "I 0 7")
				Expect(err).NotTo(HaveOccurred())

				_, _, err = i.GetImageSize()
				Expect(err).To(MatchError("image axis out of range: 1 <= M,N <= 1024"))
			})
		})

		Context("if the argument for the y axis size is less than the min value", func() {
			It("fails", func() {
				_, err := io.WriteString(buffer, "I 7 0")
				Expect(err).NotTo(HaveOccurred())

				_, _, err = i.GetImageSize()
				Expect(err).To(MatchError("image axis out of range: 1 <= M,N <= 1024"))
			})
		})

		Context("if the argument for the x axis size is greater than the max value", func() {
			It("fails", func() {
				_, err := io.WriteString(buffer, "I 1025 7")
				Expect(err).NotTo(HaveOccurred())

				_, _, err = i.GetImageSize()
				Expect(err).To(MatchError("image axis out of range: 1 <= M,N <= 1024"))
			})
		})

		Context("if the argument for the y axis size is greater than the max value", func() {
			It("fails", func() {
				_, err := io.WriteString(buffer, "I 7 1025")
				Expect(err).NotTo(HaveOccurred())

				_, _, err = i.GetImageSize()
				Expect(err).To(MatchError("image axis out of range: 1 <= M,N <= 1024"))
			})
		})
	})
})
