package scan

import (
	"context"
	"net/url"
	"strconv"

	"github.com/ice-bergtech/dnh/src/internal/model_ent"
)

func checkType(i string) {

}

func ProcessDomain(domain *model_ent.Scan, ctx context.Context, client *model_ent.Client) (*model_ent.Domain, error) {
	// <scheme>:[//[<login>[:<password>]@]<host>[:<port>]][/<URL path>][?<parameters>][#<anchor>]
	url, err := url.Parse(domain.Input)
	if err != nil {
		return nil, err
	}
	ports := []int{}

	if url.Port() != "" {
		portNum, err := strconv.Atoi(url.Port())
		if err == nil {
			ports = append(ports, portNum)
		}
	}

	client.Domain.
		Create().
		SetName(url.Hostname()).
		SetPorts(ports).
		SetTimeFirst(domain.Timestamp).
		SetTimeLast(domain.Timestamp).
		OnConflict().
		Ignore()

	entry := &model_ent.Domain{
		Name:  url.Hostname(),
		Ports: ports,
	}
	return entry, nil
}
