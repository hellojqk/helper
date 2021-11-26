package util

// WaitInitFuncs .
var WaitInitFuncs []func() error

func init() {
	WaitInitFuncs = make([]func() error, 0, 2)
}

// WaitInitFuncsAdd .
func WaitInitFuncsAdd(f func() error) {
	WaitInitFuncs = append(WaitInitFuncs, f)
}

// WaitInitFuncsExec .
func WaitInitFuncsExec() {
	for _, f := range WaitInitFuncs {
		err := f()
		if err != nil {
			panic(err)
		}
	}
}
