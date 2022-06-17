package node

import (
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"github.com/decision2016/go-crypto/ecdcs"
	"github.com/decision2016/osc/internal/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

// Validator的测试代码
func TestValidator(t *testing.T) {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&utils.Formatter{})
	logrus.SetLevel(logrus.TraceLevel)
	validator := ChameleonValidator{}

	hexPrivate := "3f707e43bd8805f379b86afacc70f39767e9a3ca896553445d30d8c5ad089610"
	chameleonSec := ecdcs.PrivateKey{}
	err := chameleonSec.FromHexString(hexPrivate)
	logrus.Info(err)

	chameleonPub, _ := chameleonSec.ExportPublicKey(elliptic.P256())

	logrus.Info(chameleonPub.ToHexString())

	genesisData := utils.GenesisData{
		Id:             "",
		Random:         "",
		CreatedTx:      "",
		Version:        0,
		Operator:       "0x9a8489051f6fA66B3104920844B1ccc7C3F8eDbC",
		Timestamp:      1655298330,
		PermissionList: []string{"0x9a8489051f6fA66B3104920844B1ccc7C3F8eDbC"},
		BlockHash:      "0xb2658061b110ef6a22503d65464c990dbf948c7ba2e1f401157359a7fb1bfa50",
		PublicKey:      chameleonPub.ToHexString(),
	}

	genesisJson, _ := json.Marshal(genesisData)
	random, sign, _ := chameleonSec.Signature(elliptic.P256(), genesisJson)

	genesisData.Id = hex.EncodeToString(sign)
	genesisData.Random = hex.EncodeToString(random)
	genesisData.CreatedTx = "test"

	genesisJson, _ = json.Marshal(genesisData)

	t.Log(validator.Validate(genesisData.Id, genesisJson))
}
