package main


import ("fmt";
	"os";
	"io/ioutil"
)


func short_wayto_read(file_path string) {
	bs, err := ioutil.ReadFile(file_path) 
	if err != nil {
		fmt.Println("Not able to read the file")
		panic ("This is pethentiuc man")
	}
	str := string(bs)
	fmt.Println("This is the shorter way to read")
	fmt.Println(str)
}

func create_file() {
	file, err := os.Create("go_created_file.txt")
	if err != nil {
		fmt.Println("Not able to read the file")
		panic ("This is pethentiuc man")
	}
	defer file.Close()
	file.WriteString("Go is awsome")
	
}


func main() {
	file, err := os.Open("need_to_read.txt")
	if err != nil {
		fmt.Println("Not able to read the file")
		panic ("This is pethentiuc man")
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Not able to get file stat")
		panic ("This is pethentiuc man")
	}
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Not able to read the file")
		panic ("This is pethentiuc man")
	}
	str := string(bs)
	fmt.Println(str)
	fmt.Println(stat.Size())
	short_wayto_read("need_to_read.txt")
	create_file()
}
