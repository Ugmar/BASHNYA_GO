package main

import (
	"flag"
	"fmt"
	"io"
	"bufio"
	"os"
	"dz4/uniq"
)

const (
	INIT_SIZE_SLICE = 10;
)

func CreateOption(option *uniq.Options){
	flag.BoolVar(&option.C, "c", false, "подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.")
	flag.BoolVar(&option.D, "d", false, "вывести только те строки, которые повторились во входных данных.")
	flag.BoolVar(&option.U, "u", false, "в вывести только те строки, которые не повторились во входных данных.")
	flag.BoolVar(&option.I, "i", false, "не учитывать регистр букв.")
	flag.IntVar(&option.F, "f", 0, "не учитывать первые n полей в строке")
	flag.IntVar(&option.S, "s", 0, "не учитывать первые m символов в строке")

	flag.Parse()

	option.Input = flag.Arg(0)
	option.Output = flag.Arg(1)
}

func Help(){
	flag.VisitAll(func(f *flag.Flag) {
            fmt.Printf("  -%-10s %s\n", f.Name, f.Usage)
        })
		fmt.Println("uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
}

func ParseOption(option uniq.Options){
	if option.C && option.D || option.C && option.U || option.D && option.U{
		Help()
		os.Exit(2)
	}
}

func CreateData(r io.Reader) ([]string){
	data := make([]string, 0, INIT_SIZE_SLICE)
	scanner := bufio.NewScanner(r)

	for scanner.Scan(){
		line := scanner.Text()
		data = append(data, line)
	}

	return data
}

func PrintData(w io.Writer, data []string){
	for i := range len(data){
		fmt.Fprintln(w, data[i])
	}
}

func main(){
	var options uniq.Options

	CreateOption(&options)
	ParseOption(options)

	var data []string

	if options.Input == ""{
		data = CreateData(os.Stdin)
	}else {
		file, err := os.Open(options.Input)

		if err != nil{
			fmt.Println("Ошибка при открытии файла")
			os.Exit(1)
		}
		defer file.Close()

		data = CreateData(file)
	}

	uniq_data := uniq.Uniq(data, options)

	if options.Output == ""{
		PrintData(os.Stdout, uniq_data)
	} else {
		file, err := os.Create(options.Output)

		if err != nil{
			fmt.Println("Ошибка при открытии файла")
			os.Exit(1)
		}
		defer file.Close()

		PrintData(file, uniq_data)
	}
}
