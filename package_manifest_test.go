package rbxbin

import (
	"errors"
	"strings"
	"testing"
)

func TestParsePackages(t *testing.T) {
	manifest := []byte(strings.Join([]string{
		"v0",
		"foo.zip",
		"026b271a21b03f2e564c036525356db5",
		"71367142",
		"109436874",
		"bar.zip",
		"4d9ec7b52a29c80f3ce1f6a65b14b563",
		"408629",
		"1191394",
		"", // Roblox behavior
	}, "\r\n"))

	pkgs, err := ParsePackages(manifest)
	if err != nil {
		t.Fatal(err)
	}

	pkgFooWant := Package{
		Name:     "foo.zip",
		Checksum: "026b271a21b03f2e564c036525356db5",
		Size:     109436874,
		ZipSize:  71367142,
	}

	pkgBarWant := Package{
		Name:     "bar.zip",
		Checksum: "4d9ec7b52a29c80f3ce1f6a65b14b563",
		Size:     1191394,
		ZipSize:  408629,
	}

	if pkgs[0] != pkgFooWant {
		t.Fatalf("package %v, want package match for %v", pkgs[0], pkgFooWant)
	}

	if pkgs[1] != pkgBarWant {
		t.Fatalf("package %v, want package match for %v", pkgs[0], pkgBarWant)
	}
}

func TestInvalidPackagePackageManifest(t *testing.T) {
	manifest := []byte(strings.Join([]string{
		"v0",
		"foo.zip",
		"026b271a21b03f2e564c036525356db5",
		"71367142",
		"", // Roblox behavior
	}, "\r\n"))

	_, err := ParsePackages(manifest)
	if !errors.Is(err, ErrInvalidPkgManifest) {
		t.Fail()
	}
}

func TestUnhandledPackagePackageManifest(t *testing.T) {
	manifest := []byte(strings.Join([]string{
		"v1",
		"foo.zip",
		"026b271a21b03f2e564c036525356db5",
		"71367142",
		"109436874",
		"", // Roblox behavior
	}, "\r\n"))

	_, err := ParsePackages(manifest)
	if !errors.Is(err, ErrUnhandledPkgManifestVer) {
		t.Fail()
	}
}
