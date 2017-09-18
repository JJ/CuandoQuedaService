# CuandoQuedaService

API para asignaturas, usable desde bots o por su cuenta. Necesita
tener Go instalado.

## Instalar dependencias

	go get . 
	
O si se ha hecho algún cambio

	go get -u github.com/JJ/HitosIV
	
Para actualizar


## Ejecutar

Desde el mismo directorio,

    go run CuantoQuedaService.go
	
Y luego

    curl http://localhost:31415 
	
devuelve un JSON con todos los hitos.
