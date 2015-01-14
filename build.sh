# 設定
APP_NAME="kcl-go"
APP_OS="linux darwin windows"
APP_ARCH="386 amd64"

# Go1.2.1をダウンロードする
pushd ~/
curl -s -o go.tar.gz https://go.googlecode.com/files/go1.2.1.linux-amd64.tar.gz
tar xzf go.tar.gz
export GOROOT=~/go
export PATH=$GOROOT/bin:$PATH
go version
popd

# goxをインストールする
go get github.com/mitchellh/gox
gox -build-toolchain -os="$APP_OS" -arch="$APP_ARCH"

# gitのコミットからバージョンを採番する
APP_VERSION=$(git log --pretty=format:"%h (%ad)" --date=short -1)
echo APP_VERSION is $APP_VERSION

# 必要なライブラリを集める
go get github.com/mitchellh/cli

# リリース用ディレクトリ作成
mkdir -p artifacts

# クロスコンパイルする
gox ./cmd/kclgo -os="$APP_OS" -arch="$APP_ARCH" -output="artifacts/{{.OS}}-{{.Arch}}/$APP_NAME" -ldflags "-X main.version '$APP_VERSION'"
find artifacts
gox ./cmd/kclsample -os="$APP_OS" -arch="$APP_ARCH" -output="artifacts/{{.OS}}-{{.Arch}}/$APP_NAME" -ldflags "-X main.version '$APP_VERSION'"
find artifacts
