# Maintainer: AlphaNecron <necronqwq@outlook.com>
pkgname=nextfetch-git
pkgver=0.0.0
_pkgname=nextfetch
pkgrel=1
pkgdesc="Simple cross-platform fetch program, written in Go"
arch=('any')
url="https://github.com/AlphaNecron/$_pkgname"
license=('MIT')
optdepends=(
	'lsb_release: Distro detection'
)
makedepends=('go>=1.17', 'git')
source=("git+https://github.com/AlphaNecron/$_pkgname.git")
sha256sums=('SKIP')

pkgver() {
  cd $_pkgname
  printf "r0.0.%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

build() {
  cd "$srcdir/$_pkgname"
  make DESTDIR="$pkgdir" PREFIX="/usr" build
}

package() {
  cd $srcdir/$_pkgname
  make DESTDIR="$pkgdir" PREFIX="/usr" install
}

post_install() {
	cd $srcdir/$_pkgname
	make install-config
}
