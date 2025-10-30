package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1: DEFINICIÓN DE ESTRUCTURAS (MODELOS)

type Cliente struct {
	ID       int
	Nombre   string
	Telefono string
	Email    string
}

type Vehiculo struct {
	Matricula    string
	Marca        string
	Modelo       string
	FechaEntrada string
	FechaSalida  string
	IDCliente    int
	IDIncidencia int
	EnTaller     bool
}

type Incidencia struct {
	ID          int
	MecanicosID []int
	Tipo        string
	Prioridad   string
	Descripcion string
	Estado      string
}

type Mecanico struct {
	ID           int
	Nombre       string
	Especialidad string
	Experiencia  int
	Activo       bool
}

// 2: VARIABLES GLOBALES Y ALMACENAMIENTO

var (
	clientes       []Cliente
	vehiculos      []Vehiculo
	incidencias    []Incidencia
	mecanicos      []Mecanico
	plazasOcupadas int
	totalPlazas    int
	scanner        = bufio.NewScanner(os.Stdin)
)

// 3: FUNCIONES AUXILIARES

func inicializarDatos() {
	totalPlazas = 0
	plazasOcupadas = 0
}

func limpiarPantalla() {
	fmt.Print("\033[H\033[2J")
}

func leerLinea() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// 4: GESTIÓN DE CLIENTES

func menuClientes() {
	for {
		limpiarPantalla()
		fmt.Println("\n--- GESTIÓN DE CLIENTES ---")
		fmt.Println("1. Crear cliente")
		fmt.Println("2. Visualizar clientes")
		fmt.Println("3. Modificar cliente")
		fmt.Println("4. Eliminar cliente")
		fmt.Println("5. Volver")
		fmt.Print("Opción: ")

		opcion := leerLinea()

		switch opcion {
		case "1":
			crearCliente()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "2":
			visualizarClientes()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "3":
			modificarCliente()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "4":
			eliminarCliente()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "5":
			return
		default:
			fmt.Println("Opción no válida")
			fmt.Print("Presione Enter para continuar...")
			leerLinea()
		}
	}
}

func crearCliente() {
	fmt.Println("\n--- CREAR CLIENTE ---")
	fmt.Print("ID: ")
	id, _ := strconv.Atoi(leerLinea())

	// Validación: verificar que el ID no exista
	for _, c := range clientes {
		if c.ID == id {
			fmt.Println("Error: Ya existe un cliente con ese ID")
			return
		}
	}

	fmt.Print("Nombre: ")
	nombre := leerLinea()
	fmt.Print("Teléfono: ")
	telefono := leerLinea()
	fmt.Print("Email: ")
	email := leerLinea()

	// Crear y agregar el nuevo cliente
	cliente := Cliente{id, nombre, telefono, email}
	clientes = append(clientes, cliente)
	fmt.Println("Cliente creado correctamente")
}

func visualizarClientes() {
	fmt.Println("\n--- LISTA DE CLIENTES ---")
	if len(clientes) == 0 {
		fmt.Println("No hay clientes registrados")
		return
	}
	for _, c := range clientes {
		fmt.Printf("ID: %d | Nombre: %s | Tel: %s | Email: %s\n",
			c.ID, c.Nombre, c.Telefono, c.Email)
	}
}

func modificarCliente() {
	fmt.Print("ID del cliente a modificar: ")
	id, _ := strconv.Atoi(leerLinea())

	for i := range clientes {
		if clientes[i].ID == id {
			fmt.Print("Nuevo nombre (enter para mantener): ")
			nombre := leerLinea()
			if nombre != "" {
				clientes[i].Nombre = nombre
			}
			fmt.Print("Nuevo teléfono (enter para mantener): ")
			telefono := leerLinea()
			if telefono != "" {
				clientes[i].Telefono = telefono
			}
			fmt.Print("Nuevo email (enter para mantener): ")
			email := leerLinea()
			if email != "" {
				clientes[i].Email = email
			}
			fmt.Println("Cliente modificado correctamente")
			return
		}
	}
	fmt.Println("Cliente no encontrado")
}

