package frotel

import (
	"fmt"
	"github.com/Ryanair/gofrlib/log"
	"testing"
)

func TestStart(t *testing.T) {
	type args struct {
		handlerFunc interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestStart",
			args: args{
				handlerFunc: func() {
					fmt.Printf("I am working")
				},
			},
		},
		{
			name: "TestPanic",
			args: args{
				handlerFunc: func() {
					someVariable := struct {
						a *struct {
							b string
						}
					}{a: nil}

					fmt.Printf(fmt.Sprintf("I am working on: %s", someVariable.a.b))
				},
			},
		},
	}
	for _, tt := range tests {
		log.Init(log.NewConfiguration("DEBUG", "TEST-APPLICATION", "TEST-PROJECT", "TEST-PROJECT-GROUP", "1.0.0", "testPrefix"))
		t.Run(tt.name, func(t *testing.T) {
			Start(tt.args.handlerFunc)
		})
	}
}
