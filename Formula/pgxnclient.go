package main

// Pgxnclient - Command-line client for the PostgreSQL Extension Network
// Homepage: https://pgxn.github.io/pgxnclient/

import (
	"fmt"
	
	"os/exec"
)

func installPgxnclient() {
	// Método 1: Descargar y extraer .tar.gz
	pgxnclient_tar_url := "https://files.pythonhosted.org/packages/54/3d/5eae61996702ce218548a98f6ccc930a80b1e4b09b7a8384b1a95129a9c2/pgxnclient-1.3.2.tar.gz"
	pgxnclient_cmd_tar := exec.Command("curl", "-L", pgxnclient_tar_url, "-o", "package.tar.gz")
	err := pgxnclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgxnclient_zip_url := "https://files.pythonhosted.org/packages/54/3d/5eae61996702ce218548a98f6ccc930a80b1e4b09b7a8384b1a95129a9c2/pgxnclient-1.3.2.zip"
	pgxnclient_cmd_zip := exec.Command("curl", "-L", pgxnclient_zip_url, "-o", "package.zip")
	err = pgxnclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgxnclient_bin_url := "https://files.pythonhosted.org/packages/54/3d/5eae61996702ce218548a98f6ccc930a80b1e4b09b7a8384b1a95129a9c2/pgxnclient-1.3.2.bin"
	pgxnclient_cmd_bin := exec.Command("curl", "-L", pgxnclient_bin_url, "-o", "binary.bin")
	err = pgxnclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgxnclient_src_url := "https://files.pythonhosted.org/packages/54/3d/5eae61996702ce218548a98f6ccc930a80b1e4b09b7a8384b1a95129a9c2/pgxnclient-1.3.2.src.tar.gz"
	pgxnclient_cmd_src := exec.Command("curl", "-L", pgxnclient_src_url, "-o", "source.tar.gz")
	err = pgxnclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgxnclient_cmd_direct := exec.Command("./binary")
	err = pgxnclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}