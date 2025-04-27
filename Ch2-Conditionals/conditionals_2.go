package main

func main() {

	//switch statement
	grade := "d"
	switch grade {
	case "A":
		println("You have scored above 85!")
	case "B":
		println("You have scored above 65!")
	case "C":
		println("You have scored above 55!")
	default:
		println("You have Failed lol")
	}

	//switch without expression(acts like if else chain)

	age := 20
	switch {
	case age < 13:
		println("Child")
	case age < 18:
		println("Teenager")
	default:
		println("Adult")

	}

	//fallthrough in switch
	// Normally, after a case matches, Go does not automatically continue to the next case.
	// If you want it to continue even after matching, use fallthrough.
	num := 2

	switch num {
	case 2:
		println("Two")
		fallthrough
	case 4:
		println("Four")
		fallthrough
	case 6:
		println("Six")
	}

}
