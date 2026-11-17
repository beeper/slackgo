package slack

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestNilClientMethods(t *testing.T) {
	skip := map[string]bool{
		// NewRTM copies the client by value, so there's nothing useful it could return.
		"NewRTM": true,
	}
	ctxType := reflect.TypeFor[context.Context]()
	errType := reflect.TypeFor[error]()
	var api *Client
	v := reflect.ValueOf(api)
	typ := v.Type()
	for i := range typ.NumMethod() {
		m := typ.Method(i)
		if skip[m.Name] {
			continue
		}
		t.Run(m.Name, func(t *testing.T) {
			mt := m.Type
			args := make([]reflect.Value, 0, mt.NumIn()-1)
			for j := 1; j < mt.NumIn(); j++ {
				in := mt.In(j)
				if mt.IsVariadic() && j == mt.NumIn()-1 {
					break
				} else if in == ctxType {
					args = append(args, reflect.ValueOf(context.Background()))
				} else {
					args = append(args, reflect.Zero(in))
				}
			}
			out := v.Method(i).Call(args)
			if mt.NumOut() > 0 && mt.Out(mt.NumOut()-1) == errType {
				err, _ := out[len(out)-1].Interface().(error)
				if !errors.Is(err, ErrNilClient) {
					t.Errorf("expected ErrNilClient, got %v", err)
				}
			}
		})
	}
}
