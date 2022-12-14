package main


/*
    Input: nums1 = [1,3], nums2 = [2]
    Output: 2.00000
    Explanation: merged array = [1,2,3] and median is 2.
    
    Input: nums1 = [1,2], nums2 = [3,4]
    Output: 2.50000
    Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var ans float64
    var nums []int
    indexN1 :=0
    indexN2 :=0
    for indexN1 < len(nums1) && indexN2 < len(nums2) {
        if(nums1[indexN1] < nums2[indexN2]){
            nums = append(nums, nums1[indexN1])
            indexN1 += 1
        } else {
            nums = append(nums, nums2[indexN2])
            indexN2 += 1
        }
    }
    for indexN1 < len(nums1) {
        nums = append(nums, nums1[indexN1])
        indexN1 += 1
    }
    for indexN2 < len(nums2) {
        nums = append(nums, nums2[indexN2])
        indexN2 += 1
    }
    l1 := len(nums)/2
    if len(nums) % 2 != 0 {
        ans = float64(nums[l1])
    } else {
        ans = float64(nums[l1] + nums[l1 - 1]) / 2.0
    }
    return ans
}
