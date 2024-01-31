package vertex

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertex_AddAdjacentVertex(t *testing.T) {
	v1 := NewVertex(1)
	v2 := NewVertex(2)

	v1.AddAdjacentVertex(v2)
	// v1 slice contains v2 testify assert
	assert.Equal(t, v2, v1.adjacentVertices[0])
	assert.Equal(t, v1, v2.adjacentVertices[0])

}

func TestVertx_DFS_Traverse(t *testing.T) {
	v1 := NewVertex(1)
	v2 := NewVertex(2)
	v3 := NewVertex(3)

	v1.AddAdjacentVertex(v2)
	v1.AddAdjacentVertex(v3)

	// create a pipe and reader
	r, w, err := os.Pipe()
	if err != nil {
		t.Error("Error creating pipe: ", err)
	}
	// save the original out stream
	origStdout := os.Stdout

	// redirect stdout to buf
	os.Stdout = w

	// capture the output in a separate goroutine because writing to the pipe is synchronous
	outCh := make(chan string)
	go func() {
		// create a buffer to capture your output
		var buf bytes.Buffer
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			buf.Write(scanner.Bytes())
			buf.WriteByte('\n')
		}
		outCh <- buf.String()
	}()

	// do work then close writer
	v1.DfsTraverse(v2)
	_ = w.Close()

	// reset the output back to orig std out after test
	os.Stdout = origStdout
	output := <-outCh

	expected := "2\n1\n3\n"
	assert.Equal(t, output, expected)
}

func TestDfs(t *testing.T) {
	v1 := NewVertex[int](1)
	v2 := NewVertex[int](2)
	v3 := NewVertex[int](3)

	v1.AddAdjacentVertex(v2)
	v1.AddAdjacentVertex(v3)

	dfs, err := v3.Dfs(1)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v1)
}
