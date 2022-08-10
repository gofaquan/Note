package Interface

const (
	SvcName = "HelloSvc"
)

type HelloSvcInterface interface {
	Hello(request string, response *string) error
}
