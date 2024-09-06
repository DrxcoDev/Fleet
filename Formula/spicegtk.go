package main

// SpiceGtk - GTK client/libraries for SPICE
// Homepage: https://www.spice-space.org

import (
	"fmt"
	
	"os/exec"
)

func installSpiceGtk() {
	// Método 1: Descargar y extraer .tar.gz
	spicegtk_tar_url := "https://www.spice-space.org/download/gtk/spice-gtk-0.42.tar.xz"
	spicegtk_cmd_tar := exec.Command("curl", "-L", spicegtk_tar_url, "-o", "package.tar.gz")
	err := spicegtk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spicegtk_zip_url := "https://www.spice-space.org/download/gtk/spice-gtk-0.42.tar.xz"
	spicegtk_cmd_zip := exec.Command("curl", "-L", spicegtk_zip_url, "-o", "package.zip")
	err = spicegtk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spicegtk_bin_url := "https://www.spice-space.org/download/gtk/spice-gtk-0.42.tar.xz"
	spicegtk_cmd_bin := exec.Command("curl", "-L", spicegtk_bin_url, "-o", "binary.bin")
	err = spicegtk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spicegtk_src_url := "https://www.spice-space.org/download/gtk/spice-gtk-0.42.tar.xz"
	spicegtk_cmd_src := exec.Command("curl", "-L", spicegtk_src_url, "-o", "source.tar.gz")
	err = spicegtk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spicegtk_cmd_direct := exec.Command("./binary")
	err = spicegtk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gstreamer")
	exec.Command("latte", "install", "gstreamer").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: libepoxy")
	exec.Command("latte", "install", "libepoxy").Run()
	fmt.Println("Instalando dependencia: libsoup")
	exec.Command("latte", "install", "libsoup").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: phodav")
	exec.Command("latte", "install", "phodav").Run()
	fmt.Println("Instalando dependencia: pixman")
	exec.Command("latte", "install", "pixman").Run()
	fmt.Println("Instalando dependencia: spice-protocol")
	exec.Command("latte", "install", "spice-protocol").Run()
	fmt.Println("Instalando dependencia: usbredir")
	exec.Command("latte", "install", "usbredir").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: libva")
	exec.Command("latte", "install", "libva").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
}