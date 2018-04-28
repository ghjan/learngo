package rpcdemo

import "errors"

//Service.Method

type DemoService struct{}

type Args struct {
	A, B int
}

//{"method":"abc.ef"}
//{"method":"DemoService.Div", "params":[{"A":3, "B":4}], "id":1}
//{"method":"DemoService.Div", "params":[{"A":3, "B":0}], "id":1234}
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
