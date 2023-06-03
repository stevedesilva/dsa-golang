package classiclinkedlist

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Add(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")
	list.Add("e")
	list.Add("f")

	assert.Equal(t, list.GetHead().data, "a")
	assert.Equal(t, list.GetHead().next.data, "b")
	assert.Equal(t, list.GetHead().next.next.data, "c")
	assert.Equal(t, list.GetHead().next.next.next.data, "d")
	assert.Equal(t, list.GetHead().next.next.next.next.data, "e")
	assert.Equal(t, list.GetHead().next.next.next.next.next.data, "f")
}

func TestLinkedList_AddByIndexToHead(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.AddByIndex(0, "a")
	assert.Equal(t, list.GetHead().data, "a")

	list.AddByIndex(0, "b")
	assert.Equal(t, list.GetHead().data, "b")

	list.AddByIndex(0, "c")
	assert.Equal(t, list.GetHead().data, "c")
}

func TestLinkedList_AddByIndexToMiddle(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	err := list.AddByIndex(1, "d")
	assert.Nil(t, err)
	assert.Equal(t, list.GetHead().next.data, "d")

	err = list.AddByIndex(1, "e")
	assert.Nil(t, err)
	assert.Equal(t, list.GetHead().next.data, "e")
	assert.Equal(t, list.GetHead().next.next.data, "d")
}

func TestLinkedList_AddByIndexReturnErrorWhenIndexOutOfBounds(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	err := list.AddByIndex(12, "d")
	assert.NotNil(t, err)
}
func TestLinkedList_AddByIndexToEndOfList(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	err := list.AddByIndex(3, "d")
	assert.Nil(t, err)
	assert.Equal(t, list.GetHead().next.next.next.data, "d")
}

func TestLinkedList_Read(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	read, err := list.Read(0)
	assert.Nil(t, err)
	assert.Equal(t, "a", read)
	read, err = list.Read(1)
	assert.Nil(t, err)
	assert.Equal(t, "b", read)
	read, err = list.Read(2)
	assert.Nil(t, err)
	assert.Equal(t, "c", read)
}

func TestLinkedList_ReadLastItem(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	read, err := list.ReadLastItem()
	assert.Nil(t, err)
	assert.Equal(t, "c", read)
}

func TestLinkedList_Search(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	idx, err := list.Search("a")
	assert.Nil(t, err)
	assert.Equal(t, 0, idx)
	idx, err = list.Search("b")
	assert.Nil(t, err)
	assert.Equal(t, 1, idx)
	idx, err = list.Search("c")
	assert.Nil(t, err)
	assert.Equal(t, 2, idx)
}

func TestLinkedList_SearchErrorWhenIndexNotFound(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	idx, err := list.Search("z")
	assert.NotNil(t, err)
	assert.Equal(t, -1, idx)
}

func TestLinkedList_DeleteHead(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	err := list.Delete(0)
	assert.Nil(t, err)
	assert.Equal(t, "b", list.GetHead().data)
	err = list.Delete(0)
	assert.Nil(t, err)
	assert.Equal(t, "c", list.GetHead().data)
}

func TestLinkedList_DeleteMiddle(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	err := list.Delete(1)
	assert.Nil(t, err)
	assert.Equal(t, "a", list.GetHead().data)
	assert.Equal(t, "c", list.GetHead().next.data)
}

func TestLinkedList_DeleteEnd(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	err := list.Delete(2)
	assert.Nil(t, err)
	assert.Equal(t, "a", list.GetHead().data)
	assert.Equal(t, "b", list.GetHead().next.data)
	err = list.Delete(0)
	assert.Nil(t, err)
	assert.Equal(t, "b", list.GetHead().data)
}

func TestLinkedList_DeleteErrorWhenIndexNotFound(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	err := list.Delete(10)
	assert.NotNil(t, err)
}

func TestLinkedList_DeleteItems(t *testing.T) {
	list := NewClassicLinkedList[string]()

	list.Add("this is a password:1 that needs obfuscating")
	list.Add("this is a password:12 that needs obfuscating")
	list.Add("this is a password:123 that needs obfuscating")
	list.Add("this is a password:12cae that needs obfuscating")
	list.Add("this is a password:some_password that needs obfuscating")
	list.Add("a")
	list.Add("b")

	predicate := func(data string) bool {
		pattern := regexp.MustCompile("\\bpassword:[\\w]+\\b")
		return pattern.MatchString(data)
	}
	list.DeleteItems(predicate)
	first, _ := list.Read(0)
	assert.Equal(t, first, "a")
	second, _ := list.Read(1)
	assert.Equal(t, "b", second)
}

func TestLinkedList_PrintItems(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	assert.Equal(t, "a -> b -> c -> nil", list.PrintItems())
}

func TestLinkedList_PrintItemsInReverse(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	assert.Equal(t, "c -> b -> a -> nil", list.PrintItemsInReverse())
}

