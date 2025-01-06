pkgname=Phoebe
pkgver=1.0.1
pkgrel=1
pkgdesc="A CLI tool for managing events"
arch=('x86_64')
url="https://github.com/Ap0ll02/PhoebeEvents"
license=("MIT")
depends=('go')
source=("$pkgname-$pkgver.tar.gz::https://github.com/Ap0ll02/PhoebeEvents/archive/refs/tags/v$pkgver.tar.gz")
sha256sums=('4c2d90ce86e2eebe257c9d8d3eecfd07f0ee332798177dbc0c9494d0fd638455')
trap 'echo "Phoebe CLI has been installed. Please note you may need to add to path if pb isnt working, or decide how to handle this on your own:";
  echo "Needed Path: /usr/bin";
  echo "For simplicity, one can run: export PATH=/usr/bin for bash/zsh";
  echo "For fish: set -gx PATH /usr/bin";
  echo "And source afterwards. source ~/.zshrc | source ~/.bashrc | source ~/.config/fish/config.fish"' EXIT

build() {
  cd "$srcdir/${pkgname}Events-$pkgver"  # Correct the directory name if necessary
  go build -o pb

  # Post Install
}

package() {
  cd "$srcdir/${pkgname}Events-$pkgver"
  # Create ~/.local/bin if it doesn't exist 
  install -Dm755 pb "$pkgdir/usr/bin/pb"  
}
