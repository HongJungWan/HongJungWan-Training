package com.test.training.next_step.level_2.pg_9;

/**
 * level 2 - 행렬의 곱셈
 */
public class Pg_0916_1 {
    public static void main(String[] args) {
        int[][] arr1 = {{1, 4}, {3, 2}, {4, 1}};
        int[][] arr2 = {{3, 3}, {3, 3}};

        System.out.println(solution(arr1, arr2));
    }

    public static int[][] solution(int[][] arr1, int[][] arr2) {
        int[][] answer = new int[arr1.length][arr2[0].length];

        for (int i = 0; i < arr1.length; i++) {
            for (int j = 0; j < arr2[0].length; j++) {
                int val = 0;
                for (int k = 0; k < arr1[0].length; k++) {
                    val += arr1[i][k] * arr2[k][j];
                }
                answer[i][j] = val;
            }
        }
        return answer;
    }
}