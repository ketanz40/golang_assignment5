package main

import (
	"fmt"
	"errors"
)


//Node struct
type Node struct{
	data int
	Next *Node
	Prev *Node
}

type ListFeature interface{
	addNode()
	deleteNode()
	printNodes()
	length() int
}

//Global Variables (List and a Node to be added in)
var listNodes = make([]Node, 0)
var node1 Node = Node{data: 102, Next:  &node2, Prev: nil}
var node2 Node = Node{data: 52, Next: nil, Prev: &node1}

func addNode(newData int){
	prevNode := listNodes[length()]
	newNode := Node{data: newData, Next: nil, Prev: &prevNode}
	listNodes = append(listNodes, newNode)
}

func deleteNode(location int){
	listNodes = append(listNodes[:location], listNodes[location:]...)
}

func printNodes(location int) error {
	if length() == 0 {
		return errors.New("List is empty")
	} else {
		if location == (length()) {
			lastNode := listNodes[location]
			fmt.Print(lastNode.data)
		}
		nodeToPrint := listNodes[location]
		fmt.Print(nodeToPrint.data, " ")
		newLocation := location + 1	
		printNodes(newLocation)
		}
	return nil
}

func length() int {
	return len(listNodes) + 1
}

func main(){
	listNodes = append(listNodes, node1)
	listNodes = append(listNodes, node2)
	loop:
		for{
			var userChoice int
			fmt.Print("Enter 1 to add to the list\nEnter 2 to delete from the list\nEnter 3 to print the list" +
			"\nEnter 4 to print the length of the list\nEnter 5 to quit the program\n")
			fmt.Scanln(&userChoice)
			switch userChoice {
			case 1:
				var userData int
				fmt.Print("Enter the data: ")
				fmt.Scanln(&userData)
				addNode(userData)
				fmt.Println()
				break
			case 2:
				var userLocation int
				fmt.Print("Where did you want to delete the node? ")
				fmt.Scanln(&userLocation)
				deleteNode(userLocation)
				fmt.Println()
				break
			case 3:
				err := printNodes(0)
				if err != nil{
					fmt.Println(err)
				}
				fmt.Println()
				break
			case 4:
				fmt.Println(length())
				fmt.Println()
				break
			case 5:
				fmt.Println("Goodbye")
				break loop
			default:
				fmt.Println("Invalid Input")
				fmt.Println("Please enter a valid input")
				fmt.Println()
		}
	}
}