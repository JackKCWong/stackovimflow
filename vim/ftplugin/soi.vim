setlocal readonly
setlocal buftype=nofile bufhidden=wipe nobuflisted nowrap

function! FetchSO()
    let line = line('.')
    let col = col('.')
    let search_url = search('\chttps\=', 'bW')
    if search_url > 0
        let url = matchstr(getline(search_url), '\chttps:\/\/\S*')
        call RunCommandToBuffer('stackovimflow fetch ' . shellescape(url))
    endif
    call cursor(line, col)
endfunction


function! RunCommandToBuffer(command)
    let output = system(a:command)
    let buffer_name = 'StackOverflow'
    let buffer_number = bufnr(buffer_name)

    if buffer_number > 0
        silent execute 'bwipeout! ' . buffer_number
    endif

    silent execute 'botright vnew ' . buffer_name
    setlocal buftype=nofile bufhidden=wipe nobuflisted nowrap
    call setline(1, split(output, '\n'))
    normal! gg
    setlocal filetype=somd
endfunction

nnoremap <buffer> <CR> :call FetchSO()<CR>
