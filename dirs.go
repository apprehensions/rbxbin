package rbxbin

import (
	"github.com/apprehensions/rbxweb/clientsettings"
)

// PackageDirectories is a map of where Binary packages should go.
type PackageDirectories map[string]string

// BinaryDirectories retrieves the PackageDirectories for the given binary.
func BinaryDirectories(bt clientsettings.BinaryType) PackageDirectories {
	switch bt {
	case clientsettings.WindowsPlayer:
		return PlayerDirectories
	case clientsettings.WindowsStudio64:
		return StudioDirectories
	default:
		panic("unhandled binary type " + string(bt) + "for dirs")
	}
}

// WindowsPlayerDirectories is retrieved from [Bloxstrap].
//
// [Bloxstrap]: https://github.com/pizzaboxer/bloxstrap/blob/main/Bloxstrap/Bootstrapper.cs
var PlayerDirectories = PackageDirectories{
	"RobloxApp.zip": "",

	// Common
	"content-avatar.zip":            "content/avatar",
	"content-configs.zip":           "content/configs",
	"content-fonts.zip":             "content/fonts",
	"content-models.zip":            "content/models",
	"content-sky.zip":               "content/sky",
	"content-sounds.zip":            "content/sounds",
	"shaders.zip":                   "shaders/",
	"ssl.zip":                       "ssl/",
	"content-textures2.zip":         "content/textures",
	"content-textures3.zip":         "PlatformContent/pc/textures",
	"content-terrain.zip":           "PlatformContent/pc/terrain",
	"content-platform-fonts.zip":    "PlatformContent/pc/fonts",
	"extracontent-places.zip":       "ExtraContent/places",
	"extracontent-luapackages.zip":  "ExtraContent/LuaPackages",
	"extracontent-translations.zip": "ExtraContent/translations",
	"extracontent-models.zip":       "ExtraContent/models",
	"extracontent-textures.zip":     "ExtraContent/textures",
	"WebView2.zip":                  "",
	"WebView2RuntimeInstaller.zip":  "WebView2RuntimeInstaller",

	// RobloxPlayerLauncher.exe intentionally ignored
}

// StudioDirectories is retrieved from [Bloxstrap].
//
// [Bloxstrap]: https://github.com/pizzaboxer/bloxstrap/blob/main/Bloxstrap/Bootstrapper.cs
var StudioDirectories = PackageDirectories{
	"ApplicationConfig.zip":           "ApplicationConfig",
	"BuiltInPlugins.zip":              "BuiltInPlugins",
	"BuiltInStandalonePlugins.zip":    "BuiltInStandalonePlugins",
	"Plugins.zip":                     "Plugins",
	"Qml.zip":                         "Qml",
	"StudioFonts.zip":                 "StudioFonts",
	"RobloxStudio.zip":                "",
	"Libraries.zip":                   "",
	"LibrariesQt5.zip":                "",
	"content-qt_translations.zip":     "content/qt_translations",
	"content-studio_svg_textures.zip": "content/studio_svg_textures",
	"content-api-docs.zip":            "content/api_docs",
	"extracontent-scripts.zip":        "ExtraContent/scripts",
	"redist.zip":                      "",

	// Common
	"content-avatar.zip":            "content/avatar",
	"content-configs.zip":           "content/configs",
	"content-fonts.zip":             "content/fonts",
	"content-models.zip":            "content/models",
	"content-sky.zip":               "content/sky",
	"content-sounds.zip":            "content/sounds",
	"shaders.zip":                   "shaders/",
	"ssl.zip":                       "ssl/",
	"content-textures2.zip":         "content/textures",
	"content-textures3.zip":         "PlatformContent/pc/textures",
	"content-terrain.zip":           "PlatformContent/pc/terrain",
	"content-platform-fonts.zip":    "PlatformContent/pc/fonts",
	"extracontent-places.zip":       "ExtraContent/places",
	"extracontent-luapackages.zip":  "ExtraContent/LuaPackages",
	"extracontent-translations.zip": "ExtraContent/translations",
	"extracontent-models.zip":       "ExtraContent/models",
	"extracontent-textures.zip":     "ExtraContent/textures",
	"WebView2.zip":                  "",
	"WebView2RuntimeInstaller.zip":  "WebView2RuntimeInstaller",
}
