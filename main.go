package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var id int
var users map[int]User

func crearUsuario() {
	clearConsole()
	fmt.Println(" Ingresa un valor de username")
	username := readLine()

	fmt.Println(" Ingresa un valor de email")
	email := readLine()

	fmt.Println(" Ingresa un valor de age")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir el string")
	}

	id++
	user := User{id, username, email, age}
	users[id] = user

	fmt.Println("Usuarios Creados")
}
func listarUsuario() {
	clearConsole()
	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}
	fmt.Println("\n")
	fmt.Println("Listado de Usuario")
}
func actualizarUsuario() {
	clearConsole()

	fmt.Print("Ingresa el id del usuario a actualizar: ")
	id, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir de un string a un entero.")
	}

	if _, ok := users[id]; ok {

		fmt.Print("Ingresa un nuevo username: ")
		username := readLine()
		fmt.Print("Ingresa un nuevo email: ")
		email := readLine()
		fmt.Print("Ingresa una nueva edad: ")
		age, err := strconv.Atoi(readLine())

		if err != nil {
			panic("No es posible convertir de un string a un entero")
		}

		fmt.Println("\nEste es el nuevo username: ", username)
		fmt.Println("Este es el nuevo email: ", email)
		fmt.Println("Este es la nueva edad: ", age)

		user := User{id, username, email, age}
		users[id] = user
		fmt.Println("Actualizar Usuario")
	}
}
func eliminarUsuario() {
	clearConsole()
	fmt.Println("Ingresar el id del usuario a eliminar")
	id, err := strconv.Atoi(readLine())

	if _, ok := users[id]; ok {
		delete(users, id)
	}

	if err != nil {
		panic("No se puede convertir el id")
	}
	fmt.Println("Usuario eliminado exitosamente")
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string {

	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor")
	} else {
		return strings.TrimSuffix(option, "\n")
	}

}

func main() {

	var option string
	users = make(map[int]User)

	reader = bufio.NewReader(os.Stdin)

	for {

		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")

		fmt.Print("Ingresar una opcion valida")

		option = readLine()
		if option == "quit" || option == "q" {
			break
		}

		switch option {
		case "a", "crear":
			crearUsuario()

		case "b", "listar":
			listarUsuario()

		case "c", "actualizar":
			actualizarUsuario()

		case "d", "eliminar":
			eliminarUsuario()
		default:
			fmt.Println("Option no valida")
		}
	}
	fmt.Println("Programa Finalizado")

}
