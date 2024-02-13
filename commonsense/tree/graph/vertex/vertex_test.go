package vertex

import (
	"bufio"
	"bytes"
	"fmt"
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

	dfs, err := v1.DepthFirstSearch(1)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v1)

	dfs, err = v2.DepthFirstSearch(1)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v1)

	dfs, err = v3.DepthFirstSearch(1)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v1)

	dfs, err = v1.DepthFirstSearch(2)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v2)

	dfs, err = v2.DepthFirstSearch(2)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v2)

	dfs, err = v3.DepthFirstSearch(2)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v2)

	dfs, err = v1.DepthFirstSearch(3)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v3)

	dfs, err = v2.DepthFirstSearch(3)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v3)

	dfs, err = v3.DepthFirstSearch(3)
	assert.Nil(t, err)
	assert.Equal(t, dfs, v3)
}

func TestVertex_BreadthFirstSearch(t *testing.T) {
	type args[T comparable] struct {
		value T
	}
	type testCase[T comparable] struct {
		name    string
		v       *Vertex[T]
		args    args[T]
		want    *Vertex[T]
		wantErr assert.ErrorAssertionFunc
	}

	v1 := NewVertex[int](1)
	v2 := NewVertex[int](2)
	v3 := NewVertex[int](3)
	v1.AddAdjacentVertex(v2)
	v1.AddAdjacentVertex(v3)

	tests := []testCase[int]{
		{
			name:    "want first vertx from v1",
			v:       v1,
			args:    args[int]{1},
			want:    v1,
			wantErr: assert.NoError,
		},
		{
			name:    "want first vertx from v2",
			v:       v2,
			args:    args[int]{1},
			want:    v1,
			wantErr: assert.NoError,
		},
		{
			name:    "want first vertx from v3",
			v:       v3,
			args:    args[int]{1},
			want:    v1,
			wantErr: assert.NoError,
		},
		{
			name:    "want second vertx from v1",
			v:       v1,
			args:    args[int]{2},
			want:    v2,
			wantErr: assert.NoError,
		},
		{
			name:    "want second vertx from v2",
			v:       v2,
			args:    args[int]{2},
			want:    v2,
			wantErr: assert.NoError,
		},
		{
			name:    "want second vertx from v3",
			v:       v3,
			args:    args[int]{2},
			want:    v2,
			wantErr: assert.NoError,
		},
		{
			name:    "want third vertx from v1",
			v:       v1,
			args:    args[int]{3},
			want:    v3,
			wantErr: assert.NoError,
		},
		{
			name:    "want third vertx from v2",
			v:       v2,
			args:    args[int]{3},
			want:    v3,
			wantErr: assert.NoError,
		},
		{
			name:    "want third vertx from v3",
			v:       v3,
			args:    args[int]{3},
			want:    v3,
			wantErr: assert.NoError,
		},
		{
			name:    "not found",
			v:       v3,
			args:    args[int]{9},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.BreadthFirstSearch(tt.args.value)
			if !tt.wantErr(t, err, fmt.Sprintf("BreadthFirstSearch(%v)", tt.args.value)) {
				return
			}
			assert.Equalf(t, tt.want, got, "BreadthFirstSearch(%v)", tt.args.value)
		})
	}
}

func TestVertex_BreadthFirstSearchTraversal(t *testing.T) {
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
	_ = v1.BreadthFirstSearchTraversal()
	_ = w.Close()

	// reset the output back to orig std out after test
	os.Stdout = origStdout
	output := <-outCh

	expected := "1\n2\n3\n"
	assert.Equal(t, expected, output)
}
