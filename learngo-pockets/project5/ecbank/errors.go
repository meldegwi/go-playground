package ecbank

type ecbankError string

func (e ecbankError) String() string {
	return string(e)
}
