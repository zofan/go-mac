package mac

import (
	"fmt"
	"github.com/zofan/go-fwrite"
	"github.com/zofan/go-req"
	"github.com/zofan/go-xmlre"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func UpdatePrefixes() error {
	var (
		httpClient = req.New(req.DefaultConfig)
		list       = make(map[string]*Prefix)
	)

	resp := httpClient.Get(`https://www.macvendorlookup.com/vendormacs-xml-download`)
	if resp.Error() != nil {
		return resp.Error()
	}

	body := string(resp.ReadAll())
	body = strings.ReplaceAll(body, `&nbsp;`, ` `)

	rowRe := xmlre.Compile(`<VendorMapping mac_prefix="([^"]+)" vendor_name="([^"]+)"/>`)

	for _, row := range rowRe.FindAllStringSubmatch(body, -1) {
		p := &Prefix{
			Prefix: strings.ToUpper(strings.TrimSpace(row[1])),
			Name:   strings.TrimSpace(row[2]),
		}

		/*for _, w := range strings.Split(strings.ToLower(p.Name), ` `) {
			p.Tags = append(p.Tags, w)
		}*/

		list[p.Prefix] = p
	}

	// ---

	var tpl []string

	tpl = append(tpl, `package mac`)
	tpl = append(tpl, ``)
	tpl = append(tpl, `// Updated at: `+time.Now().String())
	tpl = append(tpl, `var Prefixes = []Prefix{`)

	for _, l := range list {
		tpl = append(tpl, `	{`)
		tpl = append(tpl, `		Prefix: "`+l.Prefix+`",`)
		tpl = append(tpl, `		Name:   "`+l.Name+`",`)
		tpl = append(tpl, `		Tags:   `+fmt.Sprintf(`%#v`, l.Tags)+`,`)
		tpl = append(tpl, `	},`)
	}

	tpl = append(tpl, `}`)
	tpl = append(tpl, ``)

	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)

	return fwrite.WriteRaw(dir+`/prefix_db.go`, []byte(strings.Join(tpl, "\n")))
}