func eliminarCliente() {
	fmt.Print("ID del cliente a eliminar: ")
	id, _ := strconv.Atoi(leerLinea())

	for i := range clientes {
		if clientes[i].ID == id {
			// Eliminar usando slicing
			clientes = append(clientes[:i], clientes[i+1:]...)
			fmt.Println("Cliente eliminado correctamente")
			return
		}
	}
	fmt.Println("Cliente no encontrado")
}

// 5: GESTIÓN DE VEHÍCULOS

func menuVehiculos() {
	for {
		limpiarPantalla()
		fmt.Println("\n--- GESTIÓN DE VEHÍCULOS ---")
		fmt.Println("1. Crear vehículo")
		fmt.Println("2. Visualizar vehículos")
		fmt.Println("3. Modificar vehículo")
		fmt.Println("4. Eliminar vehículo")
		fmt.Println("5. Volver")
		fmt.Print("Opción: ")

		opcion := leerLinea()

		switch opcion {
		case "1":
			crearVehiculo()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "2":
			visualizarVehiculos()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "3":
			modificarVehiculo()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "4":
			eliminarVehiculo()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "5":
			return
		default:
			fmt.Println("Opción no válida")
			fmt.Print("Presione Enter para continuar...")
			leerLinea()
		}
	}
}

func crearVehiculo() {
	fmt.Println("\n--- CREAR VEHÍCULO ---")
	fmt.Print("Matrícula: ")
	matricula := leerLinea()

	// Validación: verificar que la matrícula no exista
	for _, v := range vehiculos {
		if v.Matricula == matricula {
			fmt.Println("Error: Ya existe un vehículo con esa matrícula")
			return
		}
	}

	fmt.Print("Marca: ")
	marca := leerLinea()
	fmt.Print("Modelo: ")
	modelo := leerLinea()
	fmt.Print("Fecha de entrada (DD/MM/AAAA): ")
	fechaEntrada := leerLinea()
	fmt.Print("Fecha estimada de salida (DD/MM/AAAA): ")
	fechaSalida := leerLinea()
	fmt.Print("ID del cliente: ")
	idCliente, _ := strconv.Atoi(leerLinea())

	// Crear el nuevo vehículo
	vehiculo := Vehiculo{
		Matricula:    matricula,
		Marca:        marca,
		Modelo:       modelo,
		FechaEntrada: fechaEntrada,
		FechaSalida:  fechaSalida,
		IDCliente:    idCliente,
		EnTaller:     false,
	}
	vehiculos = append(vehiculos, vehiculo)
	fmt.Println("Vehículo creado correctamente")
}

func visualizarVehiculos() {
	fmt.Println("\n--- LISTA DE VEHÍCULOS ---")
	if len(vehiculos) == 0 {
		fmt.Println("No hay vehículos registrados")
		return
	}
	for _, v := range vehiculos {
		fmt.Printf("Matrícula: %s | Marca: %s | Modelo: %s | Cliente: %d | En taller: %v\n",
			v.Matricula, v.Marca, v.Modelo, v.IDCliente, v.EnTaller)
	}
}

func modificarVehiculo() {
	fmt.Print("Matrícula del vehículo a modificar: ")
	matricula := leerLinea()

	for i := range vehiculos {
		if vehiculos[i].Matricula == matricula {
			fmt.Print("Nueva marca (enter para mantener): ")
			marca := leerLinea()
			if marca != "" {
				vehiculos[i].Marca = marca
			}
			fmt.Print("Nuevo modelo (enter para mantener): ")
			modelo := leerLinea()
			if modelo != "" {
				vehiculos[i].Modelo = modelo
			}
			fmt.Println("Vehículo modificado correctamente")
			return
		}
	}
	fmt.Println("Vehículo no encontrado")
}

