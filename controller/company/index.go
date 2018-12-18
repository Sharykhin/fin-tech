package company

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
)

type listResult struct {
	items []entity.Company
	err   error
}

type totalResult struct {
	total int64
	err   error
}

func (c CompanyController) Index(ctx context.Context, limit, offset int64) ([]entity.Company, int64, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	listResultCh := c.getListChan(ctx, limit, offset)
	totalResultCh := c.getTotalChan(ctx)
	var list []entity.Company
	var total int64

	for {
		if listResultCh == nil && totalResultCh == nil {
			break
		}
		select {
		case <-ctx.Done():
			return nil, 0, ctx.Err()
		case lr, ok := <-listResultCh:
			if !ok {
				listResultCh = nil
				continue
			}
			if lr.err != nil {
				return nil, 0, lr.err
			}
			list = lr.items

		case tr, ok := <-totalResultCh:
			if !ok {
				totalResultCh = nil
				continue
			}
			if tr.err != nil {
				return nil, 0, ctx.Err()
			}

			total = tr.total
		}
	}

	return list, total, nil
}

func (c CompanyController) getListChan(ctx context.Context, limit, offset int64) <-chan listResult {
	ch := make(chan listResult)
	go func() {
		defer close(ch)
		list, err := c.storage.List(ctx, limit, offset)
		ch <- listResult{
			items: list,
			err:   err,
		}
	}()
	return ch
}

func (c CompanyController) getTotalChan(ctx context.Context) <-chan totalResult {
	ch := make(chan totalResult)
	go func() {
		defer close(ch)
		total, err := c.storage.Count(ctx)
		ch <- totalResult{
			total: total,
			err:   err,
		}
	}()
	return ch
}
