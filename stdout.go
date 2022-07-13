package logger

type Stdout struct {
}

func (out *Stdout) Init(opts *LogOptions) {
	println("--- Using STDOUT ---")
}

func (out *Stdout) Output(msg string) {
	println(msg)
}
