package repository

import (
	model "case2/models"
	"errors"

	"github.com/stretchr/testify/mock"
)

type LogRepositoryMock struct {
	Mock mock.Mock
}

func (repository *LogRepositoryMock) Insert(log model.Logger) error {
	arguments := repository.Mock.Called(log)
	if arguments.Get(0) != nil {
		return nil
	} else {
		return errors.New("error")
	}
}
