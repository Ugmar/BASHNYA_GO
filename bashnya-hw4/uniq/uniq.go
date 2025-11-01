package uniq

import (
	"strings"
	"strconv"
)


type Options struct{
	C bool
	D bool
	U bool
	F int
	S int
	I bool
	Input string
	Output string
}

func skipfields(line string, n int) string{
	if n <= 0{
		return line
	}

	fields := strings.Fields(line)

	if n >= len(fields){
		return line
	}

	return strings.Join(fields[n:], " ")
}

func skipsymbol(line string, n int) string{
	if n <= 0 || n >= len(line){
		return line
	}

	return line[n:]
}

func Uniq(data []string, options Options) []string{
	map_uniq_string := make(map[string]string, len(data))
	map_uniq_count := make(map[string]int, len(data))

	for i, line := range data{
		if options.F != 0{
			line = skipfields(line, options.F)
		}

		if options.S != 0{
			line = skipsymbol(line, options.S)
		}

		if options.I{
			line = strings.ToLower(line)
		}

		_, exist := map_uniq_string[line]

		if !exist{
			map_uniq_string[line] = data[i]
		}

		map_uniq_count[line] += 1
	}

	data_uniq := make([]string, 0, len(data))

	if options.C{
		for key, value := range map_uniq_string{
			data_uniq = append(data_uniq, strconv.Itoa(map_uniq_count[key]) + " " + value)
		}
	} else if options.D{
		for key, value := range map_uniq_count{
			if value != 1{
				data_uniq = append(data_uniq, map_uniq_string[key])
			}
		}
	} else if options.U{
		for key, value := range map_uniq_count{
			if value == 1{
				data_uniq = append(data_uniq, map_uniq_string[key])
			}
		}
	} else {
		for _, value := range map_uniq_string{
			data_uniq = append(data_uniq, value)
		}
	}
	
	return data_uniq
}
