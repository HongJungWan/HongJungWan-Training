package com.test.training.next_step.level_2.pg_9;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * level 2 - 귤 고르기 (재도전)
 */
public class Pg_0908_1 {
    public static void main(String[] args) {
        System.out.println(solution(4, new int[]{1, 3, 2, 5, 4, 5, 2, 3}));
    }

    public static int solution(int k, int[] tangerine) {
        int answer = 0, i = 0;
        Map<Integer, Integer> map = new HashMap<>();
        for (int temp : tangerine) {
            map.put(temp, map.getOrDefault(temp, 0) + 1);
        }

        List<Integer> list = new ArrayList<>(map.keySet());
        list.sort((o1, o2) -> map.get(o2) - map.get(o1));

        while (k > 0) {
            answer++;
            k -= map.get(list.get(i));
            i++;
        }
        return answer;
    }
}