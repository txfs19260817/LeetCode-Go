import heapq
import random
from typing import Dict, List, Tuple


class Solution:
    # 兼容入口：无权重时走 Follow-up 1（唯一），有权重时走 Follow-up 2（加权唯一）。
    def generateNFT(self, config: dict, n: int) -> List[Dict[str, str]]:
        traits = config.get("traits", {})
        weighted = False
        for values in traits.values():
            if values and isinstance(values[0], dict):
                weighted = True
                break
        if weighted:
            return self.generateNFTWeightedUnique(config, n)
        return self.generateNFTUnique(config, n)

    # 主问题：允许重复，等概率随机抽样。
    def generateNFTBasic(self, config: dict, n: int) -> List[Dict[str, str]]:
        traits = config.get("traits", {})
        trait_names = sorted(traits.keys())

        result: List[Dict[str, str]] = []
        for _ in range(n):
            nft: Dict[str, str] = {}
            for trait in trait_names:
                values = traits[trait]
                if isinstance(values[0], dict):
                    candidates = [v["name"] for v in values]
                else:
                    candidates = list(values)
                nft[trait] = random.choice(candidates)
            result.append(nft)
        return result

    # Follow-up 1：唯一组合（无权重）。
    def generateNFTUnique(self, config: dict, n: int) -> List[Dict[str, str]]:
        traits = config.get("traits", {})
        pairs: List[Tuple[str, List[str]]] = []
        total_combos = 1
        for trait in sorted(traits.keys()):
            values = traits[trait]
            if isinstance(values[0], dict):
                opts = [v["name"] for v in values]
            else:
                opts = list(values)
            pairs.append((trait, opts))
            total_combos *= len(opts)

        if n > total_combos:
            raise ValueError("n exceeds number of unique combinations")

        combos: List[Dict[str, str]] = []
        current: Dict[str, str] = {}

        def build(idx: int) -> None:
            # 递归终止：所有 trait 都已选完，记录一份完整组合。
            if idx == len(pairs):
                combos.append(current.copy())
                return

            # 当前层处理第 idx 个 trait。
            trait, opts = pairs[idx]
            for name in opts:
                # 选择：给当前 trait 选一个候选值。
                current[trait] = name
                # 递归进入下一层 trait。
                build(idx + 1)
            # 回溯：撤销当前层选择，返回上一层继续尝试其它分支。
            current.pop(trait, None)

        build(0)
        random.shuffle(combos)
        return combos[:n]

    # Follow-up 2：加权 + 唯一组合。
    def generateNFTWeightedUnique(self, config: dict, n: int) -> List[Dict[str, str]]:
        traits = config.get("traits", {})
        pairs: List[Tuple[str, List[Tuple[str, int]]]] = []
        total_combos = 1

        for trait in sorted(traits.keys()):
            values = traits[trait]
            if isinstance(values[0], dict):
                opts = [(v["name"], int(v.get("weight", 1))) for v in values]
            else:
                opts = [(value, 1) for value in values]
            pairs.append((trait, opts))
            total_combos *= len(opts)

        if n > total_combos:
            raise ValueError("n exceeds number of unique combinations")

        combos: List[Tuple[Dict[str, str], float]] = []
        current: Dict[str, str] = {}

        def build(idx: int, weight: float) -> None:
            if idx == len(pairs):
                combos.append((current.copy(), weight))
                return
            trait, opts = pairs[idx]
            for name, value_weight in opts:
                current[trait] = name
                option_weight = value_weight if value_weight > 0 else 1
                build(idx + 1, weight * float(option_weight))
            current.pop(trait, None)

        build(0, 1.0)
        if n == len(combos):
            random.shuffle(combos)
            return [combo for combo, _ in combos]

        heap: List[Tuple[float, Dict[str, str]]] = []
        for combo, weight in combos:
            weight = weight if weight > 0 else 1.0
            key = random.random() ** (1.0 / weight)
            if len(heap) < n:
                heapq.heappush(heap, (key, combo))
            elif key > heap[0][0]:
                heapq.heapreplace(heap, (key, combo))
        return [combo for _, combo in heap]


def _trait_allowed_values(traits: dict) -> Dict[str, set]:
    allowed: Dict[str, set] = {}
    for trait, values in traits.items():
        if not values:
            allowed[trait] = set()
            continue
        if isinstance(values[0], dict):
            allowed[trait] = {value["name"] for value in values}
        else:
            allowed[trait] = set(values)
    return allowed


