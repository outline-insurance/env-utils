package utils

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/pkg/errors"
)

// GetSecret retrieves the value of the secret ARN given
func GetSecret(secretARN string, region string) (*string, error) {
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretARN),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return nil, errors.Wrapf(err, "while getting secret %s", secretARN)
	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	if result.SecretString != nil {
		return result.SecretString, nil
	}
	decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
	len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
	if err != nil {
		return nil, errors.Wrapf(err, "while decoding base 64 binary secret %s", secretARN)
	}
	decodedBinarySecret := string(decodedBinarySecretBytes[:len])
	return &decodedBinarySecret, nil
}
