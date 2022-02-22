package pipe

type Pipe struct {
	onErr func(error)
	funcs []func() error
}

func New() *Pipe {
	return &Pipe{
		onErr: func(e error) { /* do nothing */ },
	}
}

func (p *Pipe) OnErr(handler func(error)) {
	p.onErr = handler
}

func (p *Pipe) Next(f func() error) *Pipe {
	p.funcs = append(p.funcs, f)
	return p
}

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
