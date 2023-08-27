# stackovimflow

Search StackOverflow and view answers (minus comments) in Vim. Using Google CustomSearch apis. So I don't have to leave the terminal and keyboard to copy & paste from StackOverflow. 🙊

![demo](https://github.com/JackKCWong/stackovimflow/blob/bd103f6ed08ae68f62d0738884d1a44244886993/demo.gif)

## Installation

```
go install github.com/JackKCWong/stackovimflow@latest
stackovimflow install-vim
```


### Configuration

You need a Google CustomSearch API key to use this plugin.
You can get one from [Google CustomSearch](https://cse.google.com/cse/all).

Set the following environment variables to configure the plugin.

```
export GCS_API_KEY=your-api-key
export GCS_CX=your-cse-id

# optional proxy setup
export GCS_HTTP_PROXY=your-proxy
export GCS_HTTPS_PROXY=your-proxy
```


## Usage

```
# put this in your .bash_profile or equivalent

function sos() {
    stackovimflow search $@ | vim -
}

sos your question as if you are searching on StackOverflow
```

If you want to see some code block syntax highlight, add the following to your `.vimrc`

```vim
let g:markdown_fenced_languages = ['html', 'js=javascript', 'golang=go', 'go', 'rust', 'java', 'bash']
```

It first load the top 5 results in a vim buffer with filetype soi. When you press `<CR>`,
it will open the selected answer in a new buffer with filetype somd.

In a somd, you can press `<CR>` to copy the code block to system clipboard when you are inside a \`\`\`code block\`\`\` or \`code block\`.

Use `<C-j>` to enlarge the somd window.

Use `<C-k>` to shrink the somd window.

Use `<C-n>` to jump to next answer.

Use `d` to navigate down by half-page.

Use `u` to navigate up by half-page.

Use `q` to close the buffer..

