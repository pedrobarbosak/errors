package service

import "github.com/pedrobarbosak/errors"

func Error() error {
	return errors.New("some service failed because value as:", 100)
}
