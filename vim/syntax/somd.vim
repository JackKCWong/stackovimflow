setlocal filetype=markdown

syntax region somdCodeBlock start=/^```$/ end=/^```$/
highlight somdCodeBlock ctermfg=blue guifg=blue

highlight separator ctermbg=yellow guibg=yellow
match separator /\*\*\*/

" Automatically enable syntax highlighting for the somd filetype
if exists("b:current_syntax")
    finish
endif
let b:current_syntax = "somd"