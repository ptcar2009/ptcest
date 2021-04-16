package pkg

import (
	"reflect"

	"github.com/google/go-cmp/cmp"
)

func (t *T) Equals(one, other interface{}, options ...cmp.Option) {
	if !cmp.Equal(one, other, options...) {
		t.Errorf("assert failure \n%#v\nis not equal to\n%#v\n", one, other)
	}
}

func (t *T) Unequals(one, other interface{}, options ...cmp.Option) {
	if cmp.Equal(one, other, options...) {
		t.Errorf("assert failure \n%#v\nis not equal to\n%#v\n", one, other)
	}
}

func (t *T) IsNil(a interface{}) {
	if !(a == nil) && !reflect.ValueOf(a).IsNil() {
		t.Errorf("assert failure %#v is not nil", a)
	}
}

func (t *T) IsNotNil(a interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Errorf("assert failure %#v is nil", a)
	}
}
