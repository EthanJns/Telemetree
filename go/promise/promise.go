package promise

type Promise[T any] struct {
	OnComplete     func(func(t *T))
	OnFailure      func(func(s string))
	callOnComplete func(t *T)
	callOnFailure  func(s string)
	Complete       func(t *T)
	Fail           func(s string)
	IsCompleted    bool
}

func Create[T any]() *Promise[T] {
	p := Promise[T]{OnComplete: nil, OnSuccess: nil, OnFailure: nil, Complete: nil, IsCompleted: false}
	p.OnComplete = func(f func(t *T)) { onComplete(&p, f) }
	p.OnFailure = func(f func(s string)) { onFailure(&p, f) }
	p.Complete = func(t *T) { p.callOnComplete(t) }
	p.Fail = func(s string) { p.callOnFailure(s) }
	return &p
}

func onComplete[T any](p *Promise[T], callOnComplete func(t *T)) {
	p.callOnComplete = callOnComplete
}

func onFailure[T any](p *Promise[T], callOnFailure func(s string)) {
	p.callOnFailure = callOnFailure
}
