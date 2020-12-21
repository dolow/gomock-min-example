package service

import (
	"gomock_example/dao"
	mock_dao "gomock_example/mock/dao"

	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMockWithSignup(t *testing.T) {
	// preparing mock and mock postprocess
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// input and output
	var user   dao.User
	var expect error

	t.Run("using UserImpl", func(t *testing.T) {
		user = &dao.UserImpl{}
		expect = nil

		t.Run("should return nil error", func(t *testing.T) {
			assert.Equal(t, expect, Signup(user))
		})
	})

	t.Run("using mock", func(t *testing.T) {
		expect = errors.New("error for test")
		// *MockUser returned here implements User interface
		user = mock_dao.NewMockUser(ctrl)

		t.Run("when mock returns error", func(t *testing.T) {
			user.(*mock_dao.MockUser).EXPECT().Create().Return(expect)

			t.Run("should return error", func(t *testing.T) {
				assert.Equal(t, expect, Signup(user))
			})
		})
	})
}
