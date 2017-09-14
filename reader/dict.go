package reader

var dict = map[string]string{
	"add": `//add
@SP
M=M-1
A=M
D=M
@SP
A=M-1
M=D+M
`,
	"eq": `//eq
@SP
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@SP
A=M
M=0
@TRUE
D,JEQ
@SP
A=M
M=1
(TRUE)
`,
"lt":`



`,
	"push_constant": `//push constant {num}
@{num}
D=A
@SP
A=M
M=D
@SP
M=M+1
`,
}
