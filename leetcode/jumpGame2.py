
def jump(nums):
    jumps = 0
    endRange = 0
    furthest = 0

    for i in range(0, len(nums) - 1):
        furthest = max(furthest, nums[i] + i)

        if endRange == i:
            jumps += 1
            endRange = furthest
    return jumps

print(jump([2, 3, 1, 1, 4]))