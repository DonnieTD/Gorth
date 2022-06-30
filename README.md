# Gorth
Porth in Go xD

<!-- NEXT WE NEED TO CONVERT TO RUNES BRO -->

<!-- BIG RUNE REFACTOR -->
<!-- WE WILL DO THIS BY CREATING THE STRUCTS BELOW and then rewriting the program using it -->
<!-- 
    NEW PROJECT NAME 
    NAH: NAMES ARE HARD ( PORTH IMPLEMENTATION IN GO ) 
-->
<!--  -->
<!-- 
    NAH:
        -- LEXER lexer 
    ::compile()
    ::interpret()
 -->

<!-- IN: FILE PATH OUT: ARRAY OF TOKENS -->
<!-- 
    LEXER:
    -- PROGRAM text
    -- CURSOR int

    ::new() create a new lexer
    ::LoadFILE(filePath) get lines of text from file and set it to PROGRAM (also exit(1)) when anything failse) 
    ::getNextChar (look ahead)
    ::getNextTokenText() (increment till end of token and stack it up and return entire token as text)
    ::TextToToken() change a text token into an actual typed token
    ::ParseText() uses everything above to return a slice of tokens from a file
-->