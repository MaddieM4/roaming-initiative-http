%if   ip["server"] == ip["clearnet4"]:
    <header class="clearnet4">
        You are accessing this site via the IPv4 Clearnet.
    </header>
%elif ip["server"] == ip["clearnet6"]:
    <header class="clearnet6">
        You are accessing this site via the IPv6 Clearnet.
    </header>
%elif ip["server"] == ip["hyperboria"]:
    <header class="hyperboria">
        You are accessing this site via the Hyperboria mesh network. Awesome!
    </header>
%else:
    <header class="dev">
        You are not accessing this site from a known IP, so the server is probably a development machine ({{ ip['server'] }})
    </header>
%end
