package scan

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/ice-bergtech/dnh/src/internal/model_ent"
)

type ScanInput struct {
	Domain string
	IP     string
}

func NewScan(input string, ctx context.Context, client *model_ent.Client) (*model_ent.Scan, error) {
	Scanid := base64.StdEncoding.EncodeToString([]byte(input))

	scanbuilder, err := client.Scan.Create().
		SetScanid(Scanid).
		SetInput(input).
		SetTimestamp(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating scan: %w", err)
	}
	return scanbuilder, nil
}
