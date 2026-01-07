package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVerifyOTPSuccess(t *testing.T) {

	SkipDBInsert = true
	defer func() { SkipDBInsert = false }()

	pendingRegistrationsMutex.Lock()
	pendingRegistrations["example@mail.com"] = PendingUser{
		Username:  "testuser",
		Email:     "example@mail.com",
		Password:  "securepassword",
		OTP:       "123456",
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	pendingRegistrationsMutex.Unlock()

	reqBody := `{
        "email": "example@mail.com",
        "otp": "123456"
    }`

	req := httptest.NewRequest(http.MethodPost, "/api/verify-email-otp", strings.NewReader(reqBody))
	rr := httptest.NewRecorder()

	VerifyEmailOTPHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"message":"Verifikasi berhasil"}`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestVerifyOTPEmailNotFound(t *testing.T) {

	SkipDBInsert = true
	defer func() { SkipDBInsert = false }()

	pendingRegistrationsMutex.Lock()
	pendingRegistrations = make(map[string]PendingUser)
	pendingRegistrationsMutex.Unlock()

	reqBody := `{
        "email": "tidakada@mail.com",
        "otp": "123456"
    }`

	req := httptest.NewRequest(http.MethodPost, "/api/verify-email-otp", strings.NewReader(reqBody))
	rr := httptest.NewRecorder()

	VerifyEmailOTPHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	assert.Equal(t, "Kode verifikasi tidak valid / kadaluarsa\n", rr.Body.String())
}

func TestVerifyOTPInvalidOTP(t *testing.T) {

	SkipDBInsert = true
	defer func() { SkipDBInsert = false }()

	pendingRegistrationsMutex.Lock()
	pendingRegistrations["example@mail.com"] = PendingUser{
		Username:  "testuser",
		Email:     "example@mail.com",
		Password:  "securepassword",
		OTP:       "123456",
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	pendingRegistrationsMutex.Unlock()

	reqBody := `{
        "email": "example@mail.com",
        "otp": "999999"
    }`

	req := httptest.NewRequest(http.MethodPost, "/api/verify-email-otp", strings.NewReader(reqBody))
	rr := httptest.NewRecorder()

	VerifyEmailOTPHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	assert.Equal(t, "Kode OTP salah\n", rr.Body.String())
}
