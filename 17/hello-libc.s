# https://cs.lmu.edu/~ray/notes/gasexamples/
    .global main

    .text
main:                               # This is called by C library's startup code
    mov     $message, %rdi          # First integer (or pointer) parameter in %rdi
    call    puts                    # puts(message)
    ret                             # Return to C library code
message:
    .asciz "Hello, world"           # asciz puts a 0 byte at the end
