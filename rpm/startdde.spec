%define specrelease 1%{?dist}
%if 0%{?openeuler}
%define specrelease 1
%global _missing_build_ids_terminate_build 0
%global debug_package   %{nil}
%global dde_prefix dde
%else
%global dde_prefix deepin
%endif

Name:           startdde
Version:        5.6.0.35.1
Release:        %{specrelease}
Summary:        Starter of deepin desktop environment
License:        GPLv3
URL:            https://github.com/linuxdeepin/startdde
%if 0%{?openeuler}
Source0:        %{name}_%{version}.orig.tar.xz

BuildRequires:  gocode
Requires:       gocode
Requires:       libXfixes
Requires:       libXcursor
%else
Source0:        %{url}/archive/%{version}/%{name}-%{version}.tar.gz
ExclusiveArch:  %{?go_arches:%{go_arches}}%{!?go_arches:%{ix86} x86_64 %{arm}}

BuildRequires:  compiler(go-compiler)
BuildRequires:  golang(pkg.deepin.io/dde/api/dxinput)
BuildRequires:  golang(pkg.deepin.io/lib)
BuildRequires:  golang(github.com/godbus/dbus)
BuildRequires:  golang(github.com/linuxdeepin/go-x11-client)
BuildRequires:  golang(github.com/davecgh/go-spew/spew)
BuildRequires:  golang(golang.org/x/xerrors)
BuildRequires:  systemd-rpm-macros
BuildRequires:  make
%endif

BuildRequires:  golang
BuildRequires:  jq
BuildRequires:  glib2-devel
BuildRequires:  pkgconfig(x11)
BuildRequires:  libXcursor-devel
BuildRequires:  libXfixes-devel
BuildRequires:  gtk3-devel
BuildRequires:  pulseaudio-libs-devel
BuildRequires:  libgnome-keyring-devel
BuildRequires:  alsa-lib-devel
BuildRequires:  pkgconfig(gudev-1.0)

Provides:       x-session-manager
Requires:       %{dde_prefix}-daemon
Requires:       procps
Requires:       deepin-desktop-schemas
Requires:       %{dde_prefix}-kwin
Recommends:     %{dde_prefix}-qt5integration

%description
%{summary}.

%prep
%autosetup -p1 -n %{name}-%{version}
## Scripts in /etc/X11/Xsession.d are not executed after xorg start
sed -i 's|X11/Xsession.d|X11/xinit/xinitrc.d|g' Makefile
# fix deepin-daemon executables path
find * -type f -not -path "rpm/*" -print0 | xargs -0 sed -i 's:/lib/deepin-daemon/:/libexec/deepin-daemon/:'
# fix dde-polkit-agent path
sed -i '/polkit/s|lib|libexec|' watchdog/dde_polkit_agent.go

%build
%if 0%{?openeuler}
export GOPATH=/usr/share/gocode
%make_build GO_BUILD_FLAGS=-trimpath
%else
export GOPATH="%{gopath}"
%make_build GO_BUILD_FLAGS="%{gobuildflags}"
# rebuild other executables with different build-ids
for cmd in fix-xauthority-perm greeter-display-daemon; do
    rm $cmd
    %make_build $cmd GO_BUILD_FLAGS="%{gobuildflags}"
done
%endif

%install
%make_install

%post
xsOptsFile=/etc/X11/Xsession.options
update-alternatives --install /usr/bin/x-session-manager x-session-manager \
    /usr/bin/startdde 90 || true
if [ -f $xsOptsFile ];then
	sed -i '/^use-ssh-agent/d' $xsOptsFile
	if ! grep '^no-use-ssh-agent' $xsOptsFile >/dev/null; then
		echo no-use-ssh-agent >> $xsOptsFile
	fi
fi

%files
%doc README.md
%license LICENSE
%{_sysconfdir}/X11/xinit/xinitrc.d/00deepin-dde-env
%{_sysconfdir}/X11/xinit/xinitrc.d/01deepin-profile
%{_sysconfdir}/profile.d/deepin-xdg-dir.sh
%{_bindir}/%{name}
%{_sbindir}/deepin-fix-xauthority-perm
%{_datadir}/xsessions/deepin.desktop
%{_datadir}/lightdm/lightdm.conf.d/60-deepin.conf
%{_datadir}/%{name}/
%{_libexecdir}/deepin-daemon/greeter-display-daemon

%changelog
* Wed Oct 14 2020 guoqinglan <guoqinglan@uniontech.com> - 5.6.0.5-2
- bugfix-49318, modify /usr/lib/deepin-daemon path

* Sat Oct 10 2020 guoqinglan <guoqinglan@uniontech.com> - 5.6.0.5-1
- bugfix-49970, fix post add preun scripts
