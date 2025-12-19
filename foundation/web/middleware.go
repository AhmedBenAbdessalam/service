package web

type MidHandler func(Handler) Handler

func wrapMiddleware(mw []MidHandler, h Handler) Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		mwFunc := mw[i]
		if mwFunc != nil {
			h = mwFunc(h)
		}
	}
	return h
}