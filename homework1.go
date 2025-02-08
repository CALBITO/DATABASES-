// --------------------------------------------------
//
// HOMEWORK 1
//
// Due: Fri, Feb 7, 2025 (23h59)
//
// Name: Carlos Pena Acosta
//
// Email: cpenaacosta@olin.edu
//
// Remarks, if any:
// Only resource i utilized were format configurations on Github and syntax verification via AI models
//
// --------------------------------------------------
//
// Please fill in this file with your solutions and submit it
//
// The functions below are stubs that you should replace with your
// own implementation.
//
// PLEASE DO NOT CHANGE THE SIGNATURE IN THE STUBS BELOW.
// Doing so makes it impossible for me to test your code.
//
// --------------------------------------------------

package main

import (
	"fmt"
)

func Clamp(min int, max int, v int) int {
	if min <= max {
		if v < min {
			return min
		} else if v > max {
			return max
		}
		return v
	} else {
		if v < max {
			return max
		} else if v > min {
			return min
		}
		return v
	}
}

func Interpolate(min float32, max float32, v float32) float32 {
	return min + v*(max-min)
}

func Spaces(n int) string {
	var n_string string
	for i := 0; i < n; i++ {
		n_string += " "
	}
	return n_string
}

func PadLeft(s string, n int) string {
	if len(s) >= n {
		return s
	}
	var pad string
	for i := 0; i < n-len(s); i++ {
		pad += " "
	}
	return pad + s
}

func PadRight(s string, n int) string {
	if len(s) >= n {
		return s
	}
	var pad string
	for i := 0; i < n-len(s); i++ {
		pad += " "
	}
	return s + pad
}

func PadBoth(s string, n int) string {
	if len(s) >= n {
		return s
	}
	totalPadding := n - len(s)
	leftPadding := totalPadding / 2
	rightPadding := totalPadding - leftPadding
	return fmt.Sprintf("%*s%s%*s", leftPadding, "", s, rightPadding, "")
}

func NewVec(n1 float32, n2 float32, n3 float32) [3]float32 {
	return [3]float32{n1, n2, n3}
}

func ScaleVec(sc float32, v1 [3]float32) [3]float32 {
	return [3]float32{
		v1[0] * sc,
		v1[1] * sc,
		v1[2] * sc,
	}
}

func AddVec(v1 [3]float32, v2 [3]float32) [3]float32 {
	return [3]float32{
		v1[0] + v2[0],
		v1[1] + v2[1],
		v1[2] + v2[2],
	}
}

func DotProd(v1 [3]float32, v2 [3]float32) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func NewMat(r1 [3]float32, r2 [3]float32, r3 [3]float32) [9]float32 {
	return [9]float32{
		r1[0], r1[1], r1[2],
		r2[0], r2[1], r2[2],
		r3[0], r3[1], r3[2],
	}
}

func ScaleMat(sc float32, m [9]float32) [9]float32 {
	var result [9]float32
	for i := 0; i < 9; i++ {
		result[i] = m[i] * sc
	}
	return result
}

func AddMat(m1 [9]float32, m2 [9]float32) [9]float32 {
	var result [9]float32
	for i := 0; i < 9; i++ {
		result[i] = m1[i] + m2[i]
	}
	return result
}

func TransposeMat(m [9]float32) [9]float32 {
	var result [9]float32
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[j*3+i] = m[i*3+j]
		}
	}
	return result
}

func MultMat(m1 [9]float32, m2 [9]float32) [9]float32 {
	var result [9]float32
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum := float32(0)
			for k := 0; k < 3; k++ {
				sum += m1[i*3+k] * m2[k*3+j]
			}
			result[i*3+j] = sum
		}
	}
	return result
}

func main() {
	var f string = "Expected = %v\n     Got = %v\n"
	var wrap func(string) string = func(s string) string { return "'" + s + "'" }
	fmt.Println("You can write some sample tests for yourself here. Here are some to get started.")

	fmt.Println("****** Clamp ***************************************")
	fmt.Printf(f, 10.0, Clamp(10.0, 20.0, 5.0))

	fmt.Println("****** Interpolate ***************************************")
	fmt.Printf(f, 10.0, Interpolate(10.0, 20.0, 0))

	fmt.Println("****** Spaces ***************************************")
	fmt.Printf(f, "'          '", wrap(Spaces(10)))

	fmt.Println("****** PadLeft ***************************************")
	fmt.Printf(f, "'      test'", wrap(PadLeft("test", 10)))

	fmt.Println("****** PadRight ***************************************")
	fmt.Printf(f, "'test      '", wrap(PadRight("test", 10)))

	fmt.Println("****** PadBoth ***************************************")
	fmt.Printf(f, "'   test    '", wrap(PadBoth("test", 10)))

	var v1 [3]float32 = [3]float32{1.0, 2.0, 3.0}

	fmt.Println("****** NewVec ***************************************")
	fmt.Printf(f, [3]float32{0, 10, 20}, NewVec(0, 10, 20))

	fmt.Println("****** ScaleVec ***************************************")
	fmt.Printf(f, [3]float32{2, 4, 6}, ScaleVec(2.0, v1))

	fmt.Println("****** AddVec ***************************************")
	fmt.Printf(f, [3]float32{2, 4, 6}, AddVec(v1, v1))

	fmt.Println("****** DotProd ***************************************")
	fmt.Printf(f, 14, DotProd(v1, v1))

	var v31 [3]float32 = [3]float32{1, 2, 3}
	var v32 [3]float32 = [3]float32{4, 5, 6}
	var v33 [3]float32 = [3]float32{7, 8, 9}
	var m1 [9]float32 = [9]float32{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("****** NewMat ***************************************")
	fmt.Printf(f, [9]float32{1, 2, 3, 4, 5, 6, 7, 8, 9}, NewMat(v31, v32, v33))

	fmt.Println("****** ScaleMat ***************************************")
	fmt.Printf(f, [9]float32{2, 4, 6, 8, 10, 12, 14, 16, 18}, ScaleMat(2.0, m1))

	fmt.Println("****** TransposeMat ***************************************")
	fmt.Printf(f, [9]float32{1, 4, 7, 2, 5, 8, 3, 6, 9}, TransposeMat(m1))

	fmt.Println("****** AddMat ***************************************")
	fmt.Printf(f, [9]float32{2, 4, 6, 8, 10, 12, 14, 16, 18}, AddMat(m1, m1))

	fmt.Println("****** MultMat ***************************************")
	fmt.Printf(f, [9]float32{30, 36, 42, 66, 81, 96, 102, 126, 150}, MultMat(m1, m1))
}
