class Smallfuck {
    constructor(code, tape) {
        this.tape = tape.split('');
        this.pointer = 0

        this.code = code;
        this.char = '';
        this.readPos = 0;
        this.bracketsIndex = this.createBracketPairs()
    }

    run() {
        this.readCode();

        while (this.readPos <= this.code.length) {
            this.executeCode();
        }

        return this.tape.join('');
    }

    readCode() {
        this.char = this.code[this.readPos];
        this.readPos++;
    }

    createBracketPairs() {
        let a = this.code.split('')
        let stack = []
        let pairs = {}

        for (let i = 0; i < a.length; i++) {
            if (a[i] === "[") {
                stack.push(i)
            }
            if (a[i] === "]") {
                let openIndex = stack.pop();
                let closeIndex = i;
                pairs[openIndex] = closeIndex;
                pairs[closeIndex] = openIndex;
            }
        }
        return pairs
    }

    executeCode() {
        switch (this.char) {
            case '>':
                this.incrementPointer();
                break;

            case '<':
                this.decrementPointer();
                break;

            case '*':
                this.flipBit();
                break;

            case '[':
                if (!this.checkBit()) {
                    this.jumpToMatchingBracket();
                }
                break;

            case ']':
                if (this.checkBit()) {
                    this.jumpToMatchingBracket();
                }
                break;
        }
        this.readCode();
    }

    jumpToMatchingBracket() {
        this.char = this.code[this.bracketsIndex[this.readPos - 1]]
        this.readPos = this.bracketsIndex[this.readPos - 1] + 1
        console.log(this.char, this.readPos)
    }

    flipBit() {
        if (this.tape[this.pointer] === '0') {
            this.tape[this.pointer] = '1'
        } else {
            this.tape[this.pointer] = '0'
        }
    }

    checkBit() {
        return this.tape[this.pointer] === 1
    }

    incrementPointer() {
        if (this.pointer === this.code.length) {
            return
        } else {
            this.pointer++;
        }
    }

    decrementPointer() {
        if (this.pointer === 0) {
            return
        } else {
            this.pointer--;
        }
    }
}


function interpreter(code, tape) {
    let sf = new Smallfuck(code, tape)
    return sf.run();
}

console.log(interpreter("[>*]", "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"))