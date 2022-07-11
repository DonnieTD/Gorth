package lexer

// OPTYPES
const (
	TOKEN_PUSH         = iota
	TOKEN_PLUS         = iota
	TOKEN_DUMP         = iota
	TOKEN_MINUS        = iota
	TOKEN_EQUALS       = iota
	TOKEN_IF           = iota
	TOKEN_END          = iota
	TOKEN_ELSE         = iota
	TOKEN_GREATER_THAN = iota
	TOKEN_DUP          = iota
	TOKEN_WHILE        = iota
	TOKEN_DO           = iota
	TOKEN_MEM          = iota
	TOKEN_LOAD         = iota
	TOKEN_STORE        = iota
	TOKEN_SYSCALL1     = iota
	TOKEN_SYSCALL3     = iota
	COUNT_TOKENS       = iota
)
