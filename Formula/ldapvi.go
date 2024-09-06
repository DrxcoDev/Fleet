package main

// Ldapvi - Update LDAP entries with a text editor
// Homepage: http://www.lichteblau.com/ldapvi/

import (
	"fmt"
	
	"os/exec"
)

func installLdapvi() {
	// Método 1: Descargar y extraer .tar.gz
	ldapvi_tar_url := "http://www.lichteblau.com/download/ldapvi-1.7.tar.gz"
	ldapvi_cmd_tar := exec.Command("curl", "-L", ldapvi_tar_url, "-o", "package.tar.gz")
	err := ldapvi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ldapvi_zip_url := "http://www.lichteblau.com/download/ldapvi-1.7.zip"
	ldapvi_cmd_zip := exec.Command("curl", "-L", ldapvi_zip_url, "-o", "package.zip")
	err = ldapvi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ldapvi_bin_url := "http://www.lichteblau.com/download/ldapvi-1.7.bin"
	ldapvi_cmd_bin := exec.Command("curl", "-L", ldapvi_bin_url, "-o", "binary.bin")
	err = ldapvi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ldapvi_src_url := "http://www.lichteblau.com/download/ldapvi-1.7.src.tar.gz"
	ldapvi_cmd_src := exec.Command("curl", "-L", ldapvi_src_url, "-o", "source.tar.gz")
	err = ldapvi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ldapvi_cmd_direct := exec.Command("./binary")
	err = ldapvi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: popt")
	exec.Command("latte", "install", "popt").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}