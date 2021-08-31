package mymain

import (
    // "fmt"
    "regexp"
)

// func main() {
//     fmt.Println(isEmailValid("test44@gmail.com"))         // true 
// 	fmt.Println(isEmailValid("bad-email"))               // false
// 	fmt.Println(isEmailValid("test44$@gmail.com"))      // false
// 	fmt.Println(isEmailValid("test-email.com"))        // false
// 	fmt.Println(isEmailValid("test+email@test.com"))  // true
// 	fmt.Println(isEmailValid("test+email@tes"))  // true
// 	fmt.Println(isEmailValid("tesT44@gmail.com"))         // true 
// }


// isEmailValid checks if the email provided is valid by regex.
func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9A-Z._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return emailRegex.MatchString(e)
}