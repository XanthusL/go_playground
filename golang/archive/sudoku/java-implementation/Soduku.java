package com.liyiheng;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.*;

public class Soduku {
    // 9*9 的格子
    private int[][] grid;
    // 每个格子的候选数字
    private HashMap<Integer, Set<Integer>> candidates;


    public void start() {
        initializeGrid();
        initializeCandidates();
        // 排除
        exclude();


        Solution solution = null;
        Solution nodeCusor = null;
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                int bigIndex = i * 9 + j;
                int n = grid[i][j];
                if (n == 0) {
                    Set<Integer> set = candidates.get(bigIndex);
                    Solution node = new Solution();
                    node.indexInGrid = bigIndex;
                    node.options = new int[set.size()];
                    int _tmp_index = 0;
                    for (Integer _int : set) {
                        node.options[_tmp_index++] = _int;
                    }
                    if (solution == null) {
                        solution = node;
                    } else {
                        nodeCusor.next = node;
                    }
                    nodeCusor = node;
                }
            }
        }
        boolean ok = walk(solution);
        System.out.println(ok);
        System.out.println(grid);
    }

    /**
     * 初始化候选数据
     */

    void initializeCandidates() {
        // Initialize candidate data
        candidates = new HashMap<>();
        for (int i = 0; i < 81; i++) {
            HashSet<Integer> set = new HashSet<>();
            set.add(1);
            set.add(2);
            set.add(3);
            set.add(4);
            set.add(5);
            set.add(6);
            set.add(7);
            set.add(8);
            set.add(9);
            candidates.put(i, set);
        }
    }

    /**
     * 初始化原始格子
     */
    void initializeGrid() {
        InputStreamReader reader = new InputStreamReader(System.in);
        BufferedReader bufferedReader = new BufferedReader(reader);
        grid = new int[9][9];

        for (int i = 0; i < 9; i++) {
            try {
                String line = bufferedReader.readLine();
                if (line == null) {
                    continue;
                }
                String[] numStrs = line.split(" ");
                int indice = 0;
                for (String s : numStrs) {
                    grid[i][indice++] = Integer.parseInt(s);
                }
                // 检查输入的行
                boolean ok = check(grid[i], false);
                if (!ok) {
                    System.out.println("输入有误");
                    return;
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
        // 检查每一列
        for (int i = 0; i < 9; i++) {
            boolean ok = check(getColumn(i), false);
            if (!ok) {
                System.out.println("输入有误");
                System.exit(0);
            }
        }
        // 检查3*3的格子
        for (int i = 0; i < 3; i++) {
            for (int j = 0; j < 3; j++) {
                int x = i * 3;
                int y = j * 3;
                int[] numbers = new int[]{
                        grid[x][y],
                        grid[x][y + 1],
                        grid[x][y + 2],
                        grid[x + 1][y],
                        grid[x + 1][y + 1],
                        grid[x + 1][y + 2],
                        grid[x + 2][y],
                        grid[x + 2][y + 1],
                        grid[x + 2][y + 2],
                };
                boolean ok = check(numbers, false);
                if (!ok) {
                    System.out.println("输入有误");
                    System.exit(0);
                }
            }
        }
    }


    /**
     * 获取列
     *
     * @param index 0-8
     */
    int[] getColumn(int index) {
        return new int[]{
                grid[0][index],
                grid[1][index],
                grid[2][index],
                grid[3][index],
                grid[4][index],
                grid[5][index],
                grid[6][index],
                grid[7][index],
                grid[8][index],
        };
    }

    void exclude() {
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                int bigIndex = i * 9 + j;
                int v = grid[i][j];
                if (v != 0) {
                    candidates.remove(bigIndex);
                    continue;
                }
                removeCandidate(getColumn(i), bigIndex);
                removeCandidate(grid[i], bigIndex);

                int x = i / 3;
                int y = j / 3;
                x *= 3;
                y *= 3;
                int[] ints = {
                        grid[x][y],
                        grid[x][y + 1],
                        grid[x][y + 2],
                        grid[x + 1][y],
                        grid[x + 1][y + 1],
                        grid[x + 1][y + 2],
                        grid[x + 2][y],
                        grid[x + 2][y + 1],
                        grid[x + 2][y + 2]};
                removeCandidate(ints, bigIndex);


                Set<Integer> integers = candidates.get(bigIndex);
                if (integers.size() == 1) {
                    for (Integer intgr : integers) {
                        grid[i][j] = intgr;
                    }
                }
            }

        }
    }

    private void removeCandidate(int[] nums, int index) {
        Set<Integer> integers = candidates.get(index);
        for (int n : nums) {
            if (n != 0) {
                integers.remove(nums);
            }
        }
    }


    /**
     * 检查一组数是否是1-9且不重复
     *
     * @param numbers 长度为9的数组
     * @param fulled  是否已经完全填充（空白用0表示，未完全填充时，不考虑0的情况）
     * @return 是否可用
     */
    public boolean check(int[] numbers, boolean fulled) {
        int len = numbers.length;
        int[] copy = Arrays.copyOf(numbers, len);
        Arrays.sort(copy);

        for (int i = 1; i < len; i++) {
            int v = copy[i];
            if (!fulled && v == 0) {
                continue;
            }
            if (v == copy[i - 1]) {
                return false;
            }
            if (v > 9) {
                return false;
            }
        }
        return true;

    }

    boolean isValidate() {
        return false;
    }

    boolean walk(Solution node) {
        if (node == null) {
            return isValidate();
        }
        outer:
        for (int v : node.options) {
            int bigIndex = node.indexInGrid;

            int[] ints0 = grid[bigIndex % 9];
            int[] ints1 = getColumn(bigIndex / 9);
            int[] ints2 = getSubGrid(bigIndex / 9, bigIndex % 9);

            int[] relations = new int[27];
            for (int i = 0; i < 9; i++) {
                relations[i] = ints0[i];
            }
            for (int i = 0; i < 9; i++) {
                relations[9 + i] = ints1[i];
            }
            for (int i = 0; i < 9; i++) {
                relations[18 + i] = ints2[i];
            }

            for (int rv : relations) {
                if (rv != 0 && rv == v) {
                    continue outer;
                }
            }
            grid[node.indexInGrid / 9][node.indexInGrid % 9] = v;
            if (walk(node.next)) {
                return true;
            } else {
                grid[node.indexInGrid / 9][node.indexInGrid % 9] = 0;
            }

        }
        return false;
    }

    int[] getSubGrid(int x, int y) {

        x /= 3;
        x *= 3;
        y /= 3;
        y *= 3;
        return new int[]{
                grid[x][y],
                grid[x][y + 1],
                grid[x][y + 2],
                grid[x + 1][y],
                grid[x + 1][y + 1],
                grid[x + 1][y + 2],
                grid[x + 2][y],
                grid[x + 2][y + 1],
                grid[x + 2][y + 2]};


    }
}
