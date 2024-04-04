package rbxbin

import (
	"testing"

	cs "github.com/apprehensions/rbxweb/clientsettings"
)

func TestDirsCurrent(t *testing.T) {
	for _, b := range []cs.BinaryType{cs.WindowsPlayer, cs.WindowsStudio64} {
		d, err := GetDeployment(b, "")
		if err != nil {
			t.Skipf("%s: Deployment fetch failed: %s", b, err)
		}

		ps, err := DefaultMirror.GetPackages(d)
		if err != nil {
			t.Skipf("%s: Packages fetch failed: %s", b, err)
		}

		dirs := BinaryDirectories(b)

		for zip, _ := range dirs {
			found := false
			for _, p := range ps {
				if p.Name == zip {
					found = true
				}
			}
			if !found {
				t.Errorf("%s: dirs zip %s missing in upstream packages", b, zip)
			}
		}

		for _, p := range ps {
			if p.Name == "RobloxPlayerLauncher.exe" {
				continue // Intentionally ignored
			}

			if _, ok := dirs[p.Name]; !ok {
				t.Errorf("%s: upstream package %s missing in dirs", b, p.Name)
			}
		}
	}
}
