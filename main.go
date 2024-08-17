package main

type Storage interface {
	Insert(e *Employee)
	Get()
	List()
	Update()
	Delete()
}

func main() {

}
