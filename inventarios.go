package main

import "fmt"

var (
	sliceProducto []producto
	mapProducto   = make(map[string]producto)
)

func inventario() {
	mapProducto["Trululu"] = producto{
		nombre:   "Trululu",
		precio:   2500,
		cantidad: 10,
	}

	mapProducto["Pirulito"] = producto{
		nombre:   "Pirulito",
		precio:   5000,
		cantidad: 100,
	}

	mapProducto["Festival"] = producto{
		nombre:   "Festival",
		precio:   7800,
		cantidad: 120,
	}

	mapProducto["Cofee"] = producto{
		nombre:   "Cofee",
		precio:   2400,
		cantidad: 80,
	}

	mapProducto["Trocipollo"] = producto{
		nombre:   "Trocipollo",
		precio:   8000,
		cantidad: 95,
	}

	sliceProducto = []producto{
		mapProducto["Trululu"],
		mapProducto["Pirulito"],
		mapProducto["Festival"],
		mapProducto["Cofee"],
		mapProducto["Trocipollo"],
	}

}

func main() {

	inventario()

	actividadDiaria()

}

type producto struct {
	nombre   string
	precio   int16
	cantidad int16
}

// Esta función esta muy bien
func actividadDiaria() {
	var nombreActividadRealizar string
	fmt.Println(" Ingresa la actividad que deseas realizar: Crear, Eliminar, Actualizar, Consultar ")
	fmt.Scanln(&nombreActividadRealizar)
	fmt.Println(" Ingresaste: ", nombreActividadRealizar)

	switch nombreActividadRealizar {
	case "Crear":
		productoNuevo := ingresoProductosNuevos()
		ingresoProductos(productoNuevo)
	case "Eliminar":
		eliminacionProducto()
	case "Actualizar":
		productoActualizado := ingresoProductoActualizar()
		actualizarProductos(productoActualizado)
	case "Consultar":
		consultaProductos()
	}
}

func ingresoProductosNuevos() producto {
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

	//repasa esto por fa la declaración de struct y como se usan
	var productoNuevo producto
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
func ingresoProductos(productoNuevo producto) {

	sliceProducto = append(sliceProducto, productoNuevo)
	mapProducto[productoNuevo.nombre] = productoNuevo
	fmt.Printf("Productos Slice: %v, Productos Map: %v", sliceProducto, mapProducto) //repetido
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
			delete(mapProducto, nombreProductoEliminar)
			fmt.Printf("Productos Slice: %v, Productos Map: %v", sliceProducto, mapProducto) //repetido

			return
		}

	}
}

// adicionarProductos esta función permite adicionar la cantidad de productos disponibles en el inventario
func ingresoProductoActualizar() producto {
	var nombreProductoActualizar string
	fmt.Println(" Ingresa el nombre del producto a Actualizar: ")
	fmt.Scanln(&nombreProductoActualizar)
	fmt.Println(" Ingresaste: ", nombreProductoActualizar)

	var cantidadProductoActualizar int16
	fmt.Println(" Ingresa la Cantidad de producto a Actualizar: ")
	fmt.Scanln(&cantidadProductoActualizar)
	fmt.Println(" La cantidad de Producto Actualizar es: ", cantidadProductoActualizar)

	var productoActualizado producto
	productoActualizado.nombre = nombreProductoActualizar
	productoActualizado.cantidad = cantidadProductoActualizar
	productoActualizado.precio = mapProducto[nombreProductoActualizar].precio

	return productoActualizado

}

// actualizarProductos esta función actualiza la cantidad de producto disponible
func actualizarProductos(productoActualizado producto) {

	for posicion, productoActualizar := range sliceProducto {
		if productoActualizar.nombre == productoActualizado.nombre {
			mapProducto[productoActualizado.nombre] = productoActualizado
			sliceProducto[posicion] = productoActualizado
			fmt.Printf("El producto: %s, Usted actualizó: %d", productoActualizado.nombre, productoActualizado.cantidad)
			fmt.Printf("Productos Slice: %v, Productos Map: %v", sliceProducto, mapProducto) //repetido
		}

	}

}

// consultaProductos esta funcion muestra los datos de los productos disponibles

func consultaProductos() {

	var nombreProductoConsultar string
	fmt.Println("Ingrese el nombre del producto a consultar:")
	fmt.Scanln(&nombreProductoConsultar)
	fmt.Println(" El producto a consultar es: ", nombreProductoConsultar)

	for _, productoActualizar := range sliceProducto {
		if productoActualizar.nombre == nombreProductoConsultar {
			fmt.Println(mapProducto[nombreProductoConsultar])
			return
		}

	}

	fmt.Println("Producto inválido")
}
