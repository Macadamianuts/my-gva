package internal

type server interface {
	ListenAndServe() error
}
