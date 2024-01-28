package main

import (
	"fmt"
	"strconv"
	"strings"
	// "sort"
)

type Cipher interface {
	Encrypt() interface{}
}

type StringCipher string

// Encrypt method for String
func (s StringCipher) Encrypt() interface{} {
	mapping := map[byte]int{
		'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5,
		'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10,
		'K': 11, 'L': 12, 'M': 13, 'N': 14, 'O': 15,
		'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20,
		'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26,
	}
	var result []int
	for _, char := range s {
		if val, ok := mapping[byte(char)]; ok {
			result = append(result, val)
		}
	}
	return result
}

type IntArrayCipher []int

// Encrypt method for Integer Array
func (arr IntArrayCipher) Encrypt() interface{} {
	for i := range arr {
		if arr[i]%2 == 0 {
			arr[i] =arr[i]/ 2
		} else {
			arr[i] = 3*arr[i] + 1
		}
	}
	return arr
}

type MapCipher map[rune]int

// Encrypt method for Map
func (m MapCipher) Encrypt() interface{} {
	// keys := make([]rune, len(m))
	keys := []rune{}
	for key := range m {
		keys = append(keys, key)
	}
	result := make([]int, len(m))
	for i, key := range keys {
		result[i] = int(key) + m[key]
	}
	return result
}

func main() {
	var a = 1
	for a==1{
		var repeat string
		fmt.Println("Do you want to continue to the Program:")
		fmt.Scanln(&repeat)
		if repeat== "y" || repeat== "Y"{
			var input string
			fmt.Println("Enter input (string, integer array, or map):")
			fmt.Scanln(&input)

			var cipher Cipher

			if strings.Contains(input, ",") && !strings.Contains(input, ":") {
				arr := strings.Split(input, ",")
				intArr := make([]int, len(arr))
				for i, numStr := range arr {
					if i == 0 {
						num, _ := strconv.Atoi(strings.ReplaceAll(numStr, "[", ""))
						intArr[i] = num
					} else if i == len(arr)-1 {
						num, _ := strconv.Atoi(strings.ReplaceAll(numStr, "]", ""))
						intArr[i] = num
					} else {
						num, _ := strconv.Atoi(numStr)
						intArr[i] = num
					}					
				}
				cipher = IntArrayCipher(intArr)

			} else if strings.Contains(input, ":") {
				pairs := strings.Split(input, ",")
				m := make(MapCipher)
				for i, pair := range pairs {
					if i == 0 {
						keyVal := strings.Split(strings.ReplaceAll(pair, "[", ""), ":")
						key := rune(keyVal[0][0])
						val, _ := strconv.Atoi(keyVal[1])
						m[key] = val
					} else if i == len(pairs)-1 {
						keyVal := strings.Split(strings.ReplaceAll(pair, "]", ""), ":")
						key := rune(keyVal[0][0])
						val, _ := strconv.Atoi(keyVal[1])
						m[key] = val
					} else {
						keyVal := strings.Split(pair, ":")
						key := rune(keyVal[0][0])
						val, _ := strconv.Atoi(keyVal[1])
						m[key] = val
					}	
				}
				cipher = m
			} else {
				cipher = StringCipher(input)
			}

			// Encrypting
			result := cipher.Encrypt()
			fmt.Println("Encrypted result:", result)
		} else {
			break
		}

	}
}
