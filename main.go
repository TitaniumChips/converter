package main

import (
  convert "converter/convert"
  "fmt"
  "os"
)

func main() {

  var conversion int

  fmt.Print("\n\n\nWelcome to the convertor!\n\nbelow is a list of conversions,\nplease input the number corresponding to the conversion you would like.\n\n\n")
  fmt.Println("1 - MM -> Thousandths\n2 - Thousandths -> MM\n3 - Celsisus -> Fahrenheit\n4 - Fahrenhit -> Celsius\n5 - Radius -> Circle Area\n6 - Radius -> Circle Circumference")
  fmt.Print("\nInput Conversion: ")
  fmt.Scan(&conversion)

  if conversion < 1 ||  conversion > 6 {
    fmt.Println("\n\nERROR, not a valid conversion")
    os.Exit(1)
  }

  var input float64

  fmt.Print("\ninput units(no fractions or symbols): ")
  fmt.Scan(&input)

  if conversion == 1 {
    fmt.Println((convert.MM(input)), "Thousandths")
  } else if conversion == 2 {
    fmt.Println((convert.Thousandths(input)), "MM")
  } else if conversion == 3 {
    fmt.Println((convert.Celsius(input)), "Fahrenheit")
  } else if conversion == 4 {
    fmt.Println((convert.Fahrenheit(input)), "Celsius")
  } else if conversion == 5 {
    fmt.Println((convert.RadiusArea(input)), "Units Squared")
  } else if conversion == 6 {
    fmt.Println((convert.RadiusCircumference(input)), "Units")
  } else {
     fmt.Println("\n\nERROR!")
     os.Exit(1)
   }
}