func eliminarVehiculo() {
	fmt.Print("Matrícula del vehículo a eliminar: ")
	matricula := leerLinea()

	for i := range vehiculos {
		if vehiculos[i].Matricula == matricula {
			// Si está en el taller, liberar la plaza
			if vehiculos[i].EnTaller {
				plazasOcupadas--
			}
			vehiculos = append(vehiculos[:i], vehiculos[i+1:]...)
			fmt.Println("Vehículo eliminado correctamente")
			return
		}
	}
	fmt.Println("Vehículo no encontrado")
}

// 6: GESTIÓN DE INCIDENCIAS

func menuIncidencias() {
	for {
		limpiarPantalla()
		fmt.Println("\n--- GESTIÓN DE INCIDENCIAS ---")
		fmt.Println("1. Crear incidencia")
		fmt.Println("2. Visualizar incidencias")
		fmt.Println("3. Modificar incidencia")
		fmt.Println("4. Eliminar incidencia")
		fmt.Println("5. Cambiar estado de incidencia")
		fmt.Println("6. Volver")
		fmt.Print("Opción: ")

		opcion := leerLinea()

		switch opcion {
		case "1":
			crearIncidencia()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "2":
			visualizarIncidencias()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "3":
			modificarIncidencia()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "4":
			eliminarIncidencia()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "5":
			cambiarEstadoIncidencia()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "6":
			return
		default:
			fmt.Println("Opción no válida")
			fmt.Print("Presione Enter para continuar...")
			leerLinea()
		}
	}
}

func crearIncidencia() {
	fmt.Println("\n--- CREAR INCIDENCIA ---")
	fmt.Print("ID de incidencia: ")
	id, _ := strconv.Atoi(leerLinea())

	// Validación: verificar que el ID no exista
	for _, inc := range incidencias {
		if inc.ID == id {
			fmt.Println("Error: Ya existe una incidencia con ese ID")
			return
		}
	}

	fmt.Print("Tipo (mecánica/eléctrica/carrocería): ")
	tipo := leerLinea()
	fmt.Print("Prioridad (baja/media/alta): ")
	prioridad := leerLinea()
	fmt.Print("Descripción: ")
	descripcion := leerLinea()
	fmt.Print("Matrícula del vehículo: ")
	matricula := leerLinea()

	// Relacionar la incidencia con el vehículo
	for i := range vehiculos {
		if vehiculos[i].Matricula == matricula {
			vehiculos[i].IDIncidencia = id
			break
		}
	}

	// Asignar mecánicos a la incidencia
	fmt.Print("IDs de mecánicos asignados (separados por comas): ")
	mecanicosStr := leerLinea()
	var mecanicosID []int
	if mecanicosStr != "" {
		ids := strings.Split(mecanicosStr, ",")
		for _, idStr := range ids {
			idMec, _ := strconv.Atoi(strings.TrimSpace(idStr))
			mecanicosID = append(mecanicosID, idMec)
		}
	}

	// Crear la nueva incidencia
	incidencia := Incidencia{
		ID:          id,
		MecanicosID: mecanicosID,
		Tipo:        tipo,
		Prioridad:   prioridad,
		Descripcion: descripcion,
		Estado:      "abierta",
	}
	incidencias = append(incidencias, incidencia)
	fmt.Println("Incidencia creada correctamente")
}

func visualizarIncidencias() {
	fmt.Println("\n--- LISTA DE INCIDENCIAS ---")
	if len(incidencias) == 0 {
		fmt.Println("No hay incidencias registradas")
		return
	}
	for _, inc := range incidencias {
		fmt.Printf("ID: %d | Tipo: %s | Prioridad: %s | Estado: %s\n",
			inc.ID, inc.Tipo, inc.Prioridad, inc.Estado)
		fmt.Printf("  Descripción: %s\n", inc.Descripcion)
		fmt.Printf("  Mecánicos: %v\n", inc.MecanicosID)
	}
}

