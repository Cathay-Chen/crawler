## GO Package

### golang.org/x/text
`golang.org/x/text` 子模块提供了多个子包，每个子包中都有不同的方法和类型。下面列举一些常用的方法，供参考：

- `encoding`: 提供字符编码和解码相关的库和工具。
  - `unicode`: 包含了许多关于 Unicode 字符集的函数和类型，如 `IsDigit`、`IsLetter` 等等。
  - `charmap`: 提供了多种字符映射表的实现，如 `ISO8859_1`、`Windows1252` 等等。
  - `json`: 提供了对 JSON 数据进行编码和解码的接口和实现。
  - `xml`: 提供了对 XML 数据进行编码和解码的接口和实现。
  
- `message`: 提供国际化和本地化（i18n/l10n）相关的库和工具。
  - `catalog`: 提供了消息目录和翻译器的实现，可以用于支持多语言消息翻译。
  - `language`: 提供了表示语言标签（如 `en-US`、`zh-CN` 等）的类型和方法，方便进行语言相关操作。
  - `plural`: 提供了多种复数形式选择规则的实现，可以根据数量自动匹配正确的复数形式。
  
- `transform`: 提供 Unicode 转换和规范化相关的库和工具。
  - `transform`: 提供了转换器的接口和实现，可以用于对文本进行转换、规范化等操作。
  - `norm`: 提供了多种 Unicode 规范化的实现，如 NFD、NFC、NFKD、NFKC 等等。

### golang.org/x/net
`golang.org/x/net` 包含了多个子包，提供了网络编程和通信相关的库和工具。下面列举一些常用的用法：

- `http`: 提供了 HTTP 客户端和服务器的实现，可以用于进行 Web 开发和网络通信。
  - `http.Client`：HTTP 客户端类型，封装了 HTTP 请求和响应的发送和接收。
  - `http.Server`：HTTP 服务器类型，能够处理 HTTP 请求并向客户端发送 HTTP 响应。
  - `http.Get(url string)`：快速获取指定 URL 的内容，并返回响应结果。
    
- `jsonrpc`: 提供了 JSON-RPC 协议的实现，可以用于在不同计算机之间进行远程过程调用（RPC）。
  - `jsonrpc.NewClient(conn io.ReadWriteCloser) *Client`：创建一个 JSON-RPC 客户端，并绑定到指定的读写流上。
  - `jsonrpc.NewServerCodec(codec ServerCodec) Server`：创建一个 JSON-RPC 服务器，并绑定到指定的编解码器上。
  
- `websocket`: 提供了 WebSocket 协议的实现，可以用于实现双向通信的 WebSocket 应用程序。
  - `websocket.Dial(url string, header http.Header) (*Conn, error)`：创建一个 WebSocket 连接，并返回连接对象。
  - `websocket.Message.Send(ws *Conn, msg []byte) error`：向指定的 WebSocket 连接发送消息。
  - `websocket.Message.Receive(ws *Conn, msg *[]byte) error`：从指定的 WebSocket 连接接收消息。