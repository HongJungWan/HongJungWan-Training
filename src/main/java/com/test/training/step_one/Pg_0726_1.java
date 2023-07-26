package com.test.training.step_one;

import java.util.Arrays;

/**
 * level 2 - 최솟값 만들기
 */
public class Pg_0726_1 {
    public static void main(String[] args) {
        int a[] = {1, 4, 2};
        int b[] = {5, 4, 4};
        System.out.println(solution(a, b));
    }

    public static int solution(int[] a, int[] b) {
        int answer = 0;

        Arrays.sort(a);
        Arrays.sort(b);
        
        for (int i = 0; i < a.length; i++) {
            answer += a[i] * b[b.length - i - 1];
        }
        return answer;
    }
}
/**
 * a : 1,2,4
 * b : 4,4,5
 * <p>
 * 5 8 16
 * <p>
 * a : 1, 2
 * b : 3, 4
 * <p>
 * 4 6
 */