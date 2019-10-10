package pkg

const (
	unExported = 15
	Exported   = 10
)

func private() string {
	return "some public string"
}

func Public() string {
	return "some public string"
}
