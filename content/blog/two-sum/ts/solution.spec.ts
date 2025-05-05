import { twoSum } from "./solution";

type TestCase = {
    name: string
    nums: number[]
    target: number
    expRes: number[]
}

const cases: TestCase[] = [
    {
        name: "[2, 7, 11, 15], 9",
        nums: [2, 7, 11, 15],
        target: 9,
        expRes: [1, 0],
    },
    {
        name: "[3, 2, 4], 6",
        nums: [3, 2, 4],
        target: 6,
        expRes: [2, 1],
    },
    {
        name: "[-1, 0], -1",
        nums: [-1, 0],
        target: -1,
        expRes: [1, 0],
    },
    {
        name: "[0, 0], 5",
        nums: [0, 0],
        target: 5,
        expRes: [],
    },
]

describe('twoSum', () => {
    cases.forEach(testCase => {
        test(testCase.name, () => {
            const res = twoSum(testCase.nums, testCase.target)
            expect(res).toEqual(testCase.expRes);
        })
    })
})