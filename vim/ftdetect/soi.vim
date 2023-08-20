au BufNewFile,BufRead *.soi set filetype=soi
au StdinReadPost * if getline(2) =~ "🔗 https://stackoverflow.com/questions/" | setfiletype soi | endif
