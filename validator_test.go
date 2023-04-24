package validator_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/miraddo/validator"
)

func TestValidator(t *testing.T) {
	t.Run("IsEmail", func(t *testing.T) {
		assert.False(t, validator.IsEmail(""))                       // empty string
		assert.False(t, validator.IsEmail("http://www.google.com"))  // not email
		assert.False(t, validator.IsEmail("http://email@gmail.com")) // not email
		assert.True(t, validator.IsEmail("email@test.com"))          // email
	})

	t.Run("IsURL", func(t *testing.T) {
		assert.False(t, validator.IsURL(""))                           // empty string
		assert.True(t, validator.IsURL("http://www.google.com"))       // url
		assert.True(t, validator.IsURL("http://"+time.Now().String())) // not url

	})

	t.Run("IsURLStatusOK", func(t *testing.T) {
		ok, err := validator.IsURLStatusOK("http://www.google.com")
		require.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("IsImage", func(t *testing.T) {
		ok, err := validator.IsImage("https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png")
		require.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("IsImageLocaly", func(t *testing.T) {
		ok, err := validator.IsImageLocaly("https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png")
		require.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("IsImageURLStatus", func(t *testing.T) {
		ok, err := validator.IsImageURLStatus("https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png")
		require.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("IsImageUrl", func(t *testing.T) {
		assert.False(t, validator.IsImageUrl("")) // empty string
		assert.True(t, validator.IsImageUrl("https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"))
	})

}
