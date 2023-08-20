setlocal buftype=nofile bufhidden=wipe nobuflisted wrap

nnoremap <buffer> <CR> :call CopyBlockContent()<CR>

function! CopyBlockContent()
    let block_start = search('^```', 'bnW')
    let block_end = search('^```', 'nW')
    if block_start > 0 && block_end > 0 && block_start < block_end
        let block_content = join(getline(block_start + 1, block_end - 1), "\n")
        call setreg('+', block_content, 'c')
        call popup_notification("code block copied", {'line': 'cursor+1', 'col': 'cursor+1', 'time': "500"})
    endif
endfunction
