package com.test.training.next_step.level_2;

/**
 * level 2 - 이진 변환 반복하기
 */
public class Pg_0728_1 {
    public static void main(String[] args) {
        System.out.println(solution("110010101001"));
    }

    public static int[] solution(String s) {
        int[] answer = new int[2];
        int zeroCount = 0;
        int operationCount = 0;

        while (s.length() != 1) {
            for (char c : s.toCharArray()) {
                if (c == '0') {
                    zeroCount++;
                }
                answer[1] = zeroCount;
            }
            s = s.replaceAll("0", "");
            s = Integer.toBinaryString(s.length());
            operationCount++;
            answer[0] = operationCount;
        }
        return answer;
    }
}
/**
 * "110010101001"이 "1"이 될 때까지 이진 변환을 가하는 과정
 * <p>
 * 3번의 이진 변환을 하는 동안 8개의 0을 제거했으므로, [3,8]을 return
 */