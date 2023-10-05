package api

import (
	"context"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/rio/internal/test"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.ReleaseMode)
	test.ResetDB(context.Background(), "../..")

	code := m.Run()
	os.Exit(code)
}
