package main

import "fmt"

var (
	sliceProducto []productos
	mapProducto   = make(map[string]productos)
)

func main() {

	mapProducto["Trululu"] = productos{
		nombre:   "Trululu",
		precio:   2500,
		cantidad: 10,
	}

	mapProducto["Pirulito"] = productos{
		nombre:   "Pirulito",
		precio:   5000,
		cantidad: 100,
	}

	mapProducto["Festival"] = productos{
		nombre:   "Festival",
		precio:   7800,
		cantidad: 120,
	}

	mapProducto["Cofee"] = productos{
		nombre:   "Cofee",
		precio:   2400,
		cantidad: 80,
	}

	mapProducto["Trocipollo"] = productos{
		nombre:   "Trocipollo",
		precio:   8000,
		cantidad: 95,
	}

	sliceProducto = []productos{
		mapProducto["Trululu"],
		mapProducto["Pirulito"],
		mapProducto["Festival"],
		mapProducto["Cofee"],
		mapProducto["Trocipollo"],
	}

	actividadDiaria()

}

// producto esta estructura detalla cada producto
type productos struct {
	nombre   string
	precio   int16
	cantidad int16
}

func actividadDiaria() {
	var nombreActividadRealizar string
	fmt.Println(" Ingresa la actividad que deseas realizar: Crear, Eliminar, Adicionar, Consultar ")
	fmt.Scanln(&nombreActividadRealizar)
	fmt.Println(" Ingresaste: ", nombreActividadRealizar)

	switch nombreActividadRealizar {
	case "Crear":
		productoNuevo := ingresoProductosNuevos()
		ingresoProductos(productoNuevo)
	case "Eliminar":
		eliminacionProducto()
	case "Adicionar":
		productoAdicionado := ingresoActualizacionCantidad()
		adicionProductos(productoAdicionado)
	case "Consultar":
		var seleccionProducto string
		fmt.Println("Ingrese el producto a consultar:")
		fmt.Scanln(&seleccionProducto)
		consultaProductos(seleccionProducto)
	default:
		fmt.Println("El valor ingresado no es válido:")
	}
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
	mapProducto[productoNuevo.nombre] = productoNuevo
	fmt.Printf("Productos Slice: %v, Productos Map: %v", sliceProducto, mapProducto)
}

// eliminacionProducto esta función elimina producto del inventario
func eliminacionProducto() {
	nombreProductoEliminar := ""
	fmt.Println(" Ingresa el producto a eliminar: ")
	fmt.Scanln(&nombreProductoEliminar)
	fmt.Println(" Ingresaste: ", nombreProductoEliminar)

	for posicion, productoEliminar := range sliceProducto {
		if nombreProductoEliminar == productoEliminar.nombre {
			sliceProducto = append(sliceProducto[0:posicion], sliceProducto[posicion+1:]...)
			fmt.Printf("Productos Slice: %v, Productos Map: %v", sliceProducto, mapProducto)

			return
		}

	}
}

// adicionarProductos esta función permite adicionar la cantidad de productos disponibles en el inventario
func ingresoActualizacionCantidad() productos {
	var nombreProductoAdicionar string
	fmt.Println(" Ingresa el nombre del producto a Adicionar: ")
	fmt.Scanln(&nombreProductoAdicionar)
	fmt.Println(" Ingresaste: ", nombreProductoAdicionar)

	var cantidadProductoAdicionar int16
	fmt.Println(" Ingresa la Cantidad de producto a Adicionar: ")
	fmt.Scanln(&cantidadProductoAdicionar)
	fmt.Println(" La cantidad de Producto Adicionado es: ", cantidadProductoAdicionar)

	var productoAdicionado productos
	productoAdicionado.nombre = nombreProductoAdicionar
	productoAdicionado.cantidad = cantidadProductoAdicionar

	return productoAdicionado

}

// adicionProductos esta función agrega productos al inventario
func adicionProductos(productoAdicionado productos) {

	sliceProducto = append(sliceProducto, productoAdicionado)
	mapProducto[productoAdicionado.nombre] = productoAdicionado
	fmt.Printf("Al producto: %s, Usted adicionó: %d", productoAdicionado.nombre, productoAdicionado.cantidad)
	//fmt.Println(sliceProducto, mapProducto)
}

// consultaProductos esta funcion muestra los productos disponibles
func consultaProductos(seleccionProducto string) {

	for _, producto := range sliceProducto {
		if producto.nombre == seleccionProducto {
			fmt.Println(producto.nombre, producto.cantidad, producto.precio)
			return
		}

	}

	fmt.Println("Producto inválido")
}
