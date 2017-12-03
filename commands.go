package iputility

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func nslookup(endpoint string, ok chan bool) {

	cmd := exec.Command("nslookup", endpoint)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		ok <- false
		return
	}

	if strings.Contains(out.String(), "server can't find") ||
		strings.Contains(out.String(), "Can't find") {
		ok <- false
		return
	}
	ok <- true
}

func ping(endpoint string, ok chan bool) {
	out, _ := exec.Command("ping", "-c", "3", "-W", "3", endpoint).CombinedOutput()

	if bytes.Contains(out, []byte("bytes from")) {
		ok <- true
		return
	} else {
		idx := bytes.Index(out, []byte("packets transmitted, "))
		if idx > -1 {
			sub := bytes.Trim(out[idx:], "packets transmitted, ")
			receivedCnt, err := strconv.Atoi(string(sub[0:1]))
			if err == nil && receivedCnt > 0 {
				ok <- true
			}
		}
	}

	ok <- false
}
