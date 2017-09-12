package parser


var dict = map[string]string{
"add":`//add
@SP
M=M-1
A=M
D=M
@SP
A=M-1
M=D+M
`,
"push_constant":`//push constant {num}
@{num}
D=A
@SP
A=M
M=D
@SP
M=M+1
`,
}
