package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello SayHello
}

type SayHelloImpl struct {
}

func (sayHello SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func NewSayHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}
