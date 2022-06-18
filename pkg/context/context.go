package context

var beans = make(map[string]interface{})

func Add(name string, bean interface{}) {
	beans[name] = bean
}

func Get(name string) interface{} {
	i, ok := beans[name]

	if !ok {
		panic("not found bean")
	}

	return i
}
