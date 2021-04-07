package validator

import (
	"fmt"
	"os"
	"testing"

	"github.com/microlib/logger/pkg/multi"
)

func TestEnvars(t *testing.T) {
	logger := multi.NewLogger(multi.HTML, multi.TRACE)

	t.Run("ValidateEnvars : should fail", func(t *testing.T) {
		os.Setenv("SERVER_PORT", "")
		err := ValidateEnvars(logger)
		if err == nil {
			t.Errorf(fmt.Sprintf("Handler %s returned with no error - got (%v) wanted (%v)", "ValidateEnvars", err, nil))
		}
	})

	t.Run("ValidateEnvars : should pass", func(t *testing.T) {
		os.Setenv("LOG_LEVEL", "info")
		os.Setenv("VERSION", "1.0.3")
		os.Setenv("NAME", "test")
		os.Setenv("SERVER_PORT", "9000")
		os.Setenv("TEMPLATE_FILE", "test")
		os.Setenv("CICD_CONSOLE_DIR", "test")
		err := ValidateEnvars(logger)
		if err != nil {
			t.Errorf(fmt.Sprintf("Handler %s returned with error - got (%v) wanted (%v)", "ValidateEnvars", err, nil))
		}
	})
}
