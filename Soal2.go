package main

 import (
         "fmt"
         "math"
 )

 var (
         weight, height, bmi float64
 )

 func main() {
         fmt.Println("Masukkan berat badan dalam kilograM : ")
         fmt.Scanf("%f", &weight)

         fmt.Println("Masukkan tinggi dalam meter : ")
         fmt.Scanf("%f", &height)

         fmt.Println("Berat Anda : ", weight)
         fmt.Println("Tinggi Anda: ", height)

         bmi = weight / math.Pow(height, 2)

         fmt.Println("Skor BMI kamu adalah : ", bmi)
         fmt.Print("Kategori kamu adalah : ")

         if bmi < float64(18.5) {
                 fmt.Println("Kekurangan berat badan")
         } else if bmi < 25 {
                 fmt.Println("Berat badan normal")
         } else if bmi < 30 {
                 fmt.Println("Kelebihan Berat Badan")
         } else {
                 fmt.Println("Obesitas")
         }

         // Menghitung standar bmi yakni 25 poin
         normalWeight := 25 * math.Pow(height, 2)
         delta := weight - normalWeight
         var intNormalWeight int = int(normalWeight)

         fmt.Printf("Berat badan normal untuk tinggi Anda adalah : %0.2v kilograms.\n", intNormalWeight)

         if (delta > 0) && (bmi > 30) {
                 fmt.Printf("You need to reduce %0.2v kilograms.\n", math.Abs(delta))
         }

         if (delta < 0) && (bmi < float64(18.5)) {
                 fmt.Printf("You need to increase %0.2v kilograms.\n", math.Abs(delta))
         }

 }