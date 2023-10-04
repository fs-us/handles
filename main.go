package handles

import (
	"regexp"
	"strconv"
)

func createHandle(handleMap map[string]bool, rawFirstName string, rawLastName string) string {
	handle := ""
	regex := regexp.MustCompile(`\W`)
	firstName := regex.ReplaceAllString(rawFirstName, "")
	lastName := regex.ReplaceAllString(rawLastName, "")

	for {
		length := len(handle)

		if length > len(firstName)+len(lastName) {
			num, _ := strconv.Atoi(handle[len(handle)-1:])
			handle = handle[:len(handle)-1] + strconv.Itoa(num+1)
		} else if length == len(firstName)+len(lastName) {
			handle = firstName + lastName + "1"
		} else if length/2 >= len(lastName) {
			handle = firstName[:length-len(lastName)+1] + lastName
		} else if length/2 >= len(firstName) {
			handle = firstName + lastName[:length-len(firstName)+1]
		} else {
			if length == 0 {
				if len(firstName)+len(lastName) < 4 {
					handle = firstName + lastName
				} else if len(firstName) < 2 {
					handle = firstName + lastName[:3]
				} else if len(lastName) < 2 {
					handle = firstName[:3] + lastName
				} else {
					handle = firstName[:2] + lastName[:2]
				}

			} else if len(handle)%2 == 0 {
				handle = firstName[:length/2] + lastName[:(length/2)+1]
			} else /*if len(currentHandle)%2 == 1*/ {
				handle = firstName[:(length+1)/2] + lastName[:(length+1)/2]
			}
		}

		if !handleMap[handle] {
			handleMap[handle] = true
			return handle
		}

	}

}
