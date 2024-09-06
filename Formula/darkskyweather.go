package main

// DarkskyWeather - Command-line weather from the darksky.net API
// Homepage: https://github.com/genuinetools/weather

import (
	"fmt"
	
	"os/exec"
)

func installDarkskyWeather() {
	// Método 1: Descargar y extraer .tar.gz
	darkskyweather_tar_url := "https://github.com/genuinetools/weather/archive/refs/tags/v0.15.7.tar.gz"
	darkskyweather_cmd_tar := exec.Command("curl", "-L", darkskyweather_tar_url, "-o", "package.tar.gz")
	err := darkskyweather_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	darkskyweather_zip_url := "https://github.com/genuinetools/weather/archive/refs/tags/v0.15.7.zip"
	darkskyweather_cmd_zip := exec.Command("curl", "-L", darkskyweather_zip_url, "-o", "package.zip")
	err = darkskyweather_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	darkskyweather_bin_url := "https://github.com/genuinetools/weather/archive/refs/tags/v0.15.7.bin"
	darkskyweather_cmd_bin := exec.Command("curl", "-L", darkskyweather_bin_url, "-o", "binary.bin")
	err = darkskyweather_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	darkskyweather_src_url := "https://github.com/genuinetools/weather/archive/refs/tags/v0.15.7.src.tar.gz"
	darkskyweather_cmd_src := exec.Command("curl", "-L", darkskyweather_src_url, "-o", "source.tar.gz")
	err = darkskyweather_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	darkskyweather_cmd_direct := exec.Command("./binary")
	err = darkskyweather_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}