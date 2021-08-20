package shared

import (
	"context"
)

type (
	Usecases interface {
		Call(n string) (u interface{})
		Add(n string, u interface{})
	}

	// Vendor interface
	Vendor interface {
		Send(ctx context.Context, p interface{}) (r interface{}, err error)
		OnOff(ctx context.Context, s bool) (err error)
		IsOn() (b bool)
		IsDefault() (b bool)
	}

	// Vendors interface
	Vendors interface {
		Get(n string) (v Vendor)
		All() (v map[string]Vendor)
		ActiveVendor() (v Vendor)
	}
)
