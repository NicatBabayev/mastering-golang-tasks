package buffered_unbuffered_channels

func RunTask6() {
	ch := make(chan string)
	ch <- "Hello"
}
