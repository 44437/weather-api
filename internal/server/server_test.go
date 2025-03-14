package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestNewServer(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	var handlers []Handler
	srv := NewServer(8080, handlers)

	assert.True(t, srv.GetEchoInstance().HideBanner)
	assert.Equal(t, "GET", srv.GetEchoInstance().Routes()[0].Method)
	assert.Equal(t, "GET", srv.GetEchoInstance().Routes()[1].Method)
}

func TestStartStop(t *testing.T) {
	var handlers []Handler
	srv := NewServer(8080, handlers)

	time.AfterFunc(1*time.Second, func() {
		srv.Stop()
	})

	err := srv.Start()
	require.NotNil(t, err)
	assert.Equal(t, "http: Server closed", err.Error())
}
