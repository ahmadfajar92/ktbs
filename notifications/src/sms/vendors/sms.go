package vendors

import (
	"math/rand"
	"notifications/src/shared"
	"time"

	"github.com/BurntSushi/toml"
)

var smsStatus = struct {
	Sent      string
	Delivered string
	Pending   string
}{
	Sent:      "SENT",
	Delivered: "DELIVERED",
	Pending:   "PENDING",
}

// sms vendor factory
type (
	SMSPayload struct {
		To      string `json:"to"`
		Message string `json:"message"`
	}

	SMSResponse struct {
		To      string `json:"to"`
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	smsvndrs struct {
		vendors map[string]shared.Vendor
	}

	SMSVendrsCfg struct {
		Configs map[string]SMSVendrCfg `toml:"smsvendors"`
	}

	SMSVendrCfg struct {
		Token     string
		Enable    string
		Uri       string
		IsDefault string
	}
)

func LoadFromFile(pathfile string) shared.Vendors {
	smsvndrs := new(smsvndrs)
	smsvndrs.vendors = make(map[string]shared.Vendor, 0)
	// load vendor config
	smsvndrs.loads(pathfile)

	return smsvndrs
}

func SMSVendors() shared.Vendors {
	smsvndrs := new(smsvndrs)
	smsvndrs.vendors = make(map[string]shared.Vendor, 0)

	return smsvndrs
}

func (vndrs *smsvndrs) loads(pathfile string) {
	vendors := make(map[string]func(cfg *SMSVendrCfg) shared.Vendor, 0)

	// register init method sms vendor
	vendors["bacrit"] = BacritSMSVendor
	vendors["ngehe"] = NgeheSMSVendor

	// // read toml file
	var cfg SMSVendrsCfg
	if _, err := toml.DecodeFile(pathfile, &cfg); err != nil {
		panic(err.Error())
	}

	countDefault := 0
	countEnable := 0
	for k, c := range cfg.Configs {
		dfault := c.IsDefault == "true"
		enable := c.Enable == "true"
		if dfault && countDefault == 1 {
			panic("number of default vendors must be one")
		}

		if dfault {
			countDefault += 1
		}

		if enable {
			countEnable += 1
		}

		vndrs.vendors[k] = vendors[k](&c)
	}

	if countDefault+countEnable == 0 {
		panic("there's no active vendors as default")
	}

	if countDefault == 0 {
		panic("there's no vendors as default")
	}
}

// Get(n string) (v Vendor)
func (vndrs *smsvndrs) Get(n string) (v shared.Vendor) {
	v = vndrs.vendors[n]
	return
}

// All() (v map[string]Vendor)
func (vndrs *smsvndrs) All() map[string]shared.Vendor {
	return vndrs.vendors
}

// ActiveVendor() (v Vendor)
func (vndrs *smsvndrs) ActiveVendor() (v shared.Vendor) {
	var actvndrs = make([]shared.Vendor, 0)
	for _, vndr := range vndrs.vendors {
		if vndr.IsOn() {
			actvndrs = append(actvndrs, vndr)
		}
	}
	// get random active vendor
	rand.Seed(time.Now().UnixNano())
	v = actvndrs[rand.Intn((len(actvndrs)-1)-0+1)+0]
	return
}
