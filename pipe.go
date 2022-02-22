// Package pipe implements a Pipe.
package pipe

type Pipe struct {
	onErr func(error)
	funcs []func() error
}

// New constructs an empty Pipe.
func New() *Pipe {
	return &Pipe{
		onErr: func(e error) { /* do nothing */ },
	}
}

// OnErr is a convinience function to handle an error.
func (p *Pipe) OnErr(handler func(error)) {
	p.onErr = handler
}

// Next appends the provided funcion to the Pipe.
func (p *Pipe) Next(f func() error) *Pipe {
	p.funcs = append(p.funcs, f)
	return p
}

// Do executes the Pipe.
func (p *Pipe) Do() error {
	for _, f := range p.funcs {
		if err := f(); err != nil {
			// stop the chain prematurely
			p.onErr(err)
			return err
		}
	}

	p.onErr(nil)
	return nil
}
