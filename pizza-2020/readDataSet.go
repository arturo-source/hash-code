package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func read(filePath string) (uint32, []uint32, error) {
	maxSlices := uint32(0)
	typeSlices := []uint32{}
	file, err := os.Open(filePath)
	if err != nil {
		return maxSlices, typeSlices, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	firstLine, _, err := reader.ReadLine()
	if err != nil {
		return maxSlices, typeSlices, err
	}

	information := strings.Fields(string(firstLine))
	maxSlices64, err := strconv.ParseUint(information[0], 10, 32)
	if err != nil {
		return maxSlices, typeSlices, err
	}
	maxSlices = uint32(maxSlices64)
	typesInt, err := strconv.Atoi(information[1])
	if err != nil {
		return maxSlices, typeSlices, err
	}
	typeSlices = make([]uint32, 0, typesInt)

	isPrefix := true
	var secondLine []byte
	var auxLine []byte
	for isPrefix {
		auxLine, isPrefix, err = reader.ReadLine()
		if err != nil {
			return maxSlices, typeSlices, err
		}
		secondLine = append(secondLine, auxLine...)
	}
	information = strings.Fields(string(secondLine))
	for _, sliceLen := range information {
		sliceLen64, err := strconv.ParseUint(sliceLen, 10, 32)
		if err != nil {
			return maxSlices, typeSlices, err
		}
		typeSlices = append(typeSlices, uint32(sliceLen64))
	}

	return maxSlices, typeSlices, nil
}
