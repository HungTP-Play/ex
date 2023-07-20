export function Ex1_1(s: string, t: string): number {
    if (s.length < t.length) {
        return 0
    }

    if (t.length === 0) {
        return 1
    }

    let count = 0;
    if (s.startsWith(t[0])) {
        count += Ex1_1(s.slice(1), t.slice(1))
    }
    count += Ex1_1(s.slice(1), t)
    return count
}

type Item = {
    i: number,
    j: number
}
class Stack {
    data: Item[] = []
    push(i: number, j: number) {
        this.data.push({ i, j })
    }
    pop(): Item {
        return this.data.pop()!
    }

    empty(): boolean {
        return this.data.length === 0;
    }
}

export function Ex1_2(s: string, t: string) {
    const stack = new Stack()
    let i = 0; // t
    let j = 0; // s
    let count = 0;
    while (i < t.length) {
        while (j < s.length) {
            if (s[j] === t[i]) {
                if (i === t.length - 1) {
                    count++
                } else {
                    stack.push(i, j)
                    i++
                }
            }
            j++
        }
        if (stack.empty()) {
            break;
        }

        const item = stack.pop()
        i = item.i
        j = item.j + 1
    }
    return count
}

/**
 * Return the number of distinct subsequences of s that equal t
 * @param s 
 * @param t 
 * @returns 
 */
export function Ex1_3(s:string, t:string) {
    let arr = new Array(t.length + 1).fill(0)
    arr = arr.map(() => new Array(s.length + 1).fill(0))

    for(let j=0; j<s.length; j++) {
        arr[0][j] = 1
    }

    for(let i=1; i<=t.length; i++) {
        for(let j=1; j<=s.length; j++) {
            if(t[i-1] === s[j-1]) {
                arr[i][j] = arr[i-1][j-1] + arr[i][j-1]
            } else {
                arr[i][j] = arr[i][j-1]
            }
        }
    }

    return arr[t.length][s.length]
}