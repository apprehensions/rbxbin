package rbxbin

import (
	"os"
	"path/filepath"
)

// Contains necessary application settings required for Roblox to run.
const AppSettings = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\r\n" +
	"<Settings>\r\n" +
	"        <ContentFolder>content</ContentFolder>\r\n" +
	"        <BaseUrl>http://www.roblox.com</BaseUrl>\r\n" +
	"</Settings>\r\n"

// WriteAppSettings writes [AppSettings] to a AppSettings.xml file
// in a binary's deployment directory.
func WriteAppSettings(dir string) error {
	as := filepath.Join(dir, "AppSettings.xml")

	f, err := os.Create(as)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(AppSettings); err != nil {
		return err
	}

	return nil
}
