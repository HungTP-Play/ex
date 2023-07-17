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