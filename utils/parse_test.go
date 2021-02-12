package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// ParseTestSuite tests functionality of the parse logic
type ParseTestSuite struct {
	suite.Suite
}

// TestParse tests that the parse function works
func (ts *ParseTestSuite) TestParse() {
	// test for non secrets
	expected := map[string]string{
		"TEST_REGULAR_ENV_VAR": "non-secret-env-var",
	}
	envMap, err := ParseEnvFile("../sample/non-secret.jsonc")
	ts.Require().NoError(err)
	ts.Require().Equal(expected, *envMap)

	// Test for secrets
	expected = map[string]string{
		"TEST_SECRET_ENV_VAR": "arn:aws:secretsmanager:us-east-2:349142687622:secret:dev/TEST_SECRET_FOR_TESTING-96qtUg",
	}
	envMap, err = ParseEnvFile("../sample/secret.jsonc")
	ts.Require().NoError(err)
	ts.Require().Equal(expected, *envMap)

}

func TestParseTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping parse test in short mode.")
	}
	suite.Run(t, new(ParseTestSuite))
}