func modificarIncidencia() {
	fmt.Print("ID de la incidencia a modificar: ")
	id, _ := strconv.Atoi(leerLinea())

	for i := range incidencias {
		if incidencias[i].ID == id {
			fmt.Print("Nueva descripción (enter para mantener): ")
			desc := leerLinea()
			if desc != "" {
				incidencias[i].Descripcion = desc
			}
			fmt.Print("Nueva prioridad (enter para mantener): ")
			prioridad := leerLinea()
			if prioridad != "" {
				incidencias[i].Prioridad = prioridad
			}
			fmt.Println("Incidencia modificada correctamente")
			return
		}
	}
	fmt.Println("Incidencia no encontrada")
}

func eliminarIncidencia() {
	fmt.Print("ID de la incidencia a eliminar: ")
	id, _ := strconv.Atoi(leerLinea())

	for i := range incidencias {
		if incidencias[i].ID == id {
			incidencias = append(incidencias[:i], incidencias[i+1:]...)
			fmt.Println("Incidencia eliminada correctamente")
			return
		}
	}
	fmt.Println("Incidencia no encontrada")
}

func cambiarEstadoIncidencia() {
	fmt.Print("ID de la incidencia: ")
	id, _ := strconv.Atoi(leerLinea())

	for i := range incidencias {
		if incidencias[i].ID == id {
			fmt.Println("Estado actual:", incidencias[i].Estado)
			fmt.Print("Nuevo estado (abierta/en proceso/cerrada): ")
			estado := leerLinea()
			incidencias[i].Estado = estado
			fmt.Println("Estado actualizado correctamente")
			return
		}
	}
	fmt.Println("Incidencia no encontrada")
}

// 7: GESTIÓN DE MECÁNICOS

func menuMecanicos() {
	for {
		limpiarPantalla()
		fmt.Println("\n--- GESTIÓN DE MECÁNICOS ---")
		fmt.Println("1. Crear mecánico")
		fmt.Println("2. Visualizar mecánicos")
		fmt.Println("3. Modificar mecánico")
		fmt.Println("4. Eliminar mecánico")
		fmt.Println("5. Dar de alta/baja mecánico")
		fmt.Println("6. Volver")
		fmt.Print("Opción: ")

		opcion := leerLinea()

		switch opcion {
		case "1":
			crearMecanico()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "2":
			visualizarMecanicos()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "3":
			modificarMecanico()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "4":
			eliminarMecanico()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "5":
			cambiarEstadoMecanico()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "6":
			return
		default:
			fmt.Println("Opción no válida")
			fmt.Print("Presione Enter para continuar...")
			leerLinea()
		}
	}
}

func crearMecanico() {
	fmt.Println("\n--- CREAR MECÁNICO ---")
	fmt.Print("ID: ")
	id, _ := strconv.Atoi(leerLinea())

	// Validación: verificar que el ID no exista
	for _, m := range mecanicos {
		if m.ID == id {
			fmt.Println("Error: Ya existe un mecánico con ese ID")
			return
		}
	}

	fmt.Print("Nombre: ")
	nombre := leerLinea()
	fmt.Print("Especialidad (mecanica/electrica/carroceria): ")
	especialidad := leerLinea()
	fmt.Print("Años de experiencia: ")
	experiencia, _ := strconv.Atoi(leerLinea())

	// Crear mecánico y actualizar plazas del taller
	mecanico := Mecanico{id, nombre, especialidad, experiencia, true}
	mecanicos = append(mecanicos, mecanico)
	totalPlazas = len(mecanicos) * 2 // Cada mecánico gestiona 2 plazas
	fmt.Println("Mecánico creado correctamente")
}

