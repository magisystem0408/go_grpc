
## 用語確認
レイテンシー：送信して応答が返ってくるまでの時間


## gRPCとwebsocketについて

gRPCについて
- http/2はバイナリベース
- multiplexy
- serverからクライアントに対してpushすることができる
- headerの圧縮が必須


WSについて
  - httpからのハンドシェイクを持って切り替える
  - 一本のconnectionだけで通信を行う
  - 双方向に通信が可能
  - 完全に非同期