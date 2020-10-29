package mypack

//GetName is an example of funcion inside a package
//First letter in function name should be uppercase for function to be exported
//Functions, Types, or even Type fields which start with lowercase letters will
//be private
func GetName() string {
	name := "Муммий Тролль"
	return name
}
