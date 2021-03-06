package integration_test

import (
	"io"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Integration", func() {
	var (
		inBuf  *gbytes.Buffer
		cliCmd *exec.Cmd
	)

	BeforeEach(func() {
		cliCmd = exec.Command(cliBin)
		inBuf = gbytes.NewBuffer()
		cliCmd.Stdin = inBuf
	})

	Describe("'I': setting image size", func() {
		Context("when executing the program", func() {
			It("the user can enter a image size", func() {
				_, err := io.WriteString(inBuf, "I 5 5")
				Expect(err).NotTo(HaveOccurred())

				session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Err.Contents).Should(HaveLen(0))
			})
		})

		Context("if the user specifies something other than a image size", func() {
			// TODO: this can probably be changed to repeat the request
			It("the program will complain and exit", func() {
				_, err := io.WriteString(inBuf, "X 5 5")
				Expect(err).NotTo(HaveOccurred())

				session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("invalid image value: unrecognised command 'X', use 'I' for Image initialisation"))
			})
		})
	})

	Describe("'L': colouring an individual pixel", func() {
		It("sets the pixel at the given coordinates to a given colour", func() {
			_, err := io.WriteString(inBuf, "I 5 5\nL 1 3 A\nS")
			Expect(err).NotTo(HaveOccurred())

			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("OOOOO\nOOOOO\nAOOOO\nOOOOO\nOOOOO\n"))
		})

		Context("if the action cannot be processed", func() {
			It("prints an error", func() {
				_, err := io.WriteString(inBuf, "I 5 5\nL 1 6 A")
				Expect(err).NotTo(HaveOccurred())

				session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("given coordinate is beyond image grid"))
			})
		})
	})

	Describe("'V': colouring a vertical line", func() {
		It("sets the pixels between the specified coordinates", func() {
			_, err := io.WriteString(inBuf, "I 5 5\nV 2 3 5 W\nS")
			Expect(err).NotTo(HaveOccurred())

			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("OOOOO\nOOOOO\nOWOOO\nOWOOO\nOWOOO\n"))
		})

		Context("if the action cannot be processed", func() {
			It("prints an error", func() {
				_, err := io.WriteString(inBuf, "I 5 5\nV 7 2 5 W")
				Expect(err).NotTo(HaveOccurred())

				session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("given coordinate is beyond image grid"))
			})
		})
	})

	Describe("'H': colouring a horizontal line", func() {
		It("sets the pixels between the specified coordinates", func() {
			_, err := io.WriteString(inBuf, "I 5 5\nH 3 5 2 Z\nS")
			Expect(err).NotTo(HaveOccurred())

			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("OOOOO\nOOZZZ\nOOOOO\nOOOOO\nOOOOO\n"))
		})

		Context("if the action cannot be processed", func() {
			It("prints an error", func() {
				_, err := io.WriteString(inBuf, "I 5 5\nH 3 5 7 Z")
				Expect(err).NotTo(HaveOccurred())

				session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("given coordinate is beyond image grid"))
			})
		})
	})

	Describe("'S': showing the image", func() {
		Context("whenever the user inputs the 'S' command", func() {
			It("the image is printed in its current state", func() {
				_, err := io.WriteString(inBuf, "I 5 5\nS")
				Expect(err).NotTo(HaveOccurred())

				session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("OOOOO\nOOOOO\nOOOOO\nOOOOO\nOOOOO\n"))
			})
		})
	})

	Describe("'C': clearing the image", func() {
		It("the image pixels are cleared", func() {
			_, err := io.WriteString(inBuf, "I 2 2\nL 1 1 A\nS\nC\nS")
			Expect(err).NotTo(HaveOccurred())

			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("AO\nOO\n\nOO\nOO\n"))
		})
	})

	Context("any other attempted action", func() {
		It("complains", func() {
			_, err := io.WriteString(inBuf, "I 2 2\nP 1 1 A\n")
			Expect(err).NotTo(HaveOccurred())

			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("invalid action"))
		})
	})
})
