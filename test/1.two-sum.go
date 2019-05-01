/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (43.54%)
 * Total Accepted:    1.7M
 * Total Submissions: 3.9M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * <p>Given an array of integers, return <strong>indices</strong> of the two
 * numbers such that they add up to a specific target.</p>
 *
 * <p>You may assume that each input would have
 * <strong><em>exactly</em></strong> one solution, and you may not use the
 * <em>same</em> element twice.</p>
 *
 * <p><strong>Example:</strong></p>
 *
 * <pre>
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[<strong>0</strong>] + nums[<strong>1</strong>] = 2 + 7 = 9,
 * return [<strong>0</strong>, <strong>1</strong>].
 * </pre>
 *
 * <p>&nbsp;</p>
 *
 */
package main

import (
	"fmt"

	"github.com/MaxJohnXiang/LeetcodeGo/kit"
)

func main() {
	ss := kit.Ints2List([]int{123, 1})
	fmt.Println(ss)
}
