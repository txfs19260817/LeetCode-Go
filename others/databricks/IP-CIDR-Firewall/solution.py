class IpFirewall:
    def __init__(self, rules: list[list[str]]):
        self.rules = []
        for action, cidr in rules:
            allow = action == "ALLOW"
            network, mask = self._parse_cidr(cidr)
            self.rules.append((allow, network, mask))

    def allow_access(self, ip: str) -> bool:
        ip_val = self._parse_ip(ip)
        for allow, network, mask in self.rules:
            if ip_val & mask == network:
                return allow
        return False

    @staticmethod
    def _parse_ip(ip: str) -> int:
        result = 0
        for part in ip.split("."):
            result = (result << 8) | int(part)
        return result

    @staticmethod
    def _parse_cidr(cidr: str) -> tuple[int, int]:
        if "/" in cidr:
            ip_str, prefix_str = cidr.split("/")
            prefix_len = int(prefix_str)
        else:
            ip_str = cidr
            prefix_len = 32
        if prefix_len == 0:
            mask = 0
        else:
            mask = ((1 << 32) - 1) << (32 - prefix_len)
        ip_val = IpFirewall._parse_ip(ip_str)
        network = ip_val & mask
        return network, mask


if __name__ == "__main__":
    # Main example
    fw = IpFirewall([
        ["ALLOW", "192.168.1.100"],
        ["DENY", "192.168.1.0/24"],
        ["ALLOW", "192.168.0.0/16"],
        ["DENY", "0.0.0.0/0"],
    ])
    assert fw.allow_access("192.168.1.100") is True
    assert fw.allow_access("192.168.1.50") is False
    assert fw.allow_access("192.168.2.10") is True
    assert fw.allow_access("10.0.0.1") is False

    # Deny all
    fw2 = IpFirewall([["DENY", "0.0.0.0/0"]])
    assert fw2.allow_access("1.2.3.4") is False
    assert fw2.allow_access("255.255.255.255") is False

    # Exact /32 match
    fw3 = IpFirewall([["ALLOW", "10.0.0.5/32"], ["DENY", "0.0.0.0/0"]])
    assert fw3.allow_access("10.0.0.5") is True
    assert fw3.allow_access("10.0.0.6") is False

    print("All tests passed!")
