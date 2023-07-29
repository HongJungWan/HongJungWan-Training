package com.test.training.next_step.level_1;

/**
 * level 1 - 최소 직사각형
 */
public class Pg_0729_6 {
    public static void main(String[] args) {
        int arr[][] = {{60, 50}, {30, 70}, {60, 30}, {80, 40}};
        System.out.println(solution(arr));
    }

    public static int solution(int[][] sizes) {
        int max_width = 0;
        int max_height = 0;

        for (int[] size : sizes) {
            int width = Math.max(size[0], size[1]); // 가장 긴 변을 가로로
            int height = Math.min(size[0], size[1]); // 짧은 변을 세로로

            max_width = Math.max(max_width, width); // 가장 긴 가로 길이 갱신
            max_height = Math.max(max_height, height); // 가장 긴 세로 길이 갱신
        }

        return max_width * max_height; // 가장 긴 가로 길이와 세로 길이를 곱해서 반환
    }
}