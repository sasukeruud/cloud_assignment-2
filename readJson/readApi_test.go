package readjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCasesApi(t *testing.T) {
	jsonData := ReadCasesApi("Norway")

	jsonString := string(jsonData)

	assert.Equal(t, jsonString, `{"data":{"country":{"name":"Norway","mostRecent":{"date":"2022-04-04","confirmed":1410051,"deaths":2518,"recovered":0,"growthRate":0.0009533558409549743}}}}
`) //Needs to look like this to get the same type of structur as the response jsonString
}

func TestReadPolicyApi(t *testing.T) {
	policyInfo := ReadPolicyApi("NOR", "2021-01-01")

	assert.Equal(t, len(policyInfo), 1)
}
