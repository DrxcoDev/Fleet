package main

// SteamguardCli - CLI for steamguard
// Homepage: https://github.com/dyc3/steamguard-cli

import (
	"fmt"
	
	"os/exec"
)

func installSteamguardCli() {
	// Método 1: Descargar y extraer .tar.gz
	steamguardcli_tar_url := "https://github.com/dyc3/steamguard-cli/archive/refs/tags/v0.14.2.tar.gz"
	steamguardcli_cmd_tar := exec.Command("curl", "-L", steamguardcli_tar_url, "-o", "package.tar.gz")
	err := steamguardcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	steamguardcli_zip_url := "https://github.com/dyc3/steamguard-cli/archive/refs/tags/v0.14.2.zip"
	steamguardcli_cmd_zip := exec.Command("curl", "-L", steamguardcli_zip_url, "-o", "package.zip")
	err = steamguardcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	steamguardcli_bin_url := "https://github.com/dyc3/steamguard-cli/archive/refs/tags/v0.14.2.bin"
	steamguardcli_cmd_bin := exec.Command("curl", "-L", steamguardcli_bin_url, "-o", "binary.bin")
	err = steamguardcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	steamguardcli_src_url := "https://github.com/dyc3/steamguard-cli/archive/refs/tags/v0.14.2.src.tar.gz"
	steamguardcli_cmd_src := exec.Command("curl", "-L", steamguardcli_src_url, "-o", "source.tar.gz")
	err = steamguardcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	steamguardcli_cmd_direct := exec.Command("./binary")
	err = steamguardcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}