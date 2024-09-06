package main

// XcbProto - X.Org: XML-XCB protocol descriptions for libxcb code generation
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installXcbProto() {
	// Método 1: Descargar y extraer .tar.gz
	xcbproto_tar_url := "https://xorg.freedesktop.org/archive/individual/proto/xcb-proto-1.17.0.tar.xz"
	xcbproto_cmd_tar := exec.Command("curl", "-L", xcbproto_tar_url, "-o", "package.tar.gz")
	err := xcbproto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcbproto_zip_url := "https://xorg.freedesktop.org/archive/individual/proto/xcb-proto-1.17.0.tar.xz"
	xcbproto_cmd_zip := exec.Command("curl", "-L", xcbproto_zip_url, "-o", "package.zip")
	err = xcbproto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcbproto_bin_url := "https://xorg.freedesktop.org/archive/individual/proto/xcb-proto-1.17.0.tar.xz"
	xcbproto_cmd_bin := exec.Command("curl", "-L", xcbproto_bin_url, "-o", "binary.bin")
	err = xcbproto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcbproto_src_url := "https://xorg.freedesktop.org/archive/individual/proto/xcb-proto-1.17.0.tar.xz"
	xcbproto_cmd_src := exec.Command("curl", "-L", xcbproto_src_url, "-o", "source.tar.gz")
	err = xcbproto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcbproto_cmd_direct := exec.Command("./binary")
	err = xcbproto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}