package services

import (
	"fmt"
	"sort"
)

// Set default values
var packageSizes = []int{250, 500, 1000, 2000, 5000}

func init() {
	// reverse sort package scales
	sort.Sort(sort.Reverse(sort.IntSlice(packageSizes)))
}

func CalculatePackagesFor(itemsCount int) (map[int]int, error) {
	if itemsCount <= 0 {
		return nil, fmt.Errorf("items count cannot less than 1")
	}

	var itemsCountL = itemsCount
	var result = make(map[int]int)
	var lastPackageSize = 0

	fmt.Printf("ItemCount: %d\n", itemsCountL)
	for index, val := range packageSizes {
		if index == len(packageSizes)-1 {
			if itemsCountL > val {
				result[lastPackageSize] = result[lastPackageSize] + 1
			} else {
				result[val] = result[val] + 1
				itemsCountL = itemsCountL % val
			}
		} else {
			result[val] = itemsCountL / val
			itemsCountL = itemsCountL % val
		}
		lastPackageSize = val
		// fmt.Printf("index: %d package_size: %d\n", index, val)
	}

	// trim result (clear keys with empty values)
	for key := range result {
		if result[key] == 0 {
			delete(result, key)
		}
	}
	fmt.Printf("Result: %v\n", result)

	return result, nil
}

func AddPackageSize(packageSize int) (bool, error) {
	if packageSize <= 0 {
		return false, fmt.Errorf("package size cannot less than 1")
	} else {
		packageSizes = append(packageSizes, packageSize)
		return true, nil
	}
}

func RemovePackageSize(packageSize int) (bool, error) {
	if packageSize <= 0 {
		return false, fmt.Errorf("package size cannot less than 1")
	} else {
		fmt.Printf("Before PackageSizes: %v\n", packageSizes)
		for index, value := range packageSizes {
			if value == packageSize {
				packageSizes = append(packageSizes[:index], packageSizes[index+1:]...)
			}
		}
		fmt.Printf("After PackageSizes: %v\n", packageSizes)
		return true, nil
	}
}
