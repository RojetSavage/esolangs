class Paintfuck {
    constructor(code, iterations, width, height) {
        console.log(code, iterations, width, height)
        this.dataGrid = new Array(height);
        this.pointer = [0, 0];
        this.width = width;
        this.height = height;
        this.code = code;
        this.char = '';
        this.readPos = 0;
        this.iterations = iterations;
        this.commandsExecuted = 0;

        this.bracketsIndex = this.createBracketPairs()

        for (let i = 0; i < height; i++) {
            this.dataGrid[i] = new Array(width).fill(0);
        }
    }

    run() {
        this.readCode();

        while (this.commandsExecuted < this.iterations && this.readPos <= this.code.length) {
            this.executeCode();
        }

        return this.printDatagrid();
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

        this.bracketsIndex = pairs;
        return pairs 
    }

    executeCode() {
        switch (this.char) {
            case 'n':
                this.pointer = this.getNorthPointer();
                break;

            case 's':
                this.pointer = this.getSouthPointer();
                break;

            case 'e':
                this.pointer = this.getEastPointer();
                break;

            case 'w':
                this.pointer = this.getWestPointer();
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

            default:
                this.readCode()
                return
        }
        
        this.commandsExecuted++;
        this.readCode();    
    }
    
    jumpToMatchingBracket() {
        this.char = this.code[this.bracketsIndex[this.readPos - 1]]
        this.readPos = this.bracketsIndex[this.readPos - 1] + 1
    }

    flipBit() {
        if (this.checkBit()) {
            this.dataGrid[this.pointer[0]][this.pointer[1]] = 0;
        } else if (!this.checkBit()) {
            this.dataGrid[this.pointer[0]][this.pointer[1]] = 1;
        }
    }

    checkBit() {
        return this.dataGrid[this.pointer[0]][this.pointer[1]] === 1;
    }

    getWestPointer() {
        if (this.pointer[1] === 0) {
            return [this.pointer[0], this.width - 1];
        } else {
            return [this.pointer[0], this.pointer[1] - 1];
        }
    }

    getEastPointer() {
        if (this.pointer[1] === this.width - 1) {
            return [this.pointer[0], 0];
        } else {
            return [this.pointer[0], this.pointer[1] + 1];
        }
    }

    getNorthPointer() {
        if (this.pointer[0] === 0) {
            return [this.height - 1, this.pointer[1]];
        } else {
            return [this.pointer[0] - 1, this.pointer[1]];
        }
    }

    getSouthPointer() {
        if (this.pointer[0] === this.height - 1) {
            return [0, this.pointer[1]];
        } else {
            return [this.pointer[0] + 1, this.pointer[1]];
        }
    }

    printDatagrid() {
        let str = [];
        for (let i = 0; i < this.dataGrid.length; i++) {
            str.push(this.dataGrid[i].toString().split(',').join(''));
        }
        return str.join("\r\n");
    }
}


function interpreter(code, iterations, width, height) {
    let pf = new Paintfuck(code, iterations, width, height)
    return pf.run();
}

console.log(interpreter("*[s[e]*]", 100, 5, 5))