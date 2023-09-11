package com.test.training.next_step.level_2;

import java.util.HashSet;
import java.util.Set;

/**
 * level 2 - 연속 부분 수열 합의 개수
 */
public class Pg_0911_1 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{7, 9, 1, 1, 4}));
    }

    public static int solution(int[] elements) {
        int[] elementsExt = new int[elements.length * 2];

        for (int i = 0; i < elements.length; i++) {
            elementsExt[i] = elements[i];
            elementsExt[i + elements.length] = elements[i];
        }

        Set<Integer> set = new HashSet<>();
        for (int i = 0; i < elements.length; i++) {
            int sum = 0;
            for (int j = i; j < i + elements.length; j++) {
                sum += elementsExt[j];
                set.add(sum);
            }
        }
        return set.size();
    }
}

// 브루트 포스 (Brute Force)
// 원형 구조 문제를 -> 선형 구조로 변환하여 해결하는 방식

// 원형 수열은 선형 수열의 처음과 끝이 연결된 형태. 그래서 연속 부분 수열을 찾을 때, 처음과 끝 부분의 연결도 고려해야 한다.

// 원형 수열의 연속 부분 수열 문제를 간단히 하기 위한 전략: 원래 수열을 두 번 연결하여 확장된 선형 수열을 생성.
// 예: [a, b, c] -> [a, b, c, a, b, c]