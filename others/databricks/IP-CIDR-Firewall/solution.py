class IpFirewall:
    def __init__(self, rules: list[list[str]]):
        self.rules = []
        for r in rules:
            allow = r[0] == "ALLOW"
            network, mask = self._parse_cidr(r[1])
            self.rules.append((allow, network, mask))

    def allow_access(self, ip: str) -> bool:
        parsed_ip = self._parse_ip(ip)
        for allow, network, mask in self.rules:
            if parsed_ip & mask == network:
                return allow
        return False

    def allow_cidr(self, cidr: str) -> bool:
        """Return True if EVERY IP in *cidr* would be allowed."""
        network, _ = self._parse_cidr(cidr)
        prefix_len = int(cidr.split("/")[1]) if "/" in cidr else 32
        return self._allow_range(network, prefix_len)

    def _allow_range(self, network: int, prefix_len: int) -> bool:
        """
        Return True only if every IP in the query prefix is allowed.

        Next is checking an entire CIDR block.
        This is trickier because a query range might be only partially covered by a rule,
        or have a more specific DENY rule inside a broader ALLOW rule.
        To solve this, I'll write a recursive helper function _allow_range.
        As I check the rules, if I find a rule that completely covers our query block,
        and I haven't seen any smaller subset rules inside it, I can just return its action.
        But, if I detect that a rule only covers a subset of our query, I'll set a needs_split flag.
        This means I can't trust a broad match and must split the query into smaller blocks.
        """
        # Query range represented as a network + prefix.
        # q_mask keeps only the fixed prefix bits for this query.
        q_mask = (0xFFFFFFFF << (32 - prefix_len)) & 0xFFFFFFFF
        # True means at least one more specific rule exists inside this query,
        # so we cannot trust a broad match and must split into subranges.
        needs_split = False
        for allow, r_net, r_mask in self.rules:
            # Rule covers entire query?
            if q_mask & r_mask == r_mask and network & r_mask == r_net:
                # If no subset rules were seen, this first full-covering rule
                # determines the whole query range.
                if needs_split:
                    break
                return allow
            # Rule is a proper subset of query?
            if q_mask & r_mask != r_mask and r_net & q_mask == network:
                needs_split = True
        if prefix_len >= 32:
            # Leaf (/32): evaluate this exact IP with first-match semantics.
            return self.allow_access(str(network))
        # Split query into two child prefixes by toggling the next host bit.
        bit = 1 << (31 - prefix_len)
        # Entire query is allowed only if BOTH halves are fully allowed.
        return (self._allow_range(network, prefix_len + 1) and
                self._allow_range(network | bit, prefix_len + 1))

    @staticmethod
    def _parse_prefix_len(cidr: str) -> int:
        return int(cidr.split("/")[1]) if "/" in cidr else 32

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
        mask = (0xFFFFFFFF << (32-prefix_len)) & 0xFFFFFFFF
        network = IpFirewall._parse_ip(ip_str) & mask
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

    # --- AllowCIDR follow-up ---

    # Entire range covered by ALLOW
    fw4 = IpFirewall([["ALLOW", "192.168.0.0/16"], ["DENY", "0.0.0.0/0"]])
    assert fw4.allow_cidr("192.168.1.0/24") is True

    # Entire range covered by DENY
    fw5 = IpFirewall([["DENY", "192.168.1.0/24"], ["ALLOW", "0.0.0.0/0"]])
    assert fw5.allow_cidr("192.168.1.0/24") is False

    # Two halves both ALLOW
    fw6 = IpFirewall([
        ["ALLOW", "192.168.1.0/25"],
        ["ALLOW", "192.168.1.128/25"],
        ["DENY", "0.0.0.0/0"],
    ])
    assert fw6.allow_cidr("192.168.1.0/24") is True

    # One half missing → DENY
    fw7 = IpFirewall([["ALLOW", "192.168.1.0/25"], ["DENY", "0.0.0.0/0"]])
    assert fw7.allow_cidr("192.168.1.0/24") is False

    # Main example
    assert fw.allow_cidr("192.168.1.0/24") is False   # .50 is denied
    assert fw.allow_cidr("192.168.2.0/24") is True    # fully covered by /16
    assert fw.allow_cidr("10.0.0.0/8") is False       # deny all

    # Single IP as /32
    fw8 = IpFirewall([["ALLOW", "192.168.1.100"], ["DENY", "0.0.0.0/0"]])
    assert fw8.allow_cidr("192.168.1.100/32") is True
    assert fw8.allow_cidr("192.168.1.50/32") is False

    print("All tests passed!")
