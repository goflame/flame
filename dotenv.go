package flame

import (
	"github.com/golobby/dotenv"
	"os"
)

func DotEnv(s any, path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	err = dotenv.NewDecoder(file).Decode(s)

	if err != nil {
		return err
	}

	return nil
}
