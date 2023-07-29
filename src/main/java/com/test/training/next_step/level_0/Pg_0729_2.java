package com.test.training.next_step.level_0;

/**
 * level 0 - 수 조작하기 1
 */
public class Pg_0729_2 {
    public static void main(String[] args) {
        System.out.println(solution(0, "wsdawsdassw"));
    }

    public static int solution(int n, String control) {
        for (int i = 0; i < control.length(); i++) {
            if (control.charAt(i) == 'w') {
                n += 1;
            }
            if (control.charAt(i) == 's') {
                n -= 1;
            }
            if (control.charAt(i) == 'd') {
                n += 10;
            }
            if (control.charAt(i) == 'a') {
                n -= 10;
            }
        }
        return n;
    }
}
/**
 * "w" : n이 1 커집니다.
 * "s" : n이 1 작아집니다.
 * "d" : n이 10 커집니다.
 * "a" : n이 10 작아집니다.
 */