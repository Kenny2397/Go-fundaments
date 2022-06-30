package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarAlista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	for _, task := range t.tasks {
		fmt.Println("nombre:", task.nombre)
	}
}

func (t *taskList) imprimirListaCompletados() {

	for _, task := range t.tasks {
		if task.completado == true {
			fmt.Println("nombre:", task.nombre)

		} else {
			continue
		}
	}
}

type task struct {
	nombre      string
	description string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescription(description string) {
	t.description = description
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:      "Completar mi curso de Go",
		description: "Description",
		completado:  false,
	}

	t2 := &task{
		nombre:      "Completar mi curso de Python",
		description: "Description",
		completado:  false,
	}

	t3 := &task{
		nombre:      "Completar mi curso de Nodejs",
		description: "Description",
		completado:  false,
	}
	// // fmt.Println(t)
	// fmt.Printf("%+v\n", t)

	// t.marcarCompleta()
	// t.actualizarNombre("actualizando nombre")
	// t.actualizarDescription("nueva description")

	// fmt.Printf("%+v\n", t)

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	// fmt.Println(lista.tasks[0])
	// fmt.Println(lista.tasks[1])
	// // fmt.Println(lista.tasks[2])

	lista.agregarAlista(t3)
	// fmt.Println(lista.tasks[2])

	// lista.eliminarDeLista(1)
	// fmt.Println(len(lista.tasks))

	// for i:=0;i<len(lista.tasks);i++ {
	// 	fmt.Println("index:", i, "nombre:", lista.tasks[i].nombre)
	// }

	// for index, task := range lista.tasks {
	// 	fmt.Println("index:", index, "nombre:", task.nombre)
	// }

	// for i:=0;i<10;i++{
	// 	if i == 5 {
	// 		break
	// 	}else if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	lista.imprimirLista()

	lista.tasks[0].completado = true
	lista.tasks[1].marcarCompleta()

	// lista.imprimirLista()

	fmt.Println("Tareas completas")
	lista.imprimirListaCompletados()

	//
	t4 := &task{
		nombre:      "Completar mi curso de Go",
		description: "Description",
		completado:  false,
	}

	t5 := &task{
		nombre:      "Completar mi curso de Python",
		description: "Description",
		completado:  false,
	}

	lista2 := &taskList{
		tasks: []*task{
			t4,t5,
		},
	}

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Kenny"] = lista
	mapaTareas["Paul"]= lista2


	fmt.Println("Tareas de Kenny")
	mapaTareas["Kenny"].imprimirLista()

	fmt.Println("Tareas de Paul")
	mapaTareas["Paul"].imprimirLista()

}
