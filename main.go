package main

import "fmt"

func main() {
	fmt.Println("Выберите операцию:")
	fmt.Println("1. Сложение матриц")
	fmt.Println("2. Вычитание матриц")
	fmt.Println("3. Умножение матрицы на число")
	fmt.Println("4. Умножение матриц")

	var choice int
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		addMatrices()
	case 2:
		subtractMatrices()
	case 3:
		multiplyByScalar()
	case 4:
		multiplyMat()
	default:
		fmt.Println("Неверный выбор")
	}
}

func readMatrix(name string, n int) [][]float64 {
	matrix := make([][]float64, n)

	fmt.Printf("Введите элементы матрицы %s построчно:\n", name)
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}

func printMatrix(m [][]float64) {
	for _, row := range m {
		for _, v := range row {
			fmt.Printf("%.2f ", v)
		}
		fmt.Println()
	}
}
func addMatrices() {
	var n int
	fmt.Print("Введите размер матрицы (2 или 3): ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Ошибка: неверный ввод размера.")
		return
	}
	if n != 2 && n != 3 {
		fmt.Println("Поддерживаются только 2 или 3")
		return
	}

	a := readMatrix("A", n)
	b := readMatrix("B", n)

	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}

	fmt.Println("Результат сложения:")
	printMatrix(result)
}

func subtractMatrices() {
	var n int
	fmt.Print("Введите размер матрицы (2 или 3): ")
	fmt.Scan(&n)
	if n != 2 && n != 3 {
		fmt.Println("Поддерживаются только 2 или 3")
		return
	}

	a := readMatrix("A", n)
	b := readMatrix("B", n)

	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i][j] = a[i][j] - b[i][j]
		}
	}

	fmt.Println("Результат вычитания:")
	printMatrix(result)
}

func multiplyByScalar() {
	var n int
	fmt.Print("Введите размер матрицы (2 или 3): ")
	fmt.Scan(&n)
	if n != 2 && n != 3 {
		fmt.Println("Поддерживаются только 2 или 3")
		return
	}

	matrix := readMatrix("A", n)
	var scalar float64
	fmt.Print("Введите число: ")
	fmt.Scan(&scalar)

	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			result[i][j] = matrix[i][j] * scalar
		}
	}

	fmt.Println("Результат умножения на число:")
	printMatrix(result)
}
func multiplyMat() {
	var n int
	fmt.Print("Введите размер матрицы (2 или 3): ")
	fmt.Scan(&n)
	if n != 2 && n != 3 {
		fmt.Println("Поддерживаются только 2 или 3")
		return
	}

	a := readMatrix("A", n)
	b := readMatrix("B", n)

	result := make([][]float64, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	fmt.Println("Результат умножения матриц:")
	printMatrix(result)
}