def _assert_valid(result: List[Dict[str, str]], traits: dict) -> None:
    allowed = _trait_allowed_values(traits)
    for nft in result:
        assert len(nft) == len(traits)
        for trait, values in allowed.items():
            assert trait in nft
            assert nft[trait] in values


def _assert_unique(result: List[Dict[str, str]]) -> None:
    seen = set()
    for nft in result:
        key = tuple(sorted(nft.items()))
        assert key not in seen
        seen.add(key)


def test1() -> None:
    print(" ========= Test 1  =========")
    solution = Solution()
    config = {
        "name": "config-1",
        "size": "large",
        "traits": {
            "nose": ["pointy", "tiny", "flat"],
            "mouth": ["small", "wide", "thin"],
            "eyes": ["blue", "green", "brown"],
        },
    }
    result = solution.generateNFT(config, 5)
    _assert_valid(result, config["traits"])
    _assert_unique(result)
    printResult(result)


def test_basic() -> None:
    print(" ========= Test Basic  =========")
    solution = Solution()
    config_plain = {
        "name": "basic-plain",
        "size": "small",
        "traits": {
            "nose": ["pointy", "tiny"],
            "eyes": ["blue", "green"],
        },
    }
    result_plain = solution.generateNFTBasic(config_plain, 6)
    assert len(result_plain) == 6
    _assert_valid(result_plain, config_plain["traits"])

    # 显式覆盖 generateNFTBasic 中 values[0] 是 dict 的分支。
    config_weighted_shape = {
        "name": "basic-weighted-input",
        "size": "small",
        "traits": {
            "mouth": [
                {"name": "small", "weight": 3},
                {"name": "wide", "weight": 1},
            ],
            "eyes": [
                {"name": "blue", "weight": 1},
                {"name": "brown", "weight": 2},
            ],
        },
    }
    result_weighted = solution.generateNFTBasic(config_weighted_shape, 6)
    assert len(result_weighted) == 6
    _assert_valid(result_weighted, config_weighted_shape["traits"])
    printResult(result_weighted)


def test2() -> None:
    print(" ========= Test 2  =========")
    solution = Solution()
    config = {
        "name": "config-2",
        "size": "small",
        "traits": {
            "color": ["red", "blue", "green"],
            "shape": ["circle", "square"],
        },
    }
    result = solution.generateNFT(config, 3)
    _assert_valid(result, config["traits"])
    _assert_unique(result)
    printResult(result)


def test3() -> None:
    print(" ========= Test 3  =========")
    solution = Solution()
    config = {
        "name": "config-3",
        "size": "large",
        "traits": {
            "color": ["red", "blue", "green", "yellow", "purple"],
            "texture": ["smooth", "rough", "grainy"],
            "size": ["tiny", "small", "medium", "large"],
        },
    }
    result = solution.generateNFT(config, 3)
    _assert_valid(result, config["traits"])
    _assert_unique(result)
    printResult(result)


def test4_weighted() -> None:
    print(" ========= Test 4  =========")
    solution = Solution()
    config = {
        "name": "config-weighted",
        "size": "large",
        "traits": {
            "nose": [
                {"name": "pointy", "weight": 1},
                {"name": "tiny", "weight": 2},
                {"name": "flat", "weight": 3},
            ],
            "mouth": [
                {"name": "small", "weight": 1000},
                {"name": "wide", "weight": 1},
                {"name": "thin", "weight": 1},
            ],
            "eyes": [
                {"name": "blue", "weight": 10},
                {"name": "green", "weight": 2},
                {"name": "brown", "weight": 1},
            ],
        },
    }
    result = solution.generateNFT(config, 5)
    _assert_valid(result, config["traits"])
    _assert_unique(result)
    printResult(result)


def test5_too_many() -> None:
    print(" ========= Test 5  =========")
    solution = Solution()
    config = {
        "name": "simple",
        "size": "small",
        "traits": {
            "color": ["red", "blue", "green"],
            "shape": ["circle", "square"],
        },
    }
    try:
        solution.generateNFT(config, 10)
    except ValueError:
        print("raised ValueError for too many combinations")


def printResult(result: List[Dict[str, str]]) -> None:
    for each in result:
        print(each)


def main() -> None:
    random.seed(1)
    test_basic()
    test1()
    test2()
    test3()
    test4_weighted()
    test5_too_many()


if __name__ == "__main__":
    main()
