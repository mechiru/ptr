package ptr

import (
	"reflect"
	"testing"
)

func TestRef(t *testing.T) {
	v := true
	want := &v
	got := RefBool(true)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%+v, want=%+v", got, want)
	}
}

func TestDeref(t *testing.T) {
	for i, c := range []struct {
		in        *string
		def, want string
	}{
		{RefString("hoge"), "fuga", "hoge"},
		{(*string)(nil), "fuga", "fuga"},
	} {
		got := DerefString(c.in, c.def)
		if got != c.want {
			t.Errorf("index=%d: got=%+v, want=%+v", i, got, c.want)
		}
	}
}
