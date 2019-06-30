Led SUDARE simulator on win
===
ここは主に次の目的でDockerするためのものです。  
* とりあえず動作させる。  


# Docker🐳

## Require

* Git 
* Docker

## Get Started🏁

0. VirtualBox settings  
    **Windows Only**  
    Setting port forward 2345  
    ref: https://teratail.com/questions/151484  

1. Get src
    ```sh
    $ git clone https://github.com/led-sudare/simulator.git
    ```

2. Build
    ```sh
    $ docker image build -t sudare-sim:alpine --file=.docker-win/Dockerfile .
    $ docker image ls
    REPOSITORY          TAG                  IMAGE ID            CREATED             SIZE
    sudare-sim          alpine               cd3ce71d3681        3 minutes ago       19.2MB
    <none>              <none>               2e81adda8bf5        4 minutes ago       517MB
    ```

3. Run
    ```sh
    $ docker container run -d --rm --name sim -p 2345:2345 sudare-sim:alpine
    ```

4. Browse
    お手持ちのブラウザで次のURLにアクセスする。  
    http://localhost:2345  
      
    正常に設置・動作できていれば、黒点がぐるぐる回るシミュレーターが表示される。  



