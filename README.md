## dlv remote debug with VSCode

1. Install dlv  
在遠端環境中安裝`dlv`  
```
$ git clone https://github.com/go-delve/delve
$ cd delve
$ go install github.com/go-delve/delve/cmd/dlv
```

dlv會安裝到 `GOPATH/bin` 目錄下。  
如果提示找不到dlv command，就需要將`GOPATH/bin`加入`$PATH`，安裝golang時就可以先設定。  
`$ sudo vi .bashrc`  
在最底部加入以下代碼：  
```
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$HOME/go/bin
```
然後儲存檔案。  
接著呼叫：  
`$ source .bashrc`

2. VSCode 安裝Remote-SSH、Remote Development套件

3. VSCode SSH連線遠端host，並開啟一個遠端工作目錄作為`${workspaceFolder}`

4. 開始Debug之前，本地需要有一份Source code copy

5. 開啟VSCode左側Debug Window，新增一份配置文件launch.json，範例如下：
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Connect to server",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            //"remotePath": "/home/jordan/go/src/test",		// useless
            "port": 2345,
            "host": "192.168.1.126",
            "showLog": true,
            "trace": "log",
            "substitutePath": [
                {
                    "from": "N:\\Development\\go\\src\\test",   // absolute local path
                    "to": "/home/jordan/go/src/test"            // absolute remote path
                }
            ],
            "debugAdapter": "dlv-dap"

        }
    ]
}
```  
***注意事項***  
遠端host、port必須填妥。  
  
`remotePath` 官方文件說明已棄用。改用 `substitutePath`。  
  
本範例是在Windows環境開啟VSCode debug遠端Linux program，  
所以 `from` 填寫**Windows本地source目錄**，`to` 填寫**Linux遠端source目錄**。  
  
在VSCode下斷點可能會遭遇錯誤：  
`Unhandled error in debug adapter: TypeError: Cannot read properties of undefined (reading 'addr')`  
  
加入 `"debugAdapter": "dlv-dap"` ，在**remote/attach**模式下使用新版adapter可解決問題。  
  
6. Run Dbg Server  
在遠端環境中執行以下命令：  
`dlv --listen=":2345" --headless=true --log --api-version=2 exec $HOME/go/bin/test`  
如果需要添加program參數可以像這樣：  
`dlv --listen=":2345" --headless=true --log --api-version=2 exec $HOME/go/bin/test -- --p`  
  
7. VSCode 按下 F5 開心Debug  
