package main

import "fmt"

var (
	sliceProducto []productos
	mapProducto   map[string]productos
)

func main() {
	fmt.Print("Hola Mundo")
}

// producto esta estructura detalla cada producto
type productos struct {
	nombre   string
	precio   int16
	cantidad int16
}

// productosDisponibles esta funcion muestra los productos disponibles
func productosDisponibles(seleccionProducto string) {

	for _, producto := range sliceProducto {
		if producto.nombre == seleccionProducto {
			fmt.Println(producto.String())
		}

	}

	fmt.Println("Producto inválido:")
}

//ingresoProductos esta función agrega productos al inventario
var productoNuevo string
	fmt.Print(" Ingresa el producto nuevo: ")
	fmt.Scanln(&productoNuevo)
	fmt.Println(" Ingresaste: ", productoNuevo)

func ingresoProductos( productoNuevo []string) {
if productoNuevo != nil {
	sliceProducto = append(sliceProducto[:i], sliceProducto[i+1]...),
	fmt.Println()

}

}
