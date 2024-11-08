package rbxbin

import (
	"testing"

	cs "github.com/apprehensions/rbxweb/clientsettings"
)

func TestDirsCurrent(t *testing.T) {
	for _, b := range []cs.BinaryType{cs.WindowsPlayer, cs.WindowsStudio64} {
		d, err := GetDeployment(b, "")
		if err != nil {
			t.Fatalf("%s: %s", b, err)
		}

		ps, err := DefaultMirror.GetPackages(d)
		if err != nil {
			t.Fatalf("%s: pkgs: %s", b, err)
		}

		dirs, err := DefaultMirror.BinaryDirectories(d)
		if err != nil {
			t.Fatalf("%s: dirs: %s", b, err)
		}

		for zip := range dirs {
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
