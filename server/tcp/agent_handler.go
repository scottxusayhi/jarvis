package tcp

type handler interface {
	Handle (c []byte) error
	SendWelcome(a *JarvisAgent) error

}
