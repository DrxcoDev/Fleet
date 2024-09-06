package main

// Dronedb - Free and open source software for aerial data storage
// Homepage: https://github.com/DroneDB/DroneDB

import (
	"fmt"
	
	"os/exec"
)

func installDronedb() {
	// Método 1: Descargar y extraer .tar.gz
	dronedb_tar_url := "https://github.com/DroneDB/DroneDB.git"
	dronedb_cmd_tar := exec.Command("curl", "-L", dronedb_tar_url, "-o", "package.tar.gz")
	err := dronedb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dronedb_zip_url := "https://github.com/DroneDB/DroneDB.git"
	dronedb_cmd_zip := exec.Command("curl", "-L", dronedb_zip_url, "-o", "package.zip")
	err = dronedb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dronedb_bin_url := "https://github.com/DroneDB/DroneDB.git"
	dronedb_cmd_bin := exec.Command("curl", "-L", dronedb_bin_url, "-o", "binary.bin")
	err = dronedb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dronedb_src_url := "https://github.com/DroneDB/DroneDB.git"
	dronedb_cmd_src := exec.Command("curl", "-L", dronedb_src_url, "-o", "source.tar.gz")
	err = dronedb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dronedb_cmd_direct := exec.Command("./binary")
	err = dronedb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gdal")
	exec.Command("latte", "install", "gdal").Run()
	fmt.Println("Instalando dependencia: libspatialite")
	exec.Command("latte", "install", "libspatialite").Run()
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
	fmt.Println("Instalando dependencia: pdal")
	exec.Command("latte", "install", "pdal").Run()
}