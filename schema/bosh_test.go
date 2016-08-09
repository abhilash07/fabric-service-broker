package schema_test

import (
	"testing"

	"github.com/atulkc/fabric-service-broker/schema"

	. "gopkg.in/go-playground/assert.v1"
)

const (
	boshStemcell = "mystemcell"
	boshUuid     = "uuid-1"
	vmType       = "vmtype"
	networkNames = "net1,net2,net3"
)

func TestNewBoshDetails(t *testing.T) {
	boshDetails := schema.NewBoshDetails(boshStemcell, boshUuid, vmType, networkNames)
	NotEqual(t, boshDetails, nil)
	Equal(t, boshDetails.StemcellName, boshStemcell)
	Equal(t, boshDetails.DirectorUUID, boshUuid)
	Equal(t, boshDetails.Vmtype, vmType)
	Equal(t, len(boshDetails.NetworkNames), 3)
	err := boshDetails.Validate()
	Equal(t, err, nil)
}

func TestBoshDetailsValidate_Stemcell(t *testing.T) {
	boshDetails := schema.NewBoshDetails("", boshUuid, vmType, networkNames)
	NotEqual(t, boshDetails, nil)
	err := boshDetails.Validate()
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "StemcellName cannot be empty")
}

func TestBoshDetailsValidate_UUID(t *testing.T) {
	boshDetails := schema.NewBoshDetails(boshStemcell, "", vmType, networkNames)
	NotEqual(t, boshDetails, nil)
	err := boshDetails.Validate()
	NotEqual(t, boshDetails, nil)
	Equal(t, err.Error(), "DirectorUUID cannot be empty")
}

func TestBoshDetailsValidate_VmType(t *testing.T) {
	boshDetails := schema.NewBoshDetails(boshStemcell, boshUuid, "", networkNames)
	NotEqual(t, boshDetails, nil)
	err := boshDetails.Validate()
	NotEqual(t, boshDetails, nil)
	Equal(t, err.Error(), "Vmtype cannot be empty")
}

func TestBoshDetailsValidate_NetworkNames(t *testing.T) {
	boshDetails := schema.NewBoshDetails(boshStemcell, boshUuid, vmType, "")
	NotEqual(t, boshDetails, nil)
	err := boshDetails.Validate()
	NotEqual(t, boshDetails, nil)
	Equal(t, err.Error(), "Invalid network name in the list")
}
