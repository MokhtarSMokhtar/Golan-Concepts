package basic

/*
In Go, applications are organized into packages, which are collections of source files
located in the same directory.Each source file within a directory must begin with the
same package name. By convention,the package name typically matches the name of the directory
where the files are stored.
*/

/*
Go provides an extensive standard library of packages
which you can use in your program using the "import" keyword.
Standard library packages are imported using their name
example
.*/
import "fmt"

func sayHello() {
	fmt.Println("hello world")
}
