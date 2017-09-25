Name:   jarvis-agent
Version:    0.1
Release:    1%{?dist}
Summary:    jarvis_agentd

Group:      jarvis
License:    GPL
URL:        www.k2data.com.cn
Source0:    jarvis-agent-default
Source1:    jarvis-agent-init

BuildRequires: bash
Requires: bash 

%description
Jarvis agentd 0.1

%pre
[ -f /etc/default/jarvis-agent   ]||rm -rf /etc/defaul/jarvis-agent
[ -f /etc/init.d/jarvis-agent   ]||rm -rf /etc/init.d/jarvis-agent


%post
systemctl start jarvis-agent

%preun
systemctl stop jarvis-agent

%postun
rm  -rf /usr/lib/systemd/system/jarvis-agent.service
rm -rf /etc/default/jarvis-agent

%setup -q


%build
make build
%install
test -L %{buildroot}/etc/default/%{name} && rm -f %{buildroot}/etc/default/%{name}
test -L %{buildroot}/usr/lib/systemd/system/%{name}.service && rm -f %{buildroot}/usr/lib/systemd/system/%{name}.service
install -D %{buildroot}/../../BUILD/build/jarvis_agent_linux_amd64.bin  %{buildroot}/usr/sbin/jarvis-agent
install -D -m 664 %{SOURCE0} %{buildroot}/etc/default/jarvis-agent
install -D %{SOURCE1} %{buildroot}/usr/lib/systemd/system/jarvis-agent.service

%files
%defattr (-,root,root,0644)
/usr/lib/systemd/system/jarvis-agent.service
/etc/default/jarvis-agent
%attr(0755,root,root) /usr/sbin/jarvis-agent
%changelog
%clean
rm -rf %{buildroot}
