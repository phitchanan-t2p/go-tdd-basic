package captcha

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_captcha "github.com/phitchanan-t2p/go-tdd-basic/captcha/mocks"
)

func TestCaptcha(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCaptchaService := mock_captcha.NewMockRandomizer(ctrl)
	captcha := NewRandGenerator(mockCaptchaService)

	mockCaptchaService.EXPECT().Random().Return(100000)

	got := captcha.Captcha()
	want := "T2P100000"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
