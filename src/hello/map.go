package main
import (
        "fmt"
       )

type PersonInfo struct {
    ID string
        name string
        address string
}

func main() {
    var personDB map[string] PersonInfo
    personDB = make(map[string] PersonInfo)
    //personDB := make(map[string] PersonInfo)
    personDB["12345"] = PersonInfo{"12345", "Json", "room 1"}
    personDB["12346"] = PersonInfo{"12346", "xml", "room 2"}

    person, ok := personDB["12345"]
    if ok {
        fmt.Println("Found person", person.name, "with id", person.ID)
    } else {
        fmt.Println("Did not found person")
    }

    delete(personDB, "12345")
    person, ok = personDB["12345"]
    if ok {
        fmt.Println("Found person", person.name, "with id", person.ID)
    } else {
        fmt.Println("Did not found person")
    }
}
