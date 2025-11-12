package persistance

type (
	Storage interface {
		StoreFile(path string, data []byte) error
		GetFile(path string) ([]byte, error)
		DeleteFile(path string) error
	}

	File struct {
		Data []byte
		Name string
		User string
	}
)
