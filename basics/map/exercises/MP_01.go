package main

import "fmt"

func main()  {
	//Create a map of countries to their capitals.
	countries := make(map[string]string)
	countries["Cameroon"] = "Youande"
	countries["Nigeria"] = "Abuja"
	countries["Ghana"] = "Accra"
	countries["Togo"] = "Lome"
	countries["Gabon"] = "Libreville"
	countries["Chad"] = "N'Djamena"
	countries["Equatorial Guinea"] = "Malabo"
	countries["Central African Republic"] = "Bangui"
	countries["Democratic Republic of the Congo"] = "Kinshasa"
	countries["Republic of the Congo"] = "Brazzaville"
	countries["Sao Tome and Principe"] = "Sao Tome"
	countries["Guinea-Bissau"] = "Bissau"
	countries["Cape Verde"] = "Praia"
	countries["Liberia"] = "Monrovia"
	countries["Sierra Leone"] = "Freetown"
	countries["Guinea"] = "Conakry"
	countries["The Gambia"] = "Banjul"
	countries["Senegal"] = "Dakar"
	countries["United States"] = "Washington, D.C."
	countries["United Kingdom"] = "London"
	countries["France"] = "Paris"
	countries["Germany"] = "Berlin"
	countries["China"] = "Beijing"
	countries["India"] = "New Delhi"
	countries["Japan"] = "Tokyo"
	countries["Brazil"] = "Brasilia"
	countries["Russia"] = "Moscow"
	countries["South Africa"] = "Pretoria"
	countries["Egypt"] = "Cairo"
	countries["Kenya"] = "Nairobi"  
  
	for country, capital := range countries {
	  fmt.Println(country, capital)
	}
  }