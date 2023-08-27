package vm

import (
	"fmt"
)

type Environment map[string]Value

// value is stack element
type Value interface {
	ValueString() string
}

type Symbol struct {
	Name string
}

type Application struct {
	Left  Value
	Right Value
}

type Function struct {
	Arg  string
	Body Value
}

type Callure struct {
	Env  Environment
	Code []Instruction //left, right, Apply, Return
}

type Closure struct {
	Arg  string
	Env  Environment
	Code []Instruction //code, Return
}

type Dump struct {
	Env  Environment
	Code []Instruction
}

func (s Symbol) ValueString() string {
	return s.Name
}

func (f Function) ValueString() string {
	return "^" + f.Arg + "." + f.Body.ValueString()
}

func (a Application) ValueString() string {
	return "(" + a.Left.ValueString() + " " + a.Right.ValueString() + ")"
}

func (c Callure) ValueString() string {
	return fmt.Sprintf("<callure env=%v>", c.Env)
}

func (c Closure) ValueString() string {
	return fmt.Sprintf("<closure env=%v>", c.Env)
}

func (d Dump) ValueString() string {
	return fmt.Sprintf("<dump env=%v code=%v>", d.Env, d.Code)
}

type VM struct {
	stack []Value
	env   Environment
	code  []Instruction
}

func NewVM(code []Instruction) *VM {
	return &VM{
		stack: []Value{},
		env:   make(Environment),
		code:  code,
	}
}

func (vm *VM) isStackEmpty() bool {
	return len(vm.stack) == 0
}

func (vm *VM) pushStack(item Value) {
	vm.stack = append(vm.stack, item)
}

func (vm *VM) popStack() Value {
	if len(vm.stack) == 0 {
		return nil
	}
	item := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return item
}

func (vm *VM) next() Instruction {
	if len(vm.code) == 0 {
		return nil
	}
	inst := vm.code[0]
	vm.code = vm.code[1:]
	return inst
}

func (vm *VM) referEnv(name string) (v Value, ok bool) {
	v, ok = vm.env[name]
	return v, ok
}

func (vm *VM) setEnv(name string, v Value) {
	vm.env[name] = v
}

func (vm *VM) deleteEnv(name string) {
	delete(vm.env, name)
}

func (vm *VM) copyEnvironment() Environment {
	envcopy := make(Environment, len(vm.env))
	for k, v := range vm.env {
		envcopy[k] = v
	}
	return envcopy
}

func (vm *VM) pushDump() {
	envcopy := vm.copyEnvironment()

	//copy code
	codecopy := make([]Instruction, len(vm.code))
	copy(codecopy, vm.code)

	//push Dump
	vm.pushStack(Dump{Env: envcopy, Code: codecopy})
}

func (vm *VM) pushDumpWithInstructions(is ...Instruction) {
	vm.code = append(is, vm.code...)
	vm.pushDump()
}

func (vm *VM) Run() Value {

LOOP:
	for {

		instruction := vm.next()
		if instruction == nil {
			break //for
		}

		switch i := instruction.(type) {
		case Fetch:
			v, ok := vm.referEnv(i.Name)
			if ok {
				vm.pushStack(v)
			} else {
				vm.pushStack(Symbol{Name: i.Name})
			}
		case Apply:
			right := vm.popStack()
			left := vm.popStack()

			if closure, ok := left.(Closure); ok {
				vm.pushDump()
				vm.env = closure.Env
				vm.code = closure.Code
				vm.env[closure.Arg] = right

			} else if callure, ok := left.(Callure); ok {
				vm.pushStack(right)
				vm.pushDumpWithInstructions(Swap{}, Apply{})
				vm.env = callure.Env
				vm.code = callure.Code

			} else if _, ok := right.(Closure); ok {
				//FIXME internal reduction

			} else if callure, ok := right.(Callure); ok {
				vm.pushStack(left)
				vm.pushDumpWithInstructions(Apply{})
				vm.env = callure.Env
				vm.code = callure.Code

			} else {
				vm.pushStack(Application{Left: left, Right: right})
			}

		case Call:
			envcopy := vm.copyEnvironment()
			vm.pushStack(Callure{Env: envcopy, Code: i.Code})

		case Close:
			envcopy := vm.copyEnvironment()
			vm.pushStack(Closure{Arg: i.Arg, Env: envcopy, Code: i.Code})

		case Return:
			result := vm.popStack()
			d := vm.popStack()
			if callure, ok := result.(Callure); ok {
				vm.pushStack(d)
				vm.pushDumpWithInstructions(Return{})
				vm.env = callure.Env
				vm.code = callure.Code

			} else {
				vm.pushStack(result)
				if dump, ok := d.(Dump); ok {
					vm.env = dump.Env
					vm.code = dump.Code
				} else {
					panic("lost dump in return instruction")
				}
			}
		case Swap:
			first := vm.popStack()
			second := vm.popStack()
			vm.pushStack(first)
			vm.pushStack(second)

		default:
			panic("unknown instruction")
		}
	} //for

	if vm.isStackEmpty() {
		panic("vm.Run completed but stack is empty")
	}

	rest := vm.popStack()

	switch v := rest.(type) {
	case Callure:
		vm.pushDump()
		vm.env = v.Env
		vm.code = v.Code
		goto LOOP

	case Closure:
		//FIXME internal reduction
	}

	return rest
}