func visualizarMecanicos() {
	fmt.Println("\n--- LISTA DE MECÁNICOS ---")
	if len(mecanicos) == 0 {
		fmt.Println("No hay mecánicos registrados")
		return
	}
	for _, m := range mecanicos {
		estado := "Activo"
		if !m.Activo {
			estado = "Baja"
		}
		fmt.Printf("ID: %d | Nombre: %s | Especialidad: %s | Exp: %d años | Estado: %s\n",
			m.ID, m.Nombre, m.Especialidad, m.Experiencia, estado)
	}
}

func modificarMecanico() {
	fmt.Print("ID del mecánico a modificar: ")
	id, _ := strconv.Atoi(leerLinea())

	for i := range mecanicos {
		if mecanicos[i].ID == id {
			fmt.Print("Nuevo nombre (enter para mantener): ")
			nombre := leerLinea()
			if nombre != "" {
				mecanicos[i].Nombre = nombre
			}
			fmt.Print("Nueva especialidad (enter para mantener): ")
			especialidad := leerLinea()
			if especialidad != "" {
				mecanicos[i].Especialidad = especialidad
			}
			fmt.Println("Mecánico modificado correctamente")
			return
		}
	}
	fmt.Println("Mecánico no encontrado")
}

func eliminarMecanico() {
	fmt.Print("ID del mecánico a eliminar: ")
	id, _ := strconv.Atoi(leerLinea())

	for i := range mecanicos {
		if mecanicos[i].ID == id {
			mecanicos = append(mecanicos[:i], mecanicos[i+1:]...)
			totalPlazas = len(mecanicos) * 2 // Recalcular plazas
			fmt.Println("Mecánico eliminado correctamente")
			return
		}
	}
	fmt.Println("Mecánico no encontrado")
}

func cambiarEstadoMecanico() {
	fmt.Print("ID del mecánico: ")
	id, _ := strconv.Atoi(leerLinea())

	for i := range mecanicos {
		if mecanicos[i].ID == id {
			fmt.Println("1. Dar de alta")
			fmt.Println("2. Dar de baja")
			fmt.Print("Opción: ")
			opcion := leerLinea()

			if opcion == "1" {
				mecanicos[i].Activo = true
				fmt.Println("Mecánico dado de alta")
			} else if opcion == "2" {
				mecanicos[i].Activo = false
				fmt.Println("Mecánico dado de baja")
			}
			totalPlazas = len(mecanicos) * 2
			return
		}
	}
	fmt.Println("Mecánico no encontrado")
}

// 8: GESTIÓN DE PLAZAS DEL TALLER

func asignarVehiculoATaller() {
	fmt.Println("\n--- ASIGNAR VEHÍCULO A TALLER ---")
	fmt.Printf("Plazas disponibles: %d de %d\n", totalPlazas-plazasOcupadas, totalPlazas)

	// Verificar disponibilidad
	if plazasOcupadas >= totalPlazas {
		fmt.Println("No hay plazas disponibles en el taller")
		return
	}

	fmt.Print("Matrícula del vehículo: ")
	matricula := leerLinea()

	for i := range vehiculos {
		if vehiculos[i].Matricula == matricula {
			if vehiculos[i].EnTaller {
				fmt.Println("El vehículo ya está en el taller")
				return
			}
			// Asignar vehículo al taller
			vehiculos[i].EnTaller = true
			plazasOcupadas++
			fmt.Println("Vehículo asignado al taller correctamente")
			return
		}
	}
	fmt.Println("Vehículo no encontrado")
}

