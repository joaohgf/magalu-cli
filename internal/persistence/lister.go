package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	clierror "github.com/joaohgf/magalu-cli/internal/errors"
	"github.com/joaohgf/magalu-cli/internal/port"
	"github.com/nanobox-io/scribble"
)

type ServiceLister[D port.Domain, T port.FilterDomain[D]] struct {
	*scribble.Driver
}

func NewServiceLister[D port.Domain, T port.FilterDomain[D]](db *scribble.Driver) *ServiceLister[D, T] {
	return &ServiceLister[D, T]{Driver: db}
}

// List retrieves a list of items from the collection based on the filter criteria defined in the target.
func (s *ServiceLister[D, T]) List(_ context.Context, target T) (T, error) {
	out, err := s.readAll(target)
	if err != nil {
		return target, err
	}
	target.SetContent(out...)
	target.SetTotal(len(out))
	return target, nil
}

// readAll reads all items from the collection and filters them based on the target's IsEqual method.
// It returns a slice of items that match the filter criteria and an error if any occurs during the process.
func (s *ServiceLister[D, T]) readAll(target T) ([]D, error) {
	data, err := s.Driver.ReadAll(target.GetCollection())
	if err != nil {
		if _, ok := errors.AsType[*os.PathError](err); ok {
			return nil, nil
		}
		return nil, clierror.NotFound(fmt.Sprintf("no items found in collection %s", target.GetCollection()))
	}
	var out []D
	for _, raw := range data {
		// if the length of the output slice is greater than or equal to the total number of items that can be displayed on the current page,
		// it means that we have already collected enough items to fill the current page, so we can stop processing further items
		if len(out) >= target.GetSize()*target.GetPage() {
			break
		}
		var item D
		if err = json.Unmarshal([]byte(raw), &item); err != nil {
			return nil, clierror.Invalid(fmt.Sprintf("failed to unmarshal %s", target.GetCollection()))
		}
		if !target.IsEqual(item) {
			continue
		}
		out = append(out, item)
	}
	out = s.paginate(out, target)
	return out, nil
}

func (s *ServiceLister[D, T]) paginate(items []D, target T) []D {
	// limit starts with the length of the filtered items, but it will be adjusted to fit the pagination limits
	limit := len(items)
	// if the limit is greater than the total number of items that can be displayed on the current page,
	// we need to adjust it to fit the pagination limits
	if limit > target.GetSize()*target.GetPage() {
		limit = target.GetSize() * target.GetPage()
	}
	// offSet starts with 0, but it will be adjusted to fit the pagination limits
	offSet := target.GetSize() * (target.GetPage() - 1)
	// if the offset is greater than or equal to the total number of items that can be displayed on the current page,
	// it means that there are no items to display on the current page, so we return an empty slice
	if target.GetSize()*(target.GetPage()-1) >= limit {
		return nil
	}
	items = items[offSet:limit]
	return items
}
