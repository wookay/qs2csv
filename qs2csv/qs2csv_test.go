package qs2csv

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestQS2CSV(t *testing.T) {
	var cmd string
	cmd = "echo 'a=1&b=2' | go run qs2csv.go -c a"
	out, _ := exec.Command("sh", "-c", cmd).Output()
	assert.Equal(t, "a\n1\n", string(out))

	cmd = "echo 'a=1&b=2' | go run qs2csv.go -c b"
	out, _ = exec.Command("sh", "-c", cmd).Output()
	assert.Equal(t, "b\n2\n", string(out))
}
