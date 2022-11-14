package impl

import (
	"context"
	"fmt"

	"github.com/summingyu/restful-api-demo/apps/host"
)

func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error: %v", err)
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback error: %v", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error: %v", err)
			}
		}
	}()

	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()

	_, err = rstmt.ExecContext(ctx,
		ins.Account,
		ins.CreateAt,
		ins.Description,
		ins.ExpireAt,
		ins.Id,
		ins.Name,
		ins.PrivateIP,
		ins.PublicIP,
		ins.Region,
		ins.Status,
		ins.SyncAt,
		ins.Type,
		ins.UpdateAt,
		ins.Vendor,
	)

	return nil
}
