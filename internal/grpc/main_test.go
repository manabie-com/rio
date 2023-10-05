package grpc

import (
	"context"
	"os"
	"testing"

	"github.com/manabie-com/rio/internal/test"
)

func TestMain(m *testing.M) {
	test.ResetDB(context.Background(), "../..")

	code := m.Run()
	os.Exit(code)
}
