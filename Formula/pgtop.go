package main

// PgTop - Monitor PostgreSQL processes
// Homepage: https://pg_top.gitlab.io

import (
	"fmt"
	
	"os/exec"
)

func installPgTop() {
	// Método 1: Descargar y extraer .tar.gz
	pgtop_tar_url := "https://ftp.postgresql.org/pub/projects/pgFoundry/ptop/pg_top/3.7.0/pg_top-3.7.0.tar.bz2"
	pgtop_cmd_tar := exec.Command("curl", "-L", pgtop_tar_url, "-o", "package.tar.gz")
	err := pgtop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgtop_zip_url := "https://ftp.postgresql.org/pub/projects/pgFoundry/ptop/pg_top/3.7.0/pg_top-3.7.0.tar.bz2"
	pgtop_cmd_zip := exec.Command("curl", "-L", pgtop_zip_url, "-o", "package.zip")
	err = pgtop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgtop_bin_url := "https://ftp.postgresql.org/pub/projects/pgFoundry/ptop/pg_top/3.7.0/pg_top-3.7.0.tar.bz2"
	pgtop_cmd_bin := exec.Command("curl", "-L", pgtop_bin_url, "-o", "binary.bin")
	err = pgtop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgtop_src_url := "https://ftp.postgresql.org/pub/projects/pgFoundry/ptop/pg_top/3.7.0/pg_top-3.7.0.tar.bz2"
	pgtop_cmd_src := exec.Command("curl", "-L", pgtop_src_url, "-o", "source.tar.gz")
	err = pgtop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgtop_cmd_direct := exec.Command("./binary")
	err = pgtop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
}