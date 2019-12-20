// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "b99008927225e4deca4ed16493d52dac"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"a97480b591c17bffb5f753e1e8282d74": "1f8b08000000000000ffe2ffdfccc2c8f4bf85819187f17f130303dfff4610bf09c4e76160f8bf0a24c556965a949b98c7c897929a949998a79b9c9f5752949ff3bf372031393b313dd54aa1ba5a41cf2f313755a1b6962b2cb5a838333f0f2208e580c47d1333f34a1233f3528b2052083e48d6b1283923b32435b9a4b4086a1eb20848854b6a717251664109dc6c2401903c200000ffffddaff018cb000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("../assets", "../assets")
		b.SetResolver("templates/template.bundle", packr.Pointer{ForwardBox: gk, ForwardPath: "a97480b591c17bffb5f753e1e8282d74"})
	}()
	return nil
}()
