package vertex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertex_AddAdjacentVertex(t *testing.T) {
	v1 := New[string]("one")
	v2 := New[string]("two")

	v1.addAdjacentVertx(v2, 1)
	v2.addAdjacentVertx(v1, 2)

	// v1 slice contains v2 testify assert
	assert.Equal(t, 1, v1.Adjacent[v2])
}

//func TestVertx_addAdjacentVertx(t *testing.T) {
//	type args[T comparable] struct {
//		vertx *Vertx[T]
//		value int
//	}
//	type testCase[T comparable] struct {
//		name string
//		v    Vertx[T]
//		args args[T]
//	}
//	tests := []testCase[string]{
//		{
//			name: "add key pair",
//			v:    Vertx[string]{Value: "first"},
//			args: args[string]{
//				vertx: &Vertx[string]{Value: "second"},
//				value: 2,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.v.addAdjacentVertx(tt.args.vertx, tt.args.value)
//			assert.Equal(t, 2, tt.v.Adjacent[tt.args.vertx])
//		})
//	}
//}
