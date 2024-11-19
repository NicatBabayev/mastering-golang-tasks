package app

import "fmt"

/*
	 func DisplayWelcomeMessage() {
		fmt.Println("  _____      _     _ _                          ")
		fmt.Println(" | ____|    | |   (_) |__  _ __ __ _ _ __ _   _ ")
		fmt.Println(" |  _| _____| |   | | '_ \\| '__/ _` | '__| | | |")
		fmt.Println(" | |__|_____| |___| | |_) | | | (_| | |  | |_| |")
		fmt.Println(" |_____|    |_____|_|_.__/|_|  \\__,_|_|   \\__, |")
		fmt.Println("                                          |___/ ")

}
*/
func DisplayWelcomeMessage() {
	fmt.Println(" ███████╗    ██╗     ██╗██████╗ ██████╗  █████╗ ██████╗ ██╗   ██╗")
	fmt.Println(" ██╔════╝    ██║     ██║██╔══██╗██╔══██╗██╔══██╗██╔══██╗╚██╗ ██╔╝")
	fmt.Println(" █████╗█████╗██║     ██║██████╔╝██████╔╝███████║██████╔╝ ╚████╔╝ ")
	fmt.Println(" ██╔══╝╚════╝██║     ██║██╔══██╗██╔══██╗██╔══██║██╔══██╗  ╚██╔╝  ")
	fmt.Println(" ███████╗    ███████╗██║██████╔╝██║  ██║██║  ██║██║  ██║   ██║   ")
	fmt.Println(" ╚══════╝    ╚══════╝╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ")

}

func DisplayCommandsMessage() {

	fmt.Println(" ██████╗ ██████╗ ███╗   ███╗███╗   ███╗ █████╗ ███╗   ██╗██████╗ ███████╗")
	fmt.Println("██╔════╝██╔═══██╗████╗ ████║████╗ ████║██╔══██╗████╗  ██║██╔══██╗██╔════╝")
	fmt.Println("██║     ██║   ██║██╔████╔██║██╔████╔██║███████║██╔██╗ ██║██║  ██║███████╗")
	fmt.Println("██║     ██║   ██║██║╚██╔╝██║██║╚██╔╝██║██╔══██║██║╚██╗██║██║  ██║╚════██║")
	fmt.Println("╚██████╗╚██████╔╝██║ ╚═╝ ██║██║ ╚═╝ ██║██║  ██║██║ ╚████║██████╔╝███████║")
	fmt.Println(" ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝ ╚══════╝")

}
func ReturnCommandsMenu() string {
	res := fmt.Sprintf("- [1] List All Books\n- [2] List Available Books\n- [3] List Borrowed Books\n- [4] List All Users\n- [5] Borrow Book\n- [6] Return Book\n- [7] Show Welcome Screen \n- [8] Quit\n")
	return res
}
