package main

import "fmt"

func main() {

	//  variables y arrays
	nombres := [3]string{"diego", "carol", "juan"}

	calificaciones := [3][3]float64{
		{90, 85, 78}, // diego
		{60, 55, 70}, // carol
		{95, 88, 92}, // juan
	}

	materias := [3]string{"Matematicas", "Espanol", "Ingles"}

	fmt.Println("=== SISTEMA DE CALIFICACIONES ===")
	fmt.Println()

	// mostrar califiaciones de cada wey
	for i := 0; i < 3; i++ {

		fmt.Println("Estudiante:", nombres[i])
		fmt.Println("--------------------------")

		// sacar promedio
		suma := 0.0
		for j := 0; j < 3; j++ {
			fmt.Printf("  %s: %.1f\n", materias[j], calificaciones[i][j])
			suma += calificaciones[i][j]
		}

		promedio := suma / 3

		fmt.Printf("  Promedio: %.2f\n", promedio)

		// aprobdo o reprobado
		if promedio >= 70 {
			fmt.Println("  Estado: APROBADO")
		} else {
			fmt.Println("  Estado: REPROBADO")
		}

		fmt.Println()
	}
}