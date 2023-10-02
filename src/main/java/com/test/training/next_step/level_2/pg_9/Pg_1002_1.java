package com.test.training.next_step.level_2.pg_9;

import java.util.LinkedList;
import java.util.Queue;

/**
 * level 2 - 다리를 지나는 트럭
 */
public class Pg_1002_1 {
    public static void main(String[] args) {
        System.out.println(solution(2, 10, new int[]{7, 4, 5, 6}));
    }

    public static int solution(int bridge_length, int weight, int[] truck_weights) {
        int answer = 0;
        Queue<Integer> q = new LinkedList<>();
        int sum = 0; // 다리를 건너는 트럭들 무게 합

        for (int t : truck_weights) {
            while (true) {
                if (q.isEmpty()) { //큐가 비어있다면 다음 트럭 삽입
                    q.add(t);
                    sum += t;
                    answer++;
                    break;
                }
                if (q.size() == bridge_length) { //큐의 사이즈와 다리의 길이가 같다면 큐에서 큐에서 처음 값을 빼고 최대 무게 -
                    sum -= q.poll();
                }
                if (sum + t > weight) { //다음 트럭이 최대 무게 초과 or 최대 무게 이내
                    q.add(0);
                    answer++;
                } else {
                    q.add(t);
                    sum += t;
                    answer++;
                    break;
                }
            }
        }
        return answer + bridge_length; //걸린 시간 + 마지막 트럭의 통과시간 (다리의 길이)
    }
}