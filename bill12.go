package main

type Bill struct {
	Name string
	Items map[string]interface{} // this is a complete interface 
	Tip float64
}

func bill12(name string) Bill {

	billInstance := Bill{
		Name: name,
		Items: map[string]interface{}{"cheetos": 10}, // this is instantiating an empty interface, hence the 2 curly braces. OR this can work "make(map[string]interface{})"
		Tip: .9,
	}

	return billInstance
}
