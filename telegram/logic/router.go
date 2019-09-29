package logic

type router struct {
	cmds map[string]func()
}

func initRouter() router {
	return router{
		cmds: make(map[string]func()),
	}
}

func (r *router) on(cmd string, f func()) {
	r.cmds["/"+cmd] = f
}

func (r *router) route(args []string, f func()) {
	cmd := args[0]

	handler, ok := r.cmds[cmd]
	if !ok {
		f()
	} else {
		handler()
	}
}
