package homework06

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	println("Hello, my name is", p.Name)
}

// 封装

type Children struct {
	name string
	Age  int
}

func (c *Children) GetNames() string {
	return c.name
}

// 继承
type Animal struct {
	name string
}

func (a *Animal) Eat() {
	println(a.name, "is eating")
}

type Dog struct {
	Animal
}

func (d *Dog) Eat() {
	println(d.name, "is eating")
}

func TestInherit() {
	dog := Dog{
		Animal: Animal{
			name: "dog",
		},
	}
	dog.Eat()
}

// 多态
type Programmer interface {
	Program() string
}

type GoProgrammer struct{}

func (g *GoProgrammer) Program() string {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct{}

func (j *JavaProgrammer) Program() string {
	return "System.out.println(\"Hello World\")"
}

func TestPolymorphism() {
	goProg := GoProgrammer{}
	javaProg := JavaProgrammer{}
	programmers := []Programmer{&goProg, &javaProg}
	for _, programmer := range programmers {
		println(programmer.Program())
	}
}

type Engine struct {
}

func (e *Engine) Start() {
	println("Engine is starting")
}

type Car struct {
	Engine
}

func (c *Car) Start() {
	c.Engine.Start()
	println("Car is starting")
}

func TestEmbed() {
	car := Car{}
	car.Start()
}
