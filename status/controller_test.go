package status_test

import (
	"encoding/json"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/danielpsf/go-dummy-ms/status"
)

type ResponseWriterMock struct {
	writeCall struct {
		times   int
		receive []byte
	}
	headersCall struct {
		times   int
		content map[string][]string
	}
}

func (responseWriterMock *ResponseWriterMock) Write(data []byte) (int, error) {
	responseWriterMock.writeCall.times++
	responseWriterMock.writeCall.receive = data
	return http.StatusOK, nil
}

func (responseWriterMock *ResponseWriterMock) Header() http.Header {
	responseWriterMock.headersCall.times++
	return responseWriterMock.headersCall.content
}

func (responseWriterMock *ResponseWriterMock) WriteHeader(int) {}

var _ = Describe("Status/Controller", func() {
	var responseWriterMock *ResponseWriterMock

	BeforeEach(func() {
		responseWriterMock = &ResponseWriterMock{}
		responseWriterMock.writeCall.times = 0
		responseWriterMock.headersCall.times = 0
		responseWriterMock.headersCall.content = map[string][]string{}
	})

	It("should write the JSON in the HTTP response body", func() {
		status.Check(responseWriterMock, nil)

		Expect(responseWriterMock.writeCall.times).To(Equal(1))

		var httpResponse status.Response
		json.Unmarshal(responseWriterMock.writeCall.receive, &httpResponse)
		Expect(httpResponse).To(Equal(status.Response{Healthy: true}))
	})

	It("should write the proper HTTP headers", func() {
		status.Check(responseWriterMock, nil)

		Expect(responseWriterMock.headersCall.times).To(Equal(2))
		Expect(responseWriterMock.headersCall.content).Should(HaveKey("Content-Type"))
		Expect(responseWriterMock.headersCall.content).Should(HaveKey("X-Powered-By"))
	})
})
