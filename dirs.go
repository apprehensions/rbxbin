package rbxbin

// This file is generated by dirgen. DO NOT EDIT.

import (
	"github.com/apprehensions/rbxweb/clientsettings"
)

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

var PlayerDirectories = PackageDirectories{
	// Ordered based on how they appear in JSON binary form
	"extracontent-places.zip":       "ExtraContent/places",
	"shaders.zip":                   "shaders",
	"content-avatar.zip":            "content/avatar",
	"content-textures3.zip":         "PlatformContent/pc/textures",
	"content-terrain.zip":           "PlatformContent/pc/terrain",
	"content-platform-fonts.zip":    "PlatformContent/pc/fonts",
	"extracontent-translations.zip": "ExtraContent/translations",
	"extracontent-textures.zip":     "ExtraContent/textures",
	"WebView2RuntimeInstaller.zip":  "WebView2RuntimeInstaller",
	"ssl.zip":                       "ssl",
	"RobloxApp.zip":                 ".",
	"content-configs.zip":           "content/configs",
	"content-sky.zip":               "content/sky",
	"content-sounds.zip":            "content/sounds",
	"content-models.zip":            "content/models",
	"extracontent-luapackages.zip":  "ExtraContent/LuaPackages",
	"redist.zip":                    ".",
	"content-fonts.zip":             "content/fonts",
	"content-textures2.zip":         "content/textures",
	"extracontent-models.zip":       "ExtraContent/models",
	"WebView2.zip":                  ".",
}

var StudioDirectories = PackageDirectories{
	// Ordered based on how they appear in JSON binary form
	"shaders.zip":                     "shaders",
	"ApplicationConfig.zip":           "ApplicationConfig",
	"content-sky.zip":                 "content/sky",
	"Qml.zip":                         "Qml",
	"content-platform-fonts.zip":      "PlatformContent/pc/fonts",
	"content-api-docs.zip":            "content/api_docs",
	"RibbonConfig.zip":                "RibbonConfig",
	"WebView2.zip":                    ".",
	"WebView2RuntimeInstaller.zip":    "WebView2RuntimeInstaller",
	"Libraries.zip":                   ".",
	"content-avatar.zip":              "content/avatar",
	"extracontent-scripts.zip":        "ExtraContent/scripts",
	"extracontent-luapackages.zip":    "ExtraContent/LuaPackages",
	"extracontent-translations.zip":   "ExtraContent/translations",
	"extracontent-models.zip":         "ExtraContent/models",
	"BuiltInStandalonePlugins.zip":    "BuiltInStandalonePlugins",
	"StudioFonts.zip":                 "StudioFonts",
	"content-fonts.zip":               "content/fonts",
	"content-studio_svg_textures.zip": "content/studio_svg_textures",
	"content-textures3.zip":           "PlatformContent/pc/textures",
	"BuiltInPlugins.zip":              "BuiltInPlugins",
	"ssl.zip":                         "ssl",
	"content-sounds.zip":              "content/sounds",
	"content-models.zip":              "content/models",
	"content-configs.zip":             "content/configs",
	"content-textures2.zip":           "content/textures",
	"LibrariesQt5.zip":                ".",
	"content-terrain.zip":             "PlatformContent/pc/terrain",
	"extracontent-textures.zip":       "ExtraContent/textures",
	"content-qt_translations.zip":     "content/qt_translations",
	"Plugins.zip":                     "Plugins",
	"redist.zip":                      ".",
	"RobloxStudio.zip":                ".",
}