func verEstadoTaller() {
	fmt.Println("\n========================================")
	fmt.Println("       ESTADO DEL TALLER")
	fmt.Println("========================================")
	fmt.Printf("Total de plazas: %d\n", totalPlazas)
	fmt.Printf("Plazas ocupadas: %d\n", plazasOcupadas)
	fmt.Printf("Plazas libres: %d\n", totalPlazas-plazasOcupadas)
	fmt.Println("\nVehículos en el taller:")

	hayVehiculos := false
	for _, v := range vehiculos {
		if v.EnTaller {
			hayVehiculos = true
			fmt.Printf("  - %s (%s %s) | Cliente ID: %d\n",
				v.Matricula, v.Marca, v.Modelo, v.IDCliente)
		}
	}
	if !hayVehiculos {
		fmt.Println("  No hay vehículos en el taller")
	}
}

// 9: CONSULTAS Y REPORTES

func menuConsultas() {
	for {
		limpiarPantalla()
		fmt.Println("\n--- CONSULTAS Y LISTADOS ---")
		fmt.Println("1. Listar incidencias de un vehículo")
		fmt.Println("2. Listar vehículos de un cliente")
		fmt.Println("3. Listar mecánicos disponibles")
		fmt.Println("4. Listar incidencias de un mecánico")
		fmt.Println("5. Listar clientes con vehículos en taller")
		fmt.Println("6. Listar todas las incidencias y su estado")
		fmt.Println("7. Volver")
		fmt.Print("Opción: ")

		opcion := leerLinea()

		switch opcion {
		case "1":
			listarIncidenciasVehiculo()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "2":
			listarVehiculosCliente()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "3":
			listarMecanicosDisponibles()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "4":
			listarIncidenciasMecanico()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "5":
			listarClientesEnTaller()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "6":
			listarTodasIncidencias()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "7":
			return
		default:
			fmt.Println("Opción no válida")
			fmt.Print("Presione Enter para continuar...")
			leerLinea()
		}
	}
}

func listarIncidenciasVehiculo() {
	fmt.Print("Matrícula del vehículo: ")
	matricula := leerLinea()

	// Buscar el vehículo y obtener su ID de incidencia
	var idIncidencia int
	encontrado := false
	for _, v := range vehiculos {
		if v.Matricula == matricula {
			idIncidencia = v.IDIncidencia
			encontrado = true
			break
		}
	}

	if !encontrado {
		fmt.Println("Vehículo no encontrado")
		return
	}

	// Mostrar las incidencias del vehículo
	fmt.Printf("\nIncidencias del vehículo %s:\n", matricula)
	hayIncidencias := false
	for _, inc := range incidencias {
		if inc.ID == idIncidencia {
			hayIncidencias = true
			fmt.Printf("  ID: %d | Tipo: %s | Estado: %s | Prioridad: %s\n",
				inc.ID, inc.Tipo, inc.Estado, inc.Prioridad)
		}
	}
	if !hayIncidencias {
		fmt.Println("  No hay incidencias registradas")
	}
}

func listarVehiculosCliente() {
	fmt.Print("ID del cliente: ")
	id, _ := strconv.Atoi(leerLinea())

	fmt.Printf("\nVehículos del cliente %d:\n", id)
	hayVehiculos := false
	for _, v := range vehiculos {
		if v.IDCliente == id {
			hayVehiculos = true
			fmt.Printf("  %s - %s %s (En taller: %v)\n",
				v.Matricula, v.Marca, v.Modelo, v.EnTaller)
		}
	}
	if !hayVehiculos {
		fmt.Println("  El cliente no tiene vehículos registrados")
	}
}

func listarMecanicosDisponibles() {
	fmt.Println("\n--- MECÁNICOS DISPONIBLES ---")

	// Crear un mapa con los mecánicos que están ocupados
	asignados := make(map[int]bool)
	for _, inc := range incidencias {
		if inc.Estado != "cerrada" {
			for _, idMec := range inc.MecanicosID {
				asignados[idMec] = true
			}
		}
	}

	// Mostrar mecánicos que están activos y no asignados
	hayDisponibles := false
	for _, m := range mecanicos {
		if m.Activo && !asignados[m.ID] {
			hayDisponibles = true
			fmt.Printf("  ID: %d | %s | Especialidad: %s\n",
				m.ID, m.Nombre, m.Especialidad)
		}
	}
	if !hayDisponibles {
		fmt.Println("  No hay mecánicos disponibles")
	}
}

