setlocal buftype=nofile bufhidden=wipe nobuflisted wrap

function! CopyBlockContent()
    let line = getline('.')
    let start_col = match(line, '`')
    let end_col = match(line, '`', start_col+1)
    if start_col > 0 && end_col > 0 && start_col < end_col
        let enclosed_content = line[start_col+1:end_col-1]
        call setreg('+', enclosed_content)
        call popup_notification("code block copied", {'line': 'cursor+1', 'col': 'cursor+1', 'time': "500"})
    else
        let block_start = search('^```', 'bnW')
        let block_end = search('^```$', 'nW')
        if block_start > 0 && block_end > 0 && block_start < block_end
            let lines = getline(block_start + 1, block_end - 1)
            let block_content = join(lines, "\n")
            call setreg('+', block_content, 'c')
            call popup_notification(len(lines) . " line(s) copied", {'line': 'cursor+1', 'col': 'cursor+1', 'time': "500"})
        else
            echo 'No code block content found.'
        endif
    endif
endfunction

nnoremap <buffer> <CR> :call CopyBlockContent()<CR>

augroup somd 
  autocmd!
  autocmd FileType somd set clipboard=unnamed
  autocmd FileType somd nnoremap <buffer> <C-h> 3<C-w>>
  autocmd FileType somd nnoremap <buffer> <C-l> 3<C-w><
augroup END