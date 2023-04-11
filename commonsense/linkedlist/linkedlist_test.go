package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func testNodes() node.Node[string] {
//	nodeThree := node.Node[string]{
//		Value: "d",
//	}
//	nodeTwo := node.Node[string]{
//		Value: "c",
//		Next: &nodeThree;
//	}
//	nodeOne := node.Node[string]{
//		Value: "b",
//		Next: &nodeTwo,
//	}
//
//	nodeHead := node.Node[string]{
//		Value: "a",
//		Next: &nodeOne,
//	}
//	return nodeHead
//}

func TestLinkedList_Add(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")
	list.Add("e")
	list.Add("f")

	assert.Equal(t, list.Head.Data, "a")
	assert.Equal(t, list.Head.Next.Data, "b")
	assert.Equal(t, list.Head.Next.Next.Data, "c")
	assert.Equal(t, list.Head.Next.Next.Next.Data, "d")
	assert.Equal(t, list.Head.Next.Next.Next.Next.Data, "e")
	assert.Equal(t, list.Head.Next.Next.Next.Next.Next.Data, "f")
}
func TestLinkedList_AddByIndex(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")
	list.Add("e")
	list.Add("f")

	assert.Equal(t, list.Head.Data, "a")
	assert.Equal(t, list.Head.Next.Data, "b")
	assert.Equal(t, list.Head.Next.Next.Data, "c")
	assert.Equal(t, list.Head.Next.Next.Next.Data, "d")
	assert.Equal(t, list.Head.Next.Next.Next.Next.Data, "e")
	assert.Equal(t, list.Head.Next.Next.Next.Next.Next.Data, "f")
}

func TestLinkedList_Delete(t *testing.T) {
	//type fields struct {
	//	head node.Node
	//}
	//type args struct {
	//	index int
	//}
	//tests := []struct {
	//	name    string
	//	fields  fields
	//	args    args
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		l := &LinkedList{
	//			head: tt.fields.head,
	//		}
	//		if err := l.Delete(tt.args.index); (err != nil) != tt.wantErr {
	//			t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
	//		}
	//	})
	//}
}

func TestLinkedList_Read(t *testing.T) {
	//type fields struct {
	//	head node.Node
	//}
	//type args struct {
	//	index int
	//}
	//tests := []struct {
	//	name    string
	//	fields  fields
	//	args    args
	//	want    T
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		l := &LinkedList{
	//			head: tt.fields.head,
	//		}
	//		got, err := l.Read(tt.args.index)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("Read() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

func TestLinkedList_Search(t *testing.T) {
	//type fields struct {
	//	head node.Node
	//}
	//type args struct {
	//	index T
	//}
	//tests := []struct {
	//	name    string
	//	fields  fields
	//	args    args
	//	want    int
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		l := &LinkedList{
	//			head: tt.fields.head,
	//		}
	//		got, err := l.Search(tt.args.index)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if got != tt.want {
	//			t.Errorf("Search() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
