setlocal readonly
setlocal buftype=nofile bufhidden=wipe nobuflisted nowrap

function! FetchAnswers()
    let line = line('.')
    let col = col('.')
    let txt = getline('.')
    let search_url = search('🔗 https://stackoverflow.com/questions', 'bW')
    if search_url > 0
        let url = matchstr(getline(search_url), '\chttps:\/\/\S*')
        call LocateAnswer('stackovimflow fetch ' . shellescape(url), txt)
    endif
endfunction


function! LocateAnswer(command, ans)
    let output = system(a:command)
    let buffer_name = 'StackOverflow'
    let buffer_number = bufnr(buffer_name)

    if buffer_number > 0
        silent execute 'bwipeout! ' . buffer_number
    endif

    silent execute 'botright vnew ' . buffer_name
    setlocal filetype=somd
    call setline(1, split(output, '\n'))
    normal! gg

    if a:ans =~ "  👼  "
        let pat = a:ans[strlen("  👼  "):20]
        call search('\V' . pat, 'w')
    endif
endfunction

nnoremap <buffer> <CR> :call FetchAnswers()<CR>
