# simulator

LEDすだれのシュミレータ  
2byte(5-6-5)のすだれの情報を受取り、それを表示する  

# Requirements

- Docker
- git

# Install & Run

1. 適当なディレクトリで、`git clone https://github.com/led-sudare/simulator.git`
2. `./dockerbuild.sh`
3. ブラウザを立ち上げ"http://localhost:2345"にアクセスする

# Configuration

## ZeroMQのSUBするターゲットを変えたい場合

1. config.jsonを開く
2. zmqTargetに適切なIPアドレスとポートを記述する
3. simulatorのDockerコンテナを再起動する

### 注意点
1つのホストマシンで、[3d_led_cube_adapterも動かす場合](https://github.com/led-sudare/3d_led_cube_adapter)
config.jsonのzmqTargetにホストマシンのIPアドレスと、3d_led_cube_adapterのZmqのポート番号(デフォルト5563)を設定してください。


## goのソースコードを変更の反映したい場合
1. simulatorのDockerコンテナを再起動する


