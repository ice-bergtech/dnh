package scan

import (
	"context"
	"net/url"
	"strconv"

	"github.com/ice-bergtech/dnh/src/internal/model_ent"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/domain"
)

func checkType(i string) {

}

func ProcessDomain(dmn *model_ent.Scan, ctx context.Context, client *model_ent.Client) (*model_ent.Domain, error) {
	// <scheme>:[//[<login>[:<password>]@]<host>[:<port>]][/<URL path>][?<parameters>][#<anchor>]
	url, err := url.Parse(dmn.Input)
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

	val, err := client.Domain.Query().Where(domain.Name(url.Hostname())).First(ctx)
	if err == nil {
		val, err = val.Update().SetTimeLast(dmn.Timestamp).AppendPorts(ports).Save(ctx)
	} else {
		val, err = client.Domain.Create().
			SetName(url.Hostname()).
			SetPorts(ports).
			SetTimeFirst(dmn.Timestamp).
			SetTimeLast(dmn.Timestamp).Save(ctx)
	}

	// Queue check functions for domain
	return val, err
}
