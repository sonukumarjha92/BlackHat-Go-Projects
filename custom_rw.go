func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	// जादू यहाँ है: io.Copy सीधा reader से डेटा लेकर writer में डाल देगा
	io.Copy(&writer, &reader)
}

