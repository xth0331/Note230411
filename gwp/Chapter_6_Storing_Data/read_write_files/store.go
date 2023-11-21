package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("Hello World!\n")
	err := os.WriteFile("gwp/Chapter_6_Storing_Data/read_write_files/data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := os.ReadFile("gwp/Chapter_6_Storing_Data/read_write_files/data1")
	fmt.Print(string(read1))
	file1, _ := os.Create("gwp/Chapter_6_Storing_Data/read_write_files/data2")
	defer file1.Close()
	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("gwp/Chapter_6_Storing_Data/read_write_files/data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
