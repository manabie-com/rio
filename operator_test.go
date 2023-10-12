package rio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OpUrlEqualTo(t *testing.T) {
	t.Run("should be prefixed with echo", func(t *testing.T) {
		// arrange
		givenURL := "/api/eureka/get_token"
		expectedURL := "/echo/api/eureka/get_token"

		// act
		op := OpUrlEqualTo(givenURL)

		// assert
		assert.Equal(t, expectedURL, op.Value)
		assert.Equal(t, OpEqualTo, op.Name)
	})

	t.Run("url should be prefixed with echo and added slash", func(t *testing.T) {
		// arrange
		givenURL := "api/eureka/get_token"
		expectedURL := "/echo/api/eureka/get_token"

		// act
		op := OpUrlEqualTo(givenURL)

		// assert
		assert.Equal(t, expectedURL, op.Value)
		assert.Equal(t, OpEqualTo, op.Name)
	})
}
