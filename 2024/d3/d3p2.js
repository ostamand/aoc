const fs = require("node:fs");

// find all matches
const findAll = (pattern, memory) => {
    const matches = [];
    let match;
    while ((match = pattern.exec(memory)) !== null) {
        matches.push(match);
    }
    return matches;
};

// find instruction
const findInstructionBefore = (instructions, index) => {
    let instructionBefore = { index: 0, value: 1 };
    instructions.forEach((instruction) => {
        if (
            instruction.index < index &&
            instruction.index > instructionBefore.index
        ) {
            instructionBefore = instruction;
        }
    });
    return instructionBefore;
};

const processMemory = (memory) => {
    // process instructions
    const doList = findAll(/do\(\)/g, memory).map((match) => {
        return {
            index: match.index,
            value: 1,
        };
    });

    const dontList = findAll(/don't\(\)/g, memory).map((match) => {
        return {
            index: match.index,
            value: 0,
        };
    });

    const instructions = [...doList, ...dontList];

    // extract mul
    const pattern = /mul\((-?\d+(?:\.\d+)?),(-?\d+(?:\.\d+)?)\)/g;

    const matches = [];
    let match;
    while ((match = pattern.exec(memory)) !== null) {
        matches.push({
            value: Number(match[1]) * Number(match[2]),
            index: match.index,
        });
    }

    // calculate result
    const result = matches.reduce((accumulator, currentValue) => {
        return (
            accumulator +
            findInstructionBefore(instructions, currentValue.index).value *
                currentValue.value
        );
    }, 0);

    return result;
};

//const memory = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))";
//result = processMemory(memory);
//console.log("result", result);

fs.readFile("input.txt", "utf8", (err, data) => {
    if (err) {
        console.error(err);
        return;
    }
    const result = processMemory(data);
    console.log("result", result);
});
