package boundaries

type Action interface {
	Execute(input interface{}, output interface{})
}