func listarIncidenciasMecanico() {
	fmt.Print("ID del mecánico: ")
	id, _ := strconv.Atoi(leerLinea())

	fmt.Printf("\nIncidencias asignadas al mecánico %d:\n", id)
	hayIncidencias := false
	for _, inc := range incidencias {
		for _, idMec := range inc.MecanicosID {
			if idMec == id {
				hayIncidencias = true
				fmt.Printf("  ID: %d | Tipo: %s | Estado: %s | Prioridad: %s\n",
					inc.ID, inc.Tipo, inc.Estado, inc.Prioridad)
				break
			}
		}
	}
	if !hayIncidencias {
		fmt.Println("  No hay incidencias asignadas a este mecánico")
	}
}

func listarClientesEnTaller() {
	fmt.Println("\n--- CLIENTES CON VEHÍCULOS EN TALLER ---")

	// Crear un mapa con los IDs de clientes que tienen vehículos en taller
	clientesEnTaller := make(map[int]bool)
	for _, v := range vehiculos {
		if v.EnTaller {
			clientesEnTaller[v.IDCliente] = true
		}
	}

	// Mostrar información de esos clientes
	hayClientes := false
	for _, c := range clientes {
		if clientesEnTaller[c.ID] {
			hayClientes = true
			fmt.Printf("  ID: %d | Nombre: %s | Tel: %s\n",
				c.ID, c.Nombre, c.Telefono)
		}
	}
	if !hayClientes {
		fmt.Println("  No hay clientes con vehículos en el taller")
	}
}

func listarTodasIncidencias() {
	fmt.Println("\n--- TODAS LAS INCIDENCIAS DEL TALLER ---")
	if len(incidencias) == 0 {
		fmt.Println("No hay incidencias registradas")
		return
	}

	for _, inc := range incidencias {
		fmt.Printf("\nID: %d\n", inc.ID)
		fmt.Printf("  Tipo: %s\n", inc.Tipo)
		fmt.Printf("  Prioridad: %s\n", inc.Prioridad)
		fmt.Printf("  Estado: %s\n", inc.Estado)
		fmt.Printf("  Descripción: %s\n", inc.Descripcion)
		fmt.Printf("  Mecánicos asignados: %v\n", inc.MecanicosID)
	}
}

// 10: MENÚ PRINCIPAL

func mostrarMenuPrincipal() {
	for {
		limpiarPantalla()
		fmt.Println("\n========================================")
		fmt.Println("    SISTEMA DE GESTIÓN DE TALLER")
		fmt.Println("========================================")
		fmt.Println("1. Gestión de Clientes")
		fmt.Println("2. Gestión de Vehículos")
		fmt.Println("3. Gestión de Incidencias")
		fmt.Println("4. Gestión de Mecánicos")
		fmt.Println("5. Asignar vehículo a plaza del taller")
		fmt.Println("6. Ver estado del taller")
		fmt.Println("7. Consultas y Listados")
		fmt.Println("8. Salir")
		fmt.Print("\nSeleccione una opción: ")

		opcion := leerLinea()

		switch opcion {
		case "1":
			menuClientes()
		case "2":
			menuVehiculos()
		case "3":
			menuIncidencias()
		case "4":
			menuMecanicos()
		case "5":
			asignarVehiculoATaller()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "6":
			verEstadoTaller()
			fmt.Print("\nPresione Enter para continuar...")
			leerLinea()
		case "7":
			menuConsultas()
		case "8":
			limpiarPantalla()
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción no válida")
			fmt.Print("Presione Enter para continuar...")
			leerLinea()
		}
	}
}


func main() {
	inicializarDatos()
	mostrarMenuPrincipal()
}