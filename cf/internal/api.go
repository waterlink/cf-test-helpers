package cfinternal

import (
	"encoding/json"
	"strings"
	"time"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

type starter interface {
	Start(string, ...string) (*Session, error)
}

func ApiRequest(cmdStarter starter, method, endpoint string, response interface{}, timeout time.Duration, data ...string) {
	args := []string{
		"curl",
		endpoint,
		"-X", method,
	}

	dataArg := strings.Join(data, "")
	if len(dataArg) > 0 {
		args = append(args, "-d", dataArg)
	}

	request, err := cmdStarter.Start("cf", args...)
	ExpectWithOffset(2, err).NotTo(HaveOccurred())

	request.Wait(timeout)
	ExpectWithOffset(2, request).To(Exit(0))

	if response != nil {
		err := json.Unmarshal(request.Out.Contents(), response)
		ExpectWithOffset(2, err).ToNot(HaveOccurred())
	}
}