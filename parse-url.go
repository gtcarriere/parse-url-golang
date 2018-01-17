// parse-url.go
// A simple URL parser written in Go.
//
// Copyright (C) 2018  Gabriel Carriere
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main


//Imports requiered packages
import (
	"fmt"
    "strings"
    "os"
    "bufio"
)


func main() {
    
    // Code to get input from user (URL to parse)
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Please enter the URL to parse: ")
    url, _ := reader.ReadString('\n')
  
   // If URL is NOT equal to nothing, then parse it, and print the parsed URL.
   if url != "" {
       // Define the variable parsed_url and use strings.Replace to replace slashes with /n
       parsed_url := strings.Replace(url, "/", "\n", -1) 
       // Print one blank line before printing parsed url
       fmt.Println("\n")
       // Print parsed URL
       fmt.Println(parsed_url) 
   
       } else {
    // If URL IS empty, then ask user to enter a URL to parse. 
    fmt.Println("Please enter a valid URL to parse.")
       
       
    }
}