func TestLinkedList_Reverse(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	list.Reverse()
	assert.Equal(t, "c", list.GetHead().data)
	assert.Equal(t, "b", list.GetHead().next.data)
	assert.Equal(t, "a", list.GetHead().next.next.data)
	read, err := list.Read(0)
	assert.Nil(t, err)
	assert.Equal(t, "c", read)
	read, err = list.Read(1)
	assert.Nil(t, err)
	assert.Equal(t, "b", read)
	read, err = list.Read(2)
	assert.Nil(t, err)
	assert.Equal(t, "a", read)
}

func TestLinkedList_GetNode(t *testing.T) {
	list := NewClassicLinkedList[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	res, _ := list.GetNode(0)
	assert.Equal(t, res.data, "a")
	res, _ = list.GetNode(1)
	assert.Equal(t, res.data, "b")
	res, _ = list.GetNode(2)
	assert.Equal(t, res.data, "c")
}

func TestLinkedList_DeleteStartNode(t *testing.T) {

}

func TestLinkedList_DeleteMidNode(t *testing.T) {

}
func TestLinkedList_DeleteEndNode(t *testing.T) {

}

/*


   @Test
   public void shouldDeleteStartNodeFromLinkedList() {
       ClassicLinkedList<String> classicLinkedList = new ClassicLinkedList<>();
       classicLinkedList.add("a");
       classicLinkedList.add("b");
       classicLinkedList.add("c");
       classicLinkedList.add("d");
       classicLinkedList.add("e");
       final Node<String> start = classicLinkedList.getNode(0);
       classicLinkedList.delete(start);


       MatcherAssert.assertThat(classicLinkedList.getHead().data, Matchers.equalTo("b"));
       MatcherAssert.assertThat(classicLinkedList.read(0), Matchers.equalTo("b"));
       MatcherAssert.assertThat(classicLinkedList.getHead().next.data, Matchers.equalTo("c"));
       MatcherAssert.assertThat(classicLinkedList.read(1), Matchers.equalTo("c"));
       MatcherAssert.assertThat(classicLinkedList.getHead().next.next.data, Matchers.equalTo("d"));
       MatcherAssert.assertThat(classicLinkedList.read(2), Matchers.equalTo("d"));
       MatcherAssert.assertThat(classicLinkedList.getHead().next.next.next.data, Matchers.equalTo("e"));
       MatcherAssert.assertThat(classicLinkedList.read(3), Matchers.equalTo("e"));
   }

   @Test
   public void shouldDeleteMidNodeFromLinkedList() {
       ClassicLinkedList<String> classicLinkedList = new ClassicLinkedList<>();
       classicLinkedList.add("a");
       classicLinkedList.add("b");
       classicLinkedList.add("c");
       classicLinkedList.add("d");
       classicLinkedList.add("e");
       final Node<String> mid = classicLinkedList.getNode(1);
       classicLinkedList.delete(mid);


       MatcherAssert.assertThat(classicLinkedList.getHead().data, Matchers.equalTo("a"));
       MatcherAssert.assertThat(classicLinkedList.read(0), Matchers.equalTo("a"));
       MatcherAssert.assertThat(classicLinkedList.getHead().next.data, Matchers.equalTo("c"));
       MatcherAssert.assertThat(classicLinkedList.read(1), Matchers.equalTo("c"));
       MatcherAssert.assertThat(classicLinkedList.getHead().next.next.data, Matchers.equalTo("d"));
       MatcherAssert.assertThat(classicLinkedList.read(2), Matchers.equalTo("d"));
       MatcherAssert.assertThat(classicLinkedList.getHead().next.next.next.data, Matchers.equalTo("e"));
       MatcherAssert.assertThat(classicLinkedList.read(3), Matchers.equalTo("e"));
   }

   @Test
   public void shouldDeleteEndNodeFromLinkedList() {
       ClassicLinkedList<String> classicLinkedList = new ClassicLinkedList<>();
       classicLinkedList.add("a");
       classicLinkedList.add("b");
       classicLinkedList.add("c");
       classicLinkedList.add("d");
       classicLinkedList.add("e");
       final Node<String> end = classicLinkedList.getNode(4);
       classicLinkedList.delete(end);


       MatcherAssert.assertThat(classicLinkedList.getHead().data, Matchers.equalTo("a"));
       MatcherAssert.assertThat(classicLinkedList.read(0), Matchers.equalTo("a"));
       MatcherAssert.assertThat(classicLinkedList.getHead().next.data, Matchers.equalTo("b"));
       MatcherAssert.assertThat(classicLinkedList.read(1), Matchers.equalTo("b"));
       MatcherAssert.assertThat(classicLinkedList.getHead().next.next.data, Matchers.equalTo("c"));
       MatcherAssert.assertThat(classicLinkedList.read(2), Matchers.equalTo("c"));
       MatcherAssert.assertThat(classicLinkedList.getHead().next.next.next.data, Matchers.equalTo("d"));
       MatcherAssert.assertThat(classicLinkedList.read(3), Matchers.equalTo("d"));
   }
*/
