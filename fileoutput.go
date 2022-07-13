package logger

import "os"

type FileOutput struct {
	logFile os.File
}

func (out *FileOutput) Init(opts *LogOptions) {
	file, err := os.OpenFile(opts.outputPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		println(err)
		os.Exit(-1)
	}
	out.logFile = *file
}

func (out *FileOutput) Output(msg string) {
	out.logFile.WriteString(msg)
}
