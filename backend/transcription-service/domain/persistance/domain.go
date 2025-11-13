package persistance

type (
	Storage interface {
		StoreFile(path string, data []byte) error
		GetFile(path string) ([]byte, error)
		DeleteFile(path string) error
	}
)
