package com.test.training.next_step.level_2.pg_9;

import java.util.HashMap;

/**
 * level 2 - 의상
 */
public class Pg_0921_1 {
    public static void main(String[] args) {
        System.out.println(solution(new String[][]{
                {"yellow_hat", "headgear"},
                {"blue_sunglasses", "eyewear"},
                {"green_turban", "headgear"}
        }));
    }

    public static int solution(String[][] clothes) {
        int answer = 1;

        HashMap<String, Integer> hashMap = new HashMap<>();

        for (int i = 0; i < clothes.length; i++) {
            hashMap.put(clothes[i][1], hashMap.getOrDefault(clothes[i][1], 0) + 1);
        }

        for (int value : hashMap.values()) {
            answer *= value + 1;
        }
        return answer - 1;
    }
}
// 1. yellow_hat
// 2. blue_sunglasses
// 3. green_turban
// 4. yellow_hat + blue_sunglasses
// 5. green_turban + blue_sunglasses
// 6. X