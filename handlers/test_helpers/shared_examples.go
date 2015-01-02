package test_helpers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

func ItFailsWhenTheContentLengthIsMissing(req **http.Request, resp **httptest.ResponseRecorder, server **ghttp.Server) {
	Context("uploading the file, when the request is missing content length", func() {
		BeforeEach(func() {
			(*req).ContentLength = -1
		})

		It("does not make the request to CC", func() {
			Ω((*server).ReceivedRequests()).Should(HaveLen(0))
		})

		It("responds with 411", func() {
			Ω((*resp).Code).Should(Equal(http.StatusLengthRequired))
		})
	})
}

func ItHandlesCCFailures(postStatusCode *int, resp **httptest.ResponseRecorder, server **ghttp.Server) {
	Context("when CC returns a non-succesful status code", func() {
		BeforeEach(func() {
			*postStatusCode = 403
		})

		It("makes the request to CC", func() {
			Ω((*server).ReceivedRequests()).Should(HaveLen(1))
		})

		It("responds with the status code from the CC request", func() {
			Ω((*resp).Code).Should(Equal(*postStatusCode))

			data, err := ioutil.ReadAll((*resp).Body)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(string(data)).Should(ContainSubstring(strconv.Itoa(*postStatusCode)))
		})
	})
}
