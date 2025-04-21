package main

import (
	"fmt"
)

func main() {
	// var list = []string{"apple", "banana", "orange", "peach", "grapes","mango"}
	// fmt.Println(list)

	// list = append(list[1:3])
	// fmt.Println(list)


		var courses = []string{"java", "python", "ruby", "c", "go"}
		fmt.Println(courses)

		var index int=1

		courses = append(courses[:index],courses[index+1:]... )

		fmt.Println(courses)

		//maps in go lang

		 languages := make(map[string]string)
		 languages["js"]="Javascript"
		 languages["rb"]="ruby"
		 languages["jv"]="java"
		 languages["py"]="python"

		//  fmt.Println(languages)
		//  delete(languages,"rb")
		//  fmt.Println(languages)


for key ,value :=range languages{
	fmt.Printf("For key %v,value is %v\n", key,value)
}

}
