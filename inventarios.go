package main

import "fmt"

var (
	sliceProducto []productos
	mapProducto   map[string]productos
)

func main() {
	productoNuevo := ingresoProductosNuevos()
	ingresoProductos(productoNuevo)

}

// producto esta estructura detalla cada producto
type productos struct {
	nombre   string
	precio   int16
	cantidad int16
}

func (productos productos) String() string {
	return fmt.Sprintf("Nombre: %s, Precio: %d, Cantidad: %d", productos.nombre, productos.precio, productos.cantidad)

}

// productosDisponibles esta funcion muestra los productos disponibles
func productosDisponibles(seleccionProducto string) {

	for _, producto := range sliceProducto {
		if producto.nombre == seleccionProducto {
			fmt.Println(producto.String())
			return
		}

	}

	fmt.Println("Producto inválido")
}

func ingresoProductosNuevos() productos {
	var nombreProductoNuevo string
	fmt.Println(" Ingresa el producto nuevo: ")
	fmt.Scanln(&nombreProductoNuevo)
	fmt.Println(" Ingresaste: ", nombreProductoNuevo)

	var precioProductoNuevo int16
	fmt.Println(" Ingresa el precio del producto nuevo: ")
	fmt.Scanln(&precioProductoNuevo)
	fmt.Println(" El precio del Producto Nuevo es: ", precioProductoNuevo)

	var cantidadProductoNuevo int16
	fmt.Println(" Ingresa la Cantidad de producto nuevo: ")
	fmt.Scanln(&cantidadProductoNuevo)
	fmt.Println(" La cantidad de Producto Nuevo es: ", cantidadProductoNuevo)

	var productoNuevo productos
	productoNuevo.nombre = nombreProductoNuevo
	productoNuevo.precio = precioProductoNuevo
	productoNuevo.cantidad = cantidadProductoNuevo

	// productoNuevo := productos{
	// 	nombre:   nombreProductoNuevo,
	// 	precio:   precioProductoNuevo,
	// 	cantidad: cantidadProductoNuevo,
	// }
	return productoNuevo
}

// ingresoProductos esta función agrega productos al inventario
func ingresoProductos(productoNuevo productos) {

	sliceProducto = append(sliceProducto, productoNuevo)
	fmt.Println(sliceProducto)

}
