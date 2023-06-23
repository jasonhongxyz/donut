package main

import (
	"fmt"
	"math"
)

const (
	screenSize   = 40
	pi           = 3.14
	thetaRefresh = 0.07
	phiRefresh   = 0.02
	lumGrad      = ".,-~:;=!*#$@"
)

var (
	R1      = 1
	R2      = 2
	K1      = (screenSize * K2 * 3) / (8 * (R1 + R2))
	K2      = 10
	A       = 0.02
	B       = 0.02
	zbuffer = [][]float64{}
	output  = [][]string{}
	light   = [3]float64{0, 1, -1}
)

func MatMult(mat1 [3]float64, mat2 [3][3]float64) [3]float64 {
	var ret [3]float64

	ret[0] = (mat1[0] * mat2[0][0]) + (mat1[1] * mat2[0][1]) + (mat1[2] * mat2[0][2])

	ret[1] = (mat1[0] * mat2[1][0]) + (mat1[1] * mat2[1][1]) + (mat1[2] * mat2[1][2])

	ret[2] = (mat1[0] * mat2[2][0]) + (mat1[1] * mat2[2][1]) + (mat1[2] * mat2[2][2])

	return ret
}

func DotProd(v1, v2 [3]float64) float64 {
	ret := (v1[0] * v2[0]) + (v1[1] * v2[1]) + (v1[2] * v2[2])
	return ret
}

func CreateRotationMatrix(kind string, sinVal, cosVal float64) [3][3]float64 {
	var matrix [3][3]float64
	switch kind {
	case "X":
		matrix[0] = [3]float64{1, 0, 0}
		matrix[1] = [3]float64{0, cosVal, sinVal}
		matrix[2] = [3]float64{0, -1.0 * sinVal, cosVal}
	case "Y":
		matrix[0] = [3]float64{cosVal, 0, sinVal}
		matrix[1] = [3]float64{0, 1, 0}
		matrix[2] = [3]float64{-1.0 * sinVal, 0, cosVal}
	case "Z":
		matrix[0] = [3]float64{cosVal, sinVal, 0}
		matrix[1] = [3]float64{-1.0 * sinVal, cosVal, 0}
		matrix[2] = [3]float64{0, 0, 1}
	}

	return matrix
}

func GetRotation(a, b float64) {
	sinA := math.Sin(a)
	cosA := math.Cos(a)

	sinB := math.Sin(b)
	cosB := math.Cos(b)

	// Draw circle
	for theta := 0.0; theta < 2*pi; theta += thetaRefresh {
		sinTheta := math.Sin(theta)
		cosTheta := math.Cos(theta)

		// Revolve circle
		for phi := 0.0; phi < 2*pi; phi += phiRefresh {
			sinPhi := math.Sin(phi)
			cosPhi := math.Cos(phi)

			circle := [3]float64{float64(R2) + float64(R1)*cosTheta, float64(R1) * sinTheta, 0}
			yRot := CreateRotationMatrix("Y", sinPhi, cosPhi)
			xRot := CreateRotationMatrix("X", sinA, cosA)
			zRot := CreateRotationMatrix("Z", sinB, cosB)

			point := MatMult(MatMult(MatMult(circle, yRot), xRot), zRot)

			x := point[0]
			y := point[1]
			z := point[2] + float64(K2)
			ooz := 1.0 / z

			xprime := (screenSize / 2) + int(float64(K1)*ooz*x)
			yprime := (screenSize / 2) - int(float64(K1)*ooz*y)

			unitCircle := [3]float64{cosTheta, sinTheta, 0}
			lum := DotProd(MatMult(MatMult(MatMult(unitCircle, yRot), xRot), zRot), light)

			if lum > 0 {
				if ooz > zbuffer[xprime][yprime] {
					zbuffer[xprime][yprime] = ooz
					lumVal := int(lum * 8)
					output[xprime][yprime] = string(lumGrad[lumVal])
				}
			}
		}

	}

	fmt.Printf("\033[H\033[2J")

	for i := 0; i < screenSize; i++ {
		for j := 0; j < screenSize; j++ {
			fmt.Printf(output[j][i])
			output[j][i] = " "
			zbuffer[j][i] = 0.0
		}
		fmt.Printf("\n")
	}
}

func initBuffers() {
	zbuffer = make([][]float64, screenSize)
	for i := 0; i < screenSize; i++ {
		zbuffer[i] = make([]float64, screenSize)
	}
	output = make([][]string, screenSize)
	for i := 0; i < screenSize; i++ {
		output[i] = make([]string, screenSize)
		for j := 0; j < len(output[i]); j++ {
			output[i][j] = " "
		}
	}
}

func main() {
	initBuffers()

	for {
		GetRotation(A, B)
		A += 0.02
		B += 0.02
	}
}
