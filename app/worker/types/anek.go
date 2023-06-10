package types

type StringMap map[string]string

type Anek struct {
	Str   string
	Int   int
	Arr   []int
	Dict  StringMap
	Float float64
}

func MakeAnek() *Anek {
	return &Anek{
		Str: `Знаете почему меня называют на работе 007?
0 - желаний работать
0 - мотивации
7 - перекуров за час`,
		Int: 228,
		Arr: []int{1, 3, 3, 7},
		Dict: map[string]string{
			"Rzaka": "9-10",
			"Smysl": "5-6",
			"Ziza":  "10000",
		},
		Float: 3.141592653589793238462643383279}
}
