func merge(nums1 []int, m int, nums2 []int, n int)  {
    l,r := m-1,n-1
    for i:=m+n-1;l>=0&&r>=0;i-- {
        if nums1[l] > nums2[r] {
            nums1[i] = nums1[l]
            l -= 1
        } else {
            nums1[i] = nums2[r]
            r -= 1
        }
    }
    for r >= 0 {
        nums1[r] = nums2[r]
        r -= 1
    }
}
