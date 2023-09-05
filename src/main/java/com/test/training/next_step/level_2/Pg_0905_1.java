package com.test.training.next_step.level_2;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * level 2 - 귤 고르기
 */
public class Pg_0905_1 {
    public static void main(String[] args) {
        System.out.println(solution(6, new int[]{1, 3, 2, 5, 4, 5, 2, 3}));
        System.out.println(solution(4, new int[]{1, 3, 2, 5, 4, 5, 2, 3}));

        System.out.println(solution(2, new int[]{1, 1, 1, 1, 2, 2, 2, 3}));
    }

    public static int solution(int k, int[] tangerine) {
        int answer = 0;
        Map<Integer, Integer> sizeMap = new HashMap<>();

        //섹션 1
        for (int size : tangerine) {
            sizeMap.put(size, sizeMap.getOrDefault(size, 0) + 1);
        }

        //섹션 2
        List<Integer> keyList = new ArrayList<>(sizeMap.keySet());
        keyList.sort((o1, o2) -> sizeMap.get(o2) - sizeMap.get(o1));

        //섹션 3
        int i = 0;
        while (k > 0) {
            k -= sizeMap.get(keyList.get(i));
            answer++;
            i++;
        }
        return answer;
    }
}
// List 내림차순 정렬, keyList.sort((o1, o2) -> sizeMap.get(o2) - sizeMap.get(o1));