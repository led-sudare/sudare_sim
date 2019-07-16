# simulator

LEDすだれのシュミレータ  
2byte(5-6-5)のすだれの情報を受取り、それを表示する  

# Requirements

- Docker
- git

# Install & Run

1. 適当なディレクトリで、`git clone https://github.com/led-sudare/sudare_sim.git`
2. `./dockerbuild.sh`
3. ブラウザを立ち上げ"http://localhost:2345" にアクセスする

# Configuration

## ZeroMQのSUBするターゲットを変えたい場合

1. config.jsonを開く
2. zmqTargetに適切なIPアドレスとポートを記述する
3. sudare_simのDockerコンテナを再起動する

### 注意点
1つのホストマシンで、[cube_adapterも動かす場合](https://github.com/led-sudare/cube_adapter)
config.jsonのzmqTargetにホストマシンのIPアドレスと、cube_adapterのZmqのポート番号(デフォルト5563)を設定してください。


## goのソースコードを変更の反映したい場合
1. sudare_simのDockerコンテナを再起動する


