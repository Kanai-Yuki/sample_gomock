package sample

import (
	"testing"

	"github.com/Kanai-Yuki/sample_gomock/mock_sample"

	"github.com/golang/mock/gomock"
)

func TestMethod(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_sample.NewMockSample(ctrl)
	m.EXPECT().Method(1, 3).Return(4)

	m.Method(1, 3)
}
