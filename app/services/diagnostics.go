package services

import (
	"strings"

	"github.com/elephant-insurance/go-microservice-arch/v2/dig"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
	"github.com/elephant-insurance/ms-sites/app/cfg"
)

func TestContainer() (dig.DiagnosticResult, error) {
	c := msrqc.New(nil)
	blobService := NewBlobService(cfg.Config.StorageAccountName, cfg.Config.StorageAccountKey, "singlesearch")
	client, err := blobService.initializeClient(c)
	if err != nil {
		rp := dig.NewResult().Fail().SetDescription("Client initialization failed")
		return *rp, nil
	}
	blobService.client = client

	err = blobService.ensureContainer(c)

	res := strings.Contains(err.Error(), "ContainerAlreadyExists")
	if res {
		rp := dig.NewResult().Succeed().SetDescription("storage container test ok")
		return *rp, nil
	}
	rp := dig.NewResult().Fail().SetDescription("storage container test failed")
	rp.Error = err
	return *rp, nil
}
