package support

type HTTPAPICaller interface {
	Post(path string, payload interface{}) (response string, httpStatus int, err error)
	Get(path string) (response string, httpStatus int, err error)
	Verify(expectedState interface{})
}
