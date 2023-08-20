# stackovimflow

Search StackOverflow and view answers in Vim. Using Google CustomSearch apis. So I don't have to leave the terminal and keyboard to copy & paste from StackOverflow. 🙊

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
function sos() {
    stackovimflow search $@ | vim -
}
sos your question as if you are searching on StackOverflow
```

It first load the top 5 results in a vim buffer. When you press <CR>,
it will open the selected answer in a new buffer.
In the new buffer, if there is a code block, you can press <CR> to copy the code block to system clipboard.